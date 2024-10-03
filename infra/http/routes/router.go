package routes

import (
	"main/infra/http/routes/internals/controllers"

	"github.com/go-chi/chi/v5"
)

func Handler(router chi.Router) {
	router.Post("/status", controllers.CreateStatus)
	router.Post("/kinds", controllers.CreateKind)
	router.Post("/payment-types", controllers.CreatePaymentType)
	router.Post("/properties", controllers.CreateProperty)

	router.Get("/kinds", controllers.GetManyKinds)
	router.Get("/kinds/{id}", controllers.GetKind)
}
