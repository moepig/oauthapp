package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/oauth2"
)

var oauthConf *oauth2.Config

func init() {
	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("OAUTHAPP_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTHAPP_CLIENT_SECRET"),
		Scopes:       strings.Split(os.Getenv("OAUTHAPP_SCOPES"), ","),
		Endpoint: oauth2.Endpoint{
			AuthURL:  os.Getenv("OAUTHAPP_AUTH_URL"),
			TokenURL: os.Getenv("OAUTHAPP_TOKEN_URL"),
		},
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	url := oauthConf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータから認証コードを取得
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	// 認証コードをアクセストークンと交換
	token, err := oauthConf.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// トークンを使用してHTTPクライアントを作成
	client := oauthConf.Client(r.Context(), token)

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
	tmplPath := filepath.Join("templates", "callback.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"UserInfo":    userInfo,
		"AccessToken": token.AccessToken,
		"TokenType":   token.TokenType,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
