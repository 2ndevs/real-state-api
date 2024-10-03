package middlewares

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type DatabaseContextKey string

const databaseContextKey DatabaseContextKey = "database"

func DatabaseMiddleware(database *gorm.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(writer http.ResponseWriter, request *http.Request) {
				ctx := context.WithValue(request.Context(), databaseContextKey, database)
				next.ServeHTTP(writer, request.WithContext(ctx))
			},
		)
	}
}

func GetDatabaseFromContext(request *http.Request) (*gorm.DB, error) {
	database, ok := request.Context().Value(databaseContextKey).(*gorm.DB)
	if !ok {
		return nil, errors.New("Connection with database wasn't found")
	}

	return database, nil
}
