package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/internal/routes"
	"net/http"
)

func main() {
	// Check Server Running
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Server is running",
		})
	})

	routes.SetupRoutes()

	fmt.Println("Server running di 8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
