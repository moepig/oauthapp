package oidc

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var (
	OIDCConfig   *oauth2.Config
	OIDCVerifier *oidc.IDTokenVerifier
)

func init() {
	// 環境変数から設定を読み込む
	clientID := os.Getenv("OIDC_CLIENT_ID")
	clientSecret := os.Getenv("OIDC_CLIENT_SECRET")
	scopes := os.Getenv("OIDC_SCOPES")
	authURL := os.Getenv("OIDC_AUTH_URL")
	tokenURL := os.Getenv("OIDC_TOKEN_URL")
	issuer := os.Getenv("OIDC_ISSUER")

	// デフォルトスコープ
	if scopes == "" {
		scopes = "openid,profile,email"
	}

	// OIDC設定
	OIDCConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       strings.Split(scopes, ","),
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
	}

	// ID Token検証用のVerifierを作成
	if issuer != "" {
		ctx := context.Background()
		provider, err := oidc.NewProvider(ctx, issuer)
		if err != nil {
			log.Printf("Warning: Failed to create OIDC provider: %v", err)
			return
		}

		OIDCVerifier = provider.Verifier(&oidc.Config{
			ClientID: clientID,
		})
	}
}
