package oidc

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"golang.org/x/oauth2"
)

// generateState はランダムなstate値を生成します
func generateState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	// ランダムなstate値を生成（CSRF対策）
	state, err := generateState()
	if err != nil {
		http.Error(w, "Failed to generate state", http.StatusInternalServerError)
		return
	}

	// 認証URLを生成してリダイレクト
	url := OIDCConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
