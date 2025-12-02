package main

import (
	"log"
	"net/http"

	"github.com/moepig/oauthapp/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/auth/login", handlers.LoginHandler)
	http.HandleFunc("/auth/callback", handlers.CallbackHandler)

	log.Println("Server starting on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
