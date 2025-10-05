# Auth0 Sandbox Backend

Auth0を使用したOIDC認証を実装するバックエンドAPI

## セットアップ

### 環境変数

`.env`ファイルを作成して以下の環境変数を設定してください：

```
AUTH0_DOMAIN=your-tenant.auth0.com
AUTH0_AUDIENCE=your-api-audience
CORS_ALLOWED_ORIGINS=http://localhost:3000
PORT=8080
```

### 依存関係のインストール

```bash
go mod download
```

## OpenAPI仕様からのコード生成

このプロジェクトでは、OpenAPI仕様から型定義とサーバーインターフェースを自動生成するために [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) を使用しています。

### コード生成の実行

OpenAPI仕様（`api/openapi.yaml`）を更新した後、以下のコマンドでコードを再生成してください：

```bash
make generate
```

生成されたコードは `internal/generated/api.gen.go` に出力されます。

### 生成される内容

- **モデル型定義**: OpenAPIスキーマから生成されるGo構造体
- **ServerInterface**: APIハンドラーが実装すべきインターフェース
- **リクエスト/レスポンス型**: JSONリクエストボディとレスポンスの型定義

## 開発

### サーバーの起動

```bash
make run
```

または

```bash
go run cmd/server/main.go
```

### ビルド

```bash
make build
```

実行可能ファイルは `bin/server` に出力されます。

## API仕様

API仕様は `api/openapi.yaml` を参照してください。

主なエンドポイント：

- `GET /api/v1/users/me` - 現在のユーザー情報を取得
- `GET /api/v1/users/me/profile` - ユーザープロフィールを取得
- `PUT /api/v1/users/me/profile` - ユーザープロフィールを更新
- `GET /api/v1/users/me/data` - ユーザーデータのリストを取得
- `POST /api/v1/users/me/data` - ユーザーデータを作成
