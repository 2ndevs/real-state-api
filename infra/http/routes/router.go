package routes

import (
	"main/infra/http/middlewares"
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
		router.Get("/highlights", controllers.GetHighlightedProperties)
	})

	router.Route("/admin", func(router chi.Router) {
		router.Use(middlewares.AuthMiddleware)

		router.Put("/refresh", controllers.RefreshToken)
		router.Route("/users", func(router chi.Router) {
			router.Post("/sign-in", controllers.SignIn)
			router.Post("/sign-up", controllers.SignUp)
		})

		router.Route("/kinds", func(router chi.Router) {
			router.Post("/", controllers.CreateKind)
			router.Put("/{id}", controllers.UpdateKind)
			router.Delete("/{id}", controllers.DeleteKind)
		})

		router.Route("/payment-types", func(router chi.Router) {
			router.Post("/", controllers.CreatePaymentType)
			router.Put("/{id}", controllers.UpdatePaymentType)
			router.Delete("/{id}", controllers.DeletePaymentType)
		})

		router.Route("/statuses", func(router chi.Router) {
			router.Post("/", controllers.CreateStatus)
			router.Put("/{id}", controllers.UpdateStatus)
			router.Delete("/{id}", controllers.DeleteStatus)
		})

		router.Route("/properties", func(router chi.Router) {
			router.Post("/", controllers.CreateProperty)
			router.Put("/{id}", controllers.UpdateProperty)
			router.Delete("/{id}", controllers.DeleteProperty)
		})

		router.Route("/roles", func(router chi.Router) {
			router.Post("/", controllers.CreateRole)
			// router.Get("/", controllers.GetManyRoles)
			// router.Get("/{id}", controllers.GetRole)
		})
	})
}
