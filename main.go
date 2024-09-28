package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"main/infra/http/middlewares"
	"main/infra/http/routes"
	"main/infra/repositories"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env variables")
	}

	router := chi.NewRouter()

	database.Connect()
	middlewares.Debug(router)
	routes.Handler(router)

	port := fmt.Sprintf(":%v", os.Getenv("APP_PORT"))

	fmt.Printf("[SERVER] Running on port %v\n", port)
	http.ListenAndServe(port, router)
}
