package routes

import (
	"main/infra/http/routes/internals/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Handler(router chi.Router, db *gorm.DB) {
	router.Get(
		"/",
		func(writer http.ResponseWriter, request *http.Request) {
			message := "Hello world"
			writer.Write([]byte(message))
		},
	)

	router.Post("/status", controllers.CreateStatus)

	// KIND
	kindController := &controllers.KindController{DB: db}
	router.Post("/kind", kindController.CreateKind)
}
