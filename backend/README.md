# Backend API

Auth0 OIDC認証を使用したGo APIサーバー。

## セットアップ

```bash
cp .env.example .env
# .envにAuth0の情報を設定

go mod tidy
go run cmd/server/main.go
```

## 設計の背景

### なぜインメモリストアか
- サンドボックスとして簡潔に保つため
- 本番ではPostgreSQL等のDBに置き換え（`internal/repository/`層を実装）

### JWT検証の仕組み
- Auth0のJWKS（公開鍵）を5分間キャッシュして検証
- `sub`クレームをユーザーIDとして使用（Auth0の一意識別子）
- Audienceチェックで意図したAPI向けのトークンかを確認

### CORS設定
- フロントエンドからのクロスオリジンリクエストを許可
- `credentials: "include"`でクッキーを送信可能に

## トラブルシューティング

### Token validation error
Auth0のAPIで該当アプリケーションが認可されているか確認。Machine to Machine Applicationsタブで`Authorized`をONに。
