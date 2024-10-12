package controllers

import (
	"log"
	"main/core"
	"main/utils/libs"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func RefreshToken(writer http.ResponseWriter, request *http.Request) {
	parser := libs.JWT{}
	token := request.Header.Get("X-Refresh-Token")

	if len(token) <= 0 {
		core.HandleHTTPStatus(writer, core.MissingRefreshTokenError)
		return
	}

	oldToken, err := parser.Parse(token)
	if err != nil {
		switch err.(error) {
		case jwt.ErrTokenExpired:
			{
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

	sub, err := oldToken.Claims.GetSubject()
	if err != nil {
		core.HandleHTTPStatus(writer, err)
		return
	}
	log.Println("SUB: ", sub)

	role := oldToken.Raw
	log.Println("RAW: ", role)

	// refreshToken, err := parser.Generate(libs.CreateJWTParams{
	// 	Sub:  1,
	// 	Role: 2,
	// })
}
