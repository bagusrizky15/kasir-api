package routes

import (
	"database/sql"
	"net/http"

	"kasir-api/internal/handlers"
	"kasir-api/internal/repository"
	"kasir-api/internal/services"
)

func SetupRoutes(mux *http.ServeMux, db *sql.DB) {

	// ===== PRODUCT =====
	productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// ===== CATEGORY =====
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// ===== TRANSACTIONS =====
	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	reportRepo := repository.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handlers.GetTodaySalesReport(reportService)

	// ===== PRODUCT ROUTES =====
	mux.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetProducts(w, r)
		case http.MethodPost:
			productHandler.CreateProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetProductByID(w, r)
		case http.MethodPut:
			productHandler.UpdateProductByID(w, r)
		case http.MethodDelete:
			productHandler.DeleteProductByID(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// ===== CATEGORY ROUTES =====
	mux.HandleFunc("/api/v1/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetCategories(w, r)
		case http.MethodPost:
			categoryHandler.CreateCategory(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetCategoryByID(w, r)
		case http.MethodPut:
			categoryHandler.UpdateCategoryByID(w, r)
		case http.MethodDelete:
			categoryHandler.DeleteCategoryByID(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// ===== TRANSACTION ROUTES =====
	mux.HandleFunc("/api/v1/checkout", func(w http.ResponseWriter, r *http.Request) {
		transactionHandler.HandleCheckout(w, r)
	})

	// ===== REPORT ROUTES =====
	mux.HandleFunc("/api/v1/report/today", func(w http.ResponseWriter, r *http.Request) {
		reportHandler(w, r)
	})
}
