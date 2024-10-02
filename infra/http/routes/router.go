package routes

import (
	"main/infra/http/routes/internals/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Handler(router chi.Router) {
	router.Get(
		"/",
		func(writer http.ResponseWriter, request *http.Request) {
			message := "Hello world"
			writer.Write([]byte(message))
		},
	)

	router.Post("/status", controllers.CreateStatus)
	router.Post("/kinds", controllers.CreateKind)
	router.Post("/payment-types", controllers.CreatePaymentType)
	router.Post("/properties", controllers.CreateProperty)

	router.Get("/kinds", controllers.GetKinds)
	router.Get("/kinds/{id}", controllers.GetKind)
}
