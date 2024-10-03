package middlewares

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ValidatorKey string

const validatorKey ValidatorKey = "validator"

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

func ValidatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(write http.ResponseWriter, request *http.Request) {
			ctx := context.WithValue(request.Context(), validatorKey, validate)

			next.ServeHTTP(write, request.WithContext(ctx))
		},
	)
}

func GetValidator(request *http.Request) (*validator.Validate, error) {
	validate, ok := request.Context().Value(validatorKey).(*validator.Validate)
	if !ok {
		return nil, errors.New("Unable to retrieve validator from context")
	}

	return validate, nil
}
