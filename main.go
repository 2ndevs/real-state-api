package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"main/database"
	"main/http/middlewares"
	"main/http/routes"

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

	fmt.Printf("[SERVER] Running on port %v", port)
  http.ListenAndServe(port, router)
}
