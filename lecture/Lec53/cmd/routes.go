package cmd

import (
	"net/http"

	"Lec53/handlers"

	"Lec53/middleware"
)

func initRoutes(router *http.ServeMux, manager *middleware.Manager) {
	router.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)
	router.Handle("POST /create-products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)
	router.Handle("GET /products/{Id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductById),
		),
	)
}
