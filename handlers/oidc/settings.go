package oidc

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// SettingsHandler はOIDC関連の環境変数を表示するハンドラー
func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	// OIDC関連の環境変数を取得
	settings := map[string]interface{}{
		"ClientID":     os.Getenv("OIDC_CLIENT_ID"),
		"ClientSecret": maskSecret(os.Getenv("OIDC_CLIENT_SECRET")),
		"Scopes":       os.Getenv("OIDC_SCOPES"),
		"AuthURL":      os.Getenv("OIDC_AUTH_URL"),
		"TokenURL":     os.Getenv("OIDC_TOKEN_URL"),
		"Issuer":       os.Getenv("OIDC_ISSUER"),
		"JWKSURL":      os.Getenv("OIDC_JWKS_URL"),
	}

	// テンプレートを読み込み
	layoutPath := filepath.Join("templates", "layout.html")
	tmplPath := filepath.Join("templates", "oidc", "settings.html")
	tmpl, err := template.ParseFiles(layoutPath, tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// テンプレートを実行
	if err := tmpl.ExecuteTemplate(w, "layout", settings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// maskSecret はシークレット値をマスクする
func maskSecret(secret string) string {
	if secret == "" {
		return "(未設定)"
	}
	return "****"
}
