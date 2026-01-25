package routes

import (
	"kasir-api/internal/handlers"
	"kasir-api/internal/repository"
	"net/http"
)

func SetupRoutes() {
	productRepo := repository.NewProductRepository()
	productHandler := handlers.NewProductHandler(productRepo)

	http.HandleFunc("/api/v1/products", productHandler.GetAllProduct)
}
