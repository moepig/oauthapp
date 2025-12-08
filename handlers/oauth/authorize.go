package oauth

import (
	"net/http"

	"golang.org/x/oauth2"
)

func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	url := OAuthConf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
