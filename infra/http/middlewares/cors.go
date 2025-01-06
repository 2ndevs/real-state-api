package middlewares

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Cors(router chi.Router) {
	adminOrigin := os.Getenv("ADMIN_ORIGIN")
	webOrigin := os.Getenv("WEB_ORIGIN")

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{adminOrigin, webOrigin},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "x-refresh-token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}
