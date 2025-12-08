package oauth

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータから認証コードを取得
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	// 認証コードをアクセストークンと交換
	token, err := OAuthConf.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// トークンを使用してHTTPクライアントを作成
	client := OAuthConf.Client(r.Context(), token)

	// ユーザー情報エンドポイントを環境変数から取得
	userInfoURL := os.Getenv("OAUTHAPP_USERINFO_URL")
	if userInfoURL == "" {
		http.Error(w, "OAUTHAPP_USERINFO_URL is not set", http.StatusInternalServerError)
		return
	}

	// ユーザー情報を取得
	resp, err := client.Get(userInfoURL)
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// ユーザー情報をパース
	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to decode user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// ユーザー情報とトークンをテンプレートで表示
	layoutPath := filepath.Join("templates", "layout.html")
	tmplPath := filepath.Join("templates", "oauth", "callback.html")
	tmpl, err := template.ParseFiles(layoutPath, tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"UserInfo":    userInfo,
		"AccessToken": token.AccessToken,
		"TokenType":   token.TokenType,
	}

	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
