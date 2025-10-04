# Frontend

Next.js 15 + Auth0 OIDC認証フロントエンド。

## セットアップ

```bash
cp .env.local.example .env.local
# AUTH0_SECRETは: openssl rand -hex 32

npm install
npm run dev
```

## 設計の背景

### なぜ`@auth0/nextjs-auth0`を使わないか
Next.js 15の新しい非同期API（`await cookies()`等）に未対応のため、カスタム実装。

### トークン管理の方針
- Access/ID TokenをHTTP Onlyクッキーに保存（XSS対策）
- LocalStorageは使用しない
- `/api/backend/*`プロキシでサーバー側でトークンを付与

### なぜプロキシ経由でバックエンドを呼ぶか
- フロントエンドのJavaScriptからトークンを隠蔽
- CORS設定を簡素化
- トークンリフレッシュを一箇所で管理可能

## Auth0設定の注意点

### Grant Types
`Authorization Code`を必ず有効化。デフォルトでは無効の場合あり。

### Callback URL
完全一致が必要。`http://localhost:3000/api/auth/callback`（末尾スラッシュなし）
