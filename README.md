# OAuthApp

## 環境変数

### OAuth 2.0 設定

| 変数名 | 必須/任意 | 型 | 説明 |
|--------|----------|-----|------|
| `OAUTHAPP_CLIENT_ID` | **必須** | String | OAuthクライアントID。認証プロバイダーから発行されたクライアントIDを設定します。 |
| `OAUTHAPP_CLIENT_SECRET` | **必須** | String | OAuthクライアントシークレット。認証プロバイダーから発行されたシークレットを設定します。 |
| `OAUTHAPP_SCOPES` | **必須** | String (カンマ区切り) | 要求するスコープ。カンマ区切りで複数指定可能。例: `read:user,user:email` |
| `OAUTHAPP_AUTH_URL` | **必須** | String (URL) | OAuth認可エンドポイントURL。例: `https://github.com/login/oauth/authorize` |
| `OAUTHAPP_TOKEN_URL` | **必須** | String (URL) | OAuthトークンエンドポイントURL。例: `https://github.com/login/oauth/access_token` |
| `OAUTHAPP_USERINFO_URL` | **必須** | String (URL) | ユーザー情報取得エンドポイントURL。例: `https://api.github.com/user` |

### OIDC 設定

OIDC設定には2つの方法があります:

#### Discovery Endpoint 使用

| 変数名 | 必須/任意 | 型 | 説明 |
|--------|----------|-----|------|
| `OIDC_CLIENT_ID` | **必須** | String | OIDCクライアントID。認証プロバイダーから発行されたクライアントIDを設定します。 |
| `OIDC_CLIENT_SECRET` | **必須** | String | OIDCクライアントシークレット。認証プロバイダーから発行されたシークレットを設定します。 |
| `OIDC_ISSUER` | **必須** | String (URL) | OIDC Issuer URL。`.well-known/openid-configuration`から自動的にエンドポイント情報を取得します。例: `https://accounts.google.com` |
| `OIDC_SCOPES` | 任意 | String (カンマ区切り) | 要求するスコープ。デフォルト: `openid,profile,email` |

この方法では、`OIDC_AUTH_URL`と`OIDC_TOKEN_URL`は自動的に取得されます。

#### 手動設定

| 変数名 | 必須/任意 | 型 | 説明 |
|--------|----------|-----|------|
| `OIDC_CLIENT_ID` | **必須** | String | OIDCクライアントID。 |
| `OIDC_CLIENT_SECRET` | **必須** | String | OIDCクライアントシークレット。 |
| `OIDC_AUTH_URL` | **必須** | String (URL) | OIDC認可エンドポイントURL。 |
| `OIDC_TOKEN_URL` | **必須** | String (URL) | OIDCトークンエンドポイントURL。 |
| `OIDC_ISSUER` | **必須** | String (URL) | OIDC Issuer URL。ID Token検証に使用されます。 |
| `OIDC_SCOPES` | 任意 | String (カンマ区切り) | 要求するスコープ。デフォルト: `openid,profile,email` |

`OIDC_ISSUER`が設定されている場合でも、`OIDC_AUTH_URL`または`OIDC_TOKEN_URL`を明示的に設定すると、手動設定が優先されます。
