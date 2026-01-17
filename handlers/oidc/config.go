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
	ctx := context.Background()

	// 環境変数から設定を読み込む
	clientID := os.Getenv("OIDC_CLIENT_ID")
	clientSecret := os.Getenv("OIDC_CLIENT_SECRET")
	scopes := os.Getenv("OIDC_SCOPES")
	issuer := os.Getenv("OIDC_ISSUER")

	// デフォルトスコープ
	if scopes == "" {
		scopes = "openid,profile,email"
	}

	var authURL, tokenURL string

	// OIDC Discovery Endpointを使用してエンドポイント情報を取得
	if issuer != "" {
		provider, err := oidc.NewProvider(ctx, issuer)
		if err != nil {
			log.Printf("Warning: Failed to create OIDC provider from issuer %s: %v", issuer, err)
			log.Printf("Falling back to manual endpoint configuration")
		} else {
			// Discovery Endpointから取得
			authURL = provider.Endpoint().AuthURL
			tokenURL = provider.Endpoint().TokenURL

			// ID Token検証用のVerifierを作成
			OIDCVerifier = provider.Verifier(&oidc.Config{
				ClientID: clientID,
			})

			log.Printf("OIDC Discovery successful for issuer: %s", issuer)
			log.Printf("  Auth URL: %s", authURL)
			log.Printf("  Token URL: %s", tokenURL)
		}
	}

	// 手動設定がある場合は上書き（Discovery Endpointより優先）
	if manualAuthURL := os.Getenv("OIDC_AUTH_URL"); manualAuthURL != "" {
		authURL = manualAuthURL
		log.Printf("Using manual OIDC_AUTH_URL: %s", authURL)
	}
	if manualTokenURL := os.Getenv("OIDC_TOKEN_URL"); manualTokenURL != "" {
		tokenURL = manualTokenURL
		log.Printf("Using manual OIDC_TOKEN_URL: %s", tokenURL)
	}

	// OIDC設定
	OIDCConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  os.Getenv("OIDC_REDIRECT_URL"),
		Scopes:       strings.Split(scopes, ","),
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
	}

	// 設定が不完全な場合は警告
	if clientID == "" || authURL == "" || tokenURL == "" {
		log.Printf("Warning: OIDC configuration is incomplete")
		log.Printf("  OIDC_CLIENT_ID: %s", clientID)
		log.Printf("  Auth URL: %s", authURL)
		log.Printf("  Token URL: %s", tokenURL)
	}
}
