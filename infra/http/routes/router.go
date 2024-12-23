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

	router.Route("/negotiation-types", func(router chi.Router) {
		router.Get("/", controllers.GetManyNegotiationTypes)
		router.Get("/{id}", controllers.GetNegotiationType)
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

	router.Get("/topics", controllers.GetTopics)

	router.Route("/interested-users", func(router chi.Router) {
		router.Post("/", controllers.CreateInterestedUser)
	})

	router.Route("/admin", func(router chi.Router) {
		router.Use(middlewares.AuthMiddleware)

		router.Put("/refresh", controllers.RefreshToken)
		router.Route("/users", func(router chi.Router) {
			router.Get("/", controllers.GetManyUsers)

			router.Post("/sign-in", controllers.SignIn)
			router.Post("/sign-up", controllers.SignUp)

			router.Get("/{id}", controllers.GetUser)
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

		router.Route("/negotiation-types", func(router chi.Router) {
			router.Post("/", controllers.CreateNegotiationType)
			router.Put("/{id}", controllers.UpdateNegotiationType)
			router.Delete("/{id}", controllers.DeleteNegotiationType)
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
			router.Get("/", controllers.GetManyRoles)
			router.Get("/{id}", controllers.GetRole)
			router.Put("/{id}", controllers.UpdateRole)
			router.Delete("/{id}", controllers.DeleteRole)
		})

		router.Route("/measurement-unit", func(router chi.Router) {
			router.Get("/", controllers.GetManyMeasurementUnits)
			router.Get("/{id}", controllers.GetMeasurementUnit)
			router.Put("/{id}", controllers.UpdateMeasurementUnit)
			router.Post("/", controllers.CreateMeasurementUnit)
			router.Delete("/{id}", controllers.DeleteMeasurementUnit)
		})

		router.Route("/interested-users", func(router chi.Router) {
			router.Get("/", controllers.GetManyInterestedUsers)
			router.Get("/{id}", controllers.GetInterestedUser)
			router.Delete("/{id}", controllers.DeleteInterestedUser)
			router.Put("/{id}", controllers.UpdateInterestedUser)
		})
	})
}
