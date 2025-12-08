package main

import (
	"log"
	"net/http"

	"github.com/moepig/oauthapp/handlers"
)

func main() {
	// 静的ファイルの配信設定
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/auth/login", handlers.LoginHandler)
	http.HandleFunc("/auth/callback", handlers.CallbackHandler)

	log.Println("Server starting on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
