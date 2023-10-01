package handler

import (
	"github.com/go-chi/chi/v5"
	"github/ahmedghazey/packaging/internal/http/rest"
	"github/ahmedghazey/packaging/internal/middleware"
	"github/ahmedghazey/packaging/internal/service"
	"net/http"
)

func Handler(packagingService service.PackageService) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recovery)
	router.Get("/health", rest.Health())
	router.Post("/add-packages", rest.AddPackages(packagingService))
	router.Post("/calculate-packages", rest.CalculatePackages(packagingService))
	return router
}
