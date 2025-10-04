# Auth0 Terraform設定

このディレクトリには、Auth0の設定をコードで管理するためのTerraform設定ファイルが含まれています。

## 前提条件

1. Terraformのインストール
```bash
brew install terraform
```

2. Auth0 Management APIの認証情報を取得

### Management API認証情報の取得方法

**方法1: 既存のManagement APIアプリケーションを使用**

1. Auth0ダッシュボード → **Applications** → **Applications**
2. **Auth0 Management API (Test Application)** を探す
3. **Settings**タブでClient IDとClient Secretを確認

**方法2: 新しいMachine to Machineアプリケーションを作成**

1. Auth0ダッシュボード → **Applications** → **Applications** → **Create Application**
2. 名前: `Terraform Management API`
3. アプリケーションタイプ: **Machine to Machine Applications**
4. API: **Auth0 Management API** を選択
5. 必要な権限を選択:
   - `read:clients`
   - `create:clients`
   - `update:clients`
   - `delete:clients`
   - `read:resource_servers`
   - `create:resource_servers`
   - `update:resource_servers`
   - `delete:resource_servers`
   - `read:client_grants`
   - `create:client_grants`
   - `update:client_grants`
   - `delete:client_grants`
   - `read:connections`
   - `update:connections`
6. **Authorize**をクリック
7. Client IDとClient Secretをコピー

## セットアップ

1. `terraform.tfvars`ファイルを作成

```bash
cd terraform
cp terraform.tfvars.example terraform.tfvars
```

2. `terraform.tfvars`を編集してManagement API認証情報を設定

```hcl
auth0_management_client_id     = "YOUR_MANAGEMENT_CLIENT_ID"
auth0_management_client_secret = "YOUR_MANAGEMENT_CLIENT_SECRET"
```

3. Terraformを初期化

```bash
terraform init
```

4. 実行計画を確認

```bash
terraform plan
```

5. 適用

```bash
terraform apply
```

## 出力の確認

適用後、以下のコマンドで設定値を確認できます:

```bash
# Client IDを確認
terraform output client_id

# Client Secret を確認
terraform output -raw client_secret

# 環境変数設定を確認
terraform output -raw env_config
```

## リソース削除

Auth0から設定を削除する場合:

```bash
terraform destroy
```

## 注意事項

- `terraform.tfvars`には機密情報が含まれるため、Gitにコミットしないでください（`.gitignore`で除外済み）
- 本番環境では、Terraform Cloudやリモートバックエンドを使用することを推奨します
- Management API認証情報は安全に保管してください
