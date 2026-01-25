package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "kasir-api/docs"
	"kasir-api/internal/routes"
)

// @title           kasir API
// @version         1.0
// @description     API CRUD Category
// @host            localhost:8081
// @BasePath        /api/v1
func main() {
	// Swagger UI
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// API routes
	routes.SetupRoutes()

	// Root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// kalau bukan "/" → 404
		if r.URL.Path != "/" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   "Not Found",
				"message": "Endpoint tidak ditemukan",
			})
			return
		}

		// "/" → buka swagger
		http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
	})

	fmt.Println("Server running di 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println(err)
	}
}
