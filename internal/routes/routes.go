package routes

import (
	"kasir-api/internal/handlers"
	"kasir-api/internal/repository"
	"net/http"
)

func SetupRoutes() {

	productRepo := repository.NewProductRepository()
	productHandler := handlers.NewProductHandler(productRepo)

	categoryRepo := repository.NewCategoryRepository()
	categoryHandler := handlers.NewCategoryHandler(categoryRepo)

	http.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetProducts(w, r)
		case http.MethodPost:
			productHandler.CreateProduct(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/v1/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetProductByID(w, r)
		case http.MethodPut:
			productHandler.UpdateProductByID(w, r)
		case http.MethodDelete:
			productHandler.DeleteProductByID(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/v1/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetCategories(w, r)
		case http.MethodPost:
			categoryHandler.CreateCategory(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/v1/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetCategoryByID(w, r)
		case http.MethodPut:
			categoryHandler.UpdateCategoryByID(w, r)
		case http.MethodDelete:
			categoryHandler.DeleteCategoryByID(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
