# Auth0 API (Resource Server)
resource "auth0_resource_server" "api" {
  name       = "Auth0 Sandbox API"
  identifier = var.api_identifier

  signing_alg = "RS256"

  allow_offline_access                            = false
  token_lifetime                                  = 86400
  skip_consent_for_verifiable_first_party_clients = true
}

# Auth0 Application (Client)
resource "auth0_client" "app" {
  name        = "Auth0 Sandbox Frontend"
  description = "Next.js frontend application for Auth0 sandbox"
  app_type    = "regular_web"

  callbacks = [
    "${var.app_base_url}/api/auth/callback"
  ]

  allowed_logout_urls = [
    var.app_base_url
  ]

  web_origins = [
    var.app_base_url
  ]

  allowed_origins = [
    var.app_base_url
  ]

  oidc_conformant = true

  grant_types = [
    "authorization_code",
    "refresh_token"
  ]

  jwt_configuration {
    alg = "RS256"
  }
}

# Auth0 Application - API 連携
resource "auth0_client_grant" "app_api_grant" {
  client_id = auth0_client.app.id
  audience  = auth0_resource_server.api.identifier

  scopes = []
}

# Connection (Username-Password-Authentication) を有効化
data "auth0_connection" "username_password" {
  name = "Username-Password-Authentication"
}

resource "auth0_connection_client" "app_connection" {
  connection_id = data.auth0_connection.username_password.id
  client_id     = auth0_client.app.id
}
