variable "auth0_domain" {
  description = "Auth0 tenant domain"
  type        = string
  default     = "dev-31tn6c1v6nf53cr3.us.auth0.com"
}

variable "auth0_management_client_id" {
  description = "Auth0 Management API Client ID"
  type        = string
  sensitive   = true
}

variable "auth0_management_client_secret" {
  description = "Auth0 Management API Client Secret"
  type        = string
  sensitive   = true
}

variable "app_base_url" {
  description = "Application base URL"
  type        = string
  default     = "http://localhost:3000"
}

variable "api_identifier" {
  description = "Auth0 API identifier (audience)"
  type        = string
  default     = "https://api.auth0-sandbox.local"
}
