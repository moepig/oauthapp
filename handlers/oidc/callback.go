package oidc

import (
	"context"
	"html/template"
	"net/http"
	"path/filepath"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータから認証コードを取得
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// 認証コードをトークンと交換
	token, err := OIDCConfig.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ID Tokenを取得
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token in token response", http.StatusInternalServerError)
		return
	}

	// ID Tokenを検証
	if OIDCVerifier == nil {
		http.Error(w, "OIDC verifier not initialized", http.StatusInternalServerError)
		return
	}

	idToken, err := OIDCVerifier.Verify(ctx, rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ID Tokenからクレームを取得
	var claims map[string]interface{}
	if err := idToken.Claims(&claims); err != nil {
		http.Error(w, "Failed to parse claims: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// テンプレートで表示
	layoutPath := filepath.Join("templates", "layout.html")
	tmplPath := filepath.Join("templates", "oidc", "callback.html")
	tmpl, err := template.ParseFiles(layoutPath, tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Claims":      claims,
		"IDToken":     rawIDToken,
		"AccessToken": token.AccessToken,
		"TokenType":   token.TokenType,
	}

	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
