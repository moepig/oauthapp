package oauth

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// SettingsHandler はOAuth関連の環境変数を表示するハンドラー
func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	// OAuth関連の環境変数を取得
	settings := map[string]interface{}{
		"ClientID":     os.Getenv("OAUTHAPP_CLIENT_ID"),
		"ClientSecret": maskSecret(os.Getenv("OAUTHAPP_CLIENT_SECRET")),
		"RedirectURL":  os.Getenv("OAUTHAPP_REDIRECT_URL"),
		"Scopes":       os.Getenv("OAUTHAPP_SCOPES"),
		"ScopesList":   strings.Split(os.Getenv("OAUTHAPP_SCOPES"), ","),
		"AuthURL":      os.Getenv("OAUTHAPP_AUTH_URL"),
		"TokenURL":     os.Getenv("OAUTHAPP_TOKEN_URL"),
		"UserInfoURL":  os.Getenv("OAUTHAPP_USERINFO_URL"),
	}

	// テンプレートを読み込み
	layoutPath := filepath.Join("templates", "layout.html")
	tmplPath := filepath.Join("templates", "oauth", "settings.html")
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
