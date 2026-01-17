package oauth

import (
	"os"
	"strings"

	"golang.org/x/oauth2"
)

var OAuthConf *oauth2.Config

func init() {
	OAuthConf = &oauth2.Config{
		ClientID:     os.Getenv("OAUTHAPP_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTHAPP_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTHAPP_REDIRECT_URL"),
		Scopes:       strings.Split(os.Getenv("OAUTHAPP_SCOPES"), ","),
		Endpoint: oauth2.Endpoint{
			AuthURL:  os.Getenv("OAUTHAPP_AUTH_URL"),
			TokenURL: os.Getenv("OAUTHAPP_TOKEN_URL"),
		},
	}
}
