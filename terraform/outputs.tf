output "auth0_domain" {
  description = "Auth0 tenant domain"
  value       = var.auth0_domain
}

output "client_id" {
  description = "Auth0 Application Client ID"
  value       = auth0_client.app.client_id
}

output "client_secret" {
  description = "Auth0 Application Client Secret"
  value       = auth0_client.app.client_secret
  sensitive   = true
}

output "api_identifier" {
  description = "Auth0 API Identifier (Audience)"
  value       = auth0_resource_server.api.identifier
}

output "env_config" {
  description = "Environment configuration for .env files"
  value = <<-EOT

  Frontend (.env.local):
  AUTH0_SECRET=<generate-with-openssl-rand-hex-32>
  AUTH0_BASE_URL=${var.app_base_url}
  AUTH0_ISSUER_BASE_URL=https://${var.auth0_domain}
  AUTH0_CLIENT_ID=${auth0_client.app.client_id}
  AUTH0_CLIENT_SECRET=${auth0_client.app.client_secret}
  AUTH0_AUDIENCE=${auth0_resource_server.api.identifier}
  NEXT_PUBLIC_API_URL=http://localhost:8080

  Backend (.env):
  PORT=8080
  AUTH0_DOMAIN=${var.auth0_domain}
  AUTH0_AUDIENCE=${auth0_resource_server.api.identifier}
  CORS_ALLOWED_ORIGINS=${var.app_base_url}
  EOT
  sensitive = true
}
