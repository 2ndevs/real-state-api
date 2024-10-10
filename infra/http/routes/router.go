package routes

import (
	"main/infra/http/routes/internals/controllers"

	"github.com/go-chi/chi/v5"
)

func Handler(router chi.Router) {

	router.Route("/web", func(router chi.Router) {
		router.Route("/kinds", func(router chi.Router) {
			router.Post("/", controllers.CreateKind)
			router.Get("/", controllers.GetManyKinds)
			router.Get("/{id}", controllers.GetKind)
		})

		router.Route("/payment-types", func(router chi.Router) {
			router.Post("/", controllers.CreatePaymentType)
			router.Get("/", controllers.GetManyPaymentTypes)
			router.Get("/{id}", controllers.GetPaymentType)
		})

		router.Route("/statuses", func(router chi.Router) {
			router.Post("/", controllers.CreateStatus)
			router.Get("/", controllers.GetManyStatuses)
			router.Get("/{id}", controllers.GetStatus)
		})

		router.Route("/properties", func(router chi.Router) {
			router.Post("/", controllers.CreateProperty)
			router.Get("/", controllers.GetManyProperties)
			router.Get("/{id}", controllers.GetProperty)
		})
	})

	router.Route("/admin", func(router chi.Router) {
		router.Route("/users", func(router chi.Router) {
			// router.Post("/sign-in", controllers.SignIn)
      // router.Post("/sign-up", controllers.SignUp) // FIXME: remove, it's likely that it's not even an option
		})

		router.Route("/roles", func(router chi.Router) {
			// router.Post("/", controllers.CreateRole)
			// router.Get("/", controllers.GetManyRoles)
			// router.Get("/{id}", controllers.GetRole)
		})
		// ADMIN ROUTES
	})
}
