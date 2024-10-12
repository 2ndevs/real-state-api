package routes

import (
	"main/infra/http/routes/internals/controllers"

	"github.com/go-chi/chi/v5"
)

func Handler(appRouter chi.Router) {

	appRouter.Route("/web", func(webRouter chi.Router) {
		webRouter.Route("/kinds", func(router chi.Router) {
			router.Post("/", controllers.CreateKind)
			router.Get("/", controllers.GetManyKinds)
			router.Get("/{id}", controllers.GetKind)
			router.Put("/{id}", controllers.UpdateKind)
			router.Delete("/{id}", controllers.DeleteKind)
		})

		webRouter.Route("/payment-types", func(router chi.Router) {
			router.Post("/", controllers.CreatePaymentType)
			router.Get("/", controllers.GetManyPaymentTypes)
			router.Get("/{id}", controllers.GetPaymentType)
			router.Put("/{id}", controllers.UpdatePaymentType)
			router.Delete("/{id}", controllers.DeletePaymentType)
		})

		webRouter.Route("/statuses", func(router chi.Router) {
			router.Post("/", controllers.CreateStatus)
			router.Get("/", controllers.GetManyStatuses)
			router.Get("/{id}", controllers.GetStatus)
			router.Put("/{id}", controllers.UpdateStatus)
			router.Delete("/{id}", controllers.DeleteStatus)
		})

		webRouter.Route("/properties", func(router chi.Router) {
			router.Post("/", controllers.CreateProperty)
			router.Get("/", controllers.GetManyProperties)
			router.Get("/{id}", controllers.GetProperty)
			router.Put("/{id}", controllers.UpdateProperty)
			router.Delete("/{id}", controllers.DeleteProperty)
		})
	})

	appRouter.Route("/admin", func(router chi.Router) {
		// ADMIN ROUTES
	})
}
