package main

import (
	"fmt"
	"main/http/middlewares"
	"main/http/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	middlewares.Debug(router)
	routes.Handler(router)

	fmt.Println("[INFO] Running on port 3333")
	http.ListenAndServe(":3333", router)
}
