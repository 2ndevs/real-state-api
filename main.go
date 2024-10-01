package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"main/infra/http/middlewares"
	"main/infra/http/routes"
	database "main/infra/repositories"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env variables")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	router := chi.NewRouter()

	middlewares.Debug(router)

	router.Use(middlewares.DatabaseMiddleware(db))

	routes.Handler(router)

	port := fmt.Sprintf(":%v", os.Getenv("APP_PORT"))

	fmt.Printf("[SERVER] Running on port %v\n", port)
	serverError := http.ListenAndServe(port, router)

	if serverError != nil {
		log.Fatal("Unable to run server")
	}

}
