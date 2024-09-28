package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Handler(router chi.Router) {
	router.Get(
		"/",
		func(writer http.ResponseWriter, request *http.Request) {
			message := "Hello world"
			writer.Write([]byte(message))
		},
	)
}
