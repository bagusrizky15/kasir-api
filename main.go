package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "kasir-api/docs"
	"kasir-api/internal/database"
	"kasir-api/internal/middleware"
	"kasir-api/internal/routes"
)

// @title           kasir API
// @version         1.0
// @description     API CRUD Category & Product
// @host            localhost:8081
// @BasePath        /api/v1
// @schemes         http
type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	// ===== VIPER SETUP =====
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// load .env if exists
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Error reading .env file:", err)
		}
	}

	cfg := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	if cfg.Port == "" {
		cfg.Port = "8081"
	}

	// ===== DATABASE =====
	db, err := database.InitDB(cfg.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// ===== ROUTER =====
	mux := http.NewServeMux()

	// swagger
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// api routes (inject DB)
	routes.SetupRoutes(mux, db)

	// root handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   "Not Found",
				"message": "Endpoint tidak ditemukan",
			})
			return
		}

		http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
	})

	// ===== MIDDLEWARE =====
	handler := middleware.EnableCORS(mux)

	// ===== SERVER =====
	log.Printf("Server running on port %s\n", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Println("server stopped:", err)
	}

}
