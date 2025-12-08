package main

import (
	"log"
	"net/http"

	"github.com/moepig/oauthapp/handlers"
	"github.com/moepig/oauthapp/handlers/oauth"
	"github.com/moepig/oauthapp/handlers/oidc"
)

func main() {
	// 静的ファイルの配信設定
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/oauth/authorize", oauth.AuthorizeHandler)
	http.HandleFunc("/oauth/callback", oauth.CallbackHandler)
	http.HandleFunc("/oauth/settings", oauth.SettingsHandler)
	http.HandleFunc("/oidc/authorize", oidc.AuthorizeHandler)
	http.HandleFunc("/oidc/callback", oidc.CallbackHandler)
	http.HandleFunc("/oidc/settings", oidc.SettingsHandler)

	log.Println("Server starting on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
