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

	database, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	router := chi.NewRouter()

	middlewares.Debug(router)
	router.Use(middlewares.ValidatorMiddleware)
	router.Use(middlewares.DatabaseMiddleware(database))

	router.Use()

	routes.Handler(router)

	port := fmt.Sprintf(":%v", os.Getenv("APP_PORT"))

	fmt.Printf("[SERVER] Running on port %v\n", port)
	serverError := http.ListenAndServe(port, router)

	if serverError != nil {
		log.Fatal("Unable to run server")
	}

}
