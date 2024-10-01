package middlewares

import (
	"context"
	"net/http"

	"gorm.io/gorm"
)

type contextKey string

const dbContextKey = contextKey("db")

func DatabaseMiddleware(db *gorm.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
			appContext := context.WithValue(request.Context(), dbContextKey, db)
			next.ServeHTTP(write, request.WithContext(appContext))
		})
	}
}

func GetDBFromContext(appContext context.Context) *gorm.DB {
	db, ok := appContext.Value(dbContextKey).(*gorm.DB)
	if !ok {
		return nil
	}
	return db
}
