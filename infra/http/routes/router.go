package routes

import (
	"main/infra/http/routes/internals/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Handler(router chi.Router) {
	router.Get("/ping", func(write http.ResponseWriter, request *http.Request) {
		write.Write([]byte("pong"))
	})

	router.Route("/kinds", func(router chi.Router) {
		router.Get("/", controllers.GetManyKinds)
		router.Get("/{id}", controllers.GetKind)
	})

	router.Route("/payment-types", func(router chi.Router) {
		router.Get("/", controllers.GetManyPaymentTypes)
		router.Get("/{id}", controllers.GetPaymentType)
	})

	router.Route("/statuses", func(router chi.Router) {
		router.Get("/", controllers.GetManyStatuses)
		router.Get("/{id}", controllers.GetStatus)
	})

	router.Route("/properties", func(router chi.Router) {
		router.Get("/", controllers.GetManyProperties)
		router.Get("/{id}", controllers.GetProperty)
	})

	router.Route("/admin", func(router chi.Router) {
		router.Post("/kinds", controllers.CreateKind)
		router.Post("/statuses", controllers.CreateStatus)
		router.Post("/properties", controllers.CreateProperty)
		router.Post("/payment-types", controllers.CreatePaymentType)

		router.Route("/users", func(router chi.Router) {
			// router.Post("/sign-in", controllers.SignIn)
			// router.Post("/sign-up", controllers.SignUp) // FIXME: remove, it's likely that it's not even an option
		})

		router.Route("/roles", func(router chi.Router) {
			router.Post("/", controllers.CreateRole)
			// router.Get("/", controllers.GetManyRoles)
			// router.Get("/{id}", controllers.GetRole)
		})
	})
}
