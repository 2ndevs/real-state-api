package middlewares

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Debug(router chi.Router) {
	router.Use(middleware.Logger)
}
