package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/naoyafurudono/auth0-sandbox/backend/internal/config"
	"github.com/naoyafurudono/auth0-sandbox/backend/internal/handler"
	"github.com/naoyafurudono/auth0-sandbox/backend/internal/middleware"
	"github.com/naoyafurudono/auth0-sandbox/backend/internal/model"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	if cfg.Auth0Domain == "" || cfg.Auth0Audience == "" {
		log.Fatal("AUTH0_DOMAIN and AUTH0_AUDIENCE must be set")
	}

	authMiddleware, err := middleware.NewAuthMiddleware(cfg.Auth0Domain, cfg.Auth0Audience)
	if err != nil {
		log.Fatalf("Failed to create auth middleware: %v", err)
	}

	corsMiddleware := middleware.NewCORSMiddleware(cfg.CORSAllowedOrigins)

	store := model.NewStore()
	userHandler := handler.NewUserHandler(store)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/users/me", userHandler.GetCurrentUser)
	mux.HandleFunc("/api/v1/users/me/profile", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetUserProfile(w, r)
		case http.MethodPut:
			userHandler.UpdateUserProfile(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/v1/users/me/data", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetUserData(w, r)
		case http.MethodPost:
			userHandler.CreateUserData(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	handler := corsMiddleware.Handler(authMiddleware.ValidateJWT(mux))

	log.Printf("Starting server on port %s", cfg.Port)
	log.Printf("Auth0 Domain: %s", cfg.Auth0Domain)
	log.Printf("Auth0 Audience: %s", cfg.Auth0Audience)
	log.Printf("CORS Allowed Origins: %s", cfg.CORSAllowedOrigins)

	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
