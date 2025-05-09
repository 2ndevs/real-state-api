package main

import (
	"fmt"
	"log"
	"main/infra/http/middlewares"
	"main/infra/http/routes"
	"net/http"
	"os"

	database "main/infra/repositories"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env variables")
	}

	database, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	router := chi.NewRouter()

	middlewares.Cors(router)

	middlewares.Debug(router)

	router.Use(middlewares.S3Middleware)
	router.Use(middlewares.ValidatorMiddleware)
	router.Use(middlewares.DatabaseMiddleware(database))

	routes.Handler(router)

	port := fmt.Sprintf(":%v", os.Getenv("APP_PORT"))

	fmt.Printf("[SERVER] Running on port %v\n", port)
	serverError := http.ListenAndServe(port, router)
	if serverError != nil {
		log.Printf("%v\n", serverError)
		log.Fatal("Unable to run server")
	}
}
