package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/naoyafurudono/auth0-sandbox/backend/internal/config"
	"github.com/naoyafurudono/auth0-sandbox/backend/internal/generated"
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

	// 生成されたHandlerを使用してルーティングを設定
	apiHandler := generated.Handler(userHandler)

	// ミドルウェアを適用
	handler := corsMiddleware.Handler(authMiddleware.ValidateJWT(apiHandler))

	log.Printf("Starting server on port %s", cfg.Port)
	log.Printf("Auth0 Domain: %s", cfg.Auth0Domain)
	log.Printf("Auth0 Audience: %s", cfg.Auth0Audience)
	log.Printf("CORS Allowed Origins: %s", cfg.CORSAllowedOrigins)

	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
