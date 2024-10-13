package middlewares

import (
	"errors"
	"main/core"
	"main/utils/libs"
	"net/http"
	"slices"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			signPaths := []string{"/admin/users/sign-up", "/admin/users/sign-in"}
			path := request.URL.Path

			parser := libs.JWT{}

			var token string
			bearer := request.Header.Get("Authorization")
			tokenArr := strings.Split(bearer, " ")

			if len(tokenArr) > 1 {
				token = tokenArr[1]
			}

			if len(token) == 0 && !slices.Contains(signPaths, path) {
				core.HandleHTTPStatus(writer, core.MissingAuthorizationTokenError)
				return
			}

			_, err := parser.Parse(token)

			if err != nil && !slices.Contains(signPaths, path) {
				switch {
				case errors.Is(err, jwt.ErrTokenExpired):
					{
						if path == "/admin/refresh" {
							break
						}

						core.HandleHTTPStatus(writer, core.AuthorizationTokenExpiredError)
						return
					}

				default:
					{
						core.HandleHTTPStatus(writer, core.MissingAuthorizationTokenError)
						return
					}
				}
			}

			next.ServeHTTP(writer, request)
		},
	)
}
