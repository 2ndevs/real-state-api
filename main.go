package main

import (
	"main/http/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	routes.Handler(router)

	http.ListenAndServe(":3333", router)
}
