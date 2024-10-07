package routes

import (
	"main/infra/http/routes/internals/controllers"

	"github.com/go-chi/chi/v5"
)

func Handler(router chi.Router) {
	router.Post("/kinds", controllers.CreateKind)
	router.Get("/kinds", controllers.GetManyKinds)
	router.Get("/kinds/{id}", controllers.GetKind)

	router.Post("/payment-types", controllers.CreatePaymentType)
	router.Get("/payment-types", controllers.GetManyPaymentTypes)
	router.Get("/payment-types/{id}", controllers.GetPaymentType)

	router.Post("/statuses", controllers.CreateStatus)
	router.Get("/statuses", controllers.GetManyStatuses)
	router.Get("/statuses/{id}", controllers.GetStatus)

	router.Post("/properties", controllers.CreateProperty)
	router.Get("/properties", controllers.GetManyProperties)
	router.Get("/properties/{id}", controllers.GetProperty)
}
