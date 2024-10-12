package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	"main/core"
	"main/utils/libs"
	"net/http"
	"slices"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			signPaths := []string{"/admin/users/sign-up", "/admin/users/sign-in"}
			parser := libs.JWT{}

			token := strings.Split(request.Header.Get("Authorization"), "Bearer ")[1]
			path := request.URL.Path

			if len(token) <= 0 && !slices.Contains(signPaths, path) {
				core.HandleHTTPStatus(writer, core.MissingAuthorizationTokenError)
				return
			}

			_, err := parser.Parse(token)

			if err != nil && !slices.Contains(signPaths, path) {
				switch err.(error) {
				case jwt.ErrTokenExpired:
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
