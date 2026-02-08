package handlers

import (
	"encoding/json"
	"kasir-api/internal/services"
	"net/http"
)

type ReportHandler struct {
	repo *services.ReportService
}

// GetTodaySalesReport godoc
// @Summary      Sales report hari ini
// @Description  Menampilkan total revenue, total transaksi, dan produk terlaris hari ini
// @Tags         Reports
// @Accept       json
// @Produce      json
// @Success      200 {object} models.SalesReport
// @Failure      405 {object} map[string]string "Method not allowed"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /report/today [get]
func GetTodaySalesReport(service *services.ReportService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		report, err := service.GetTodaySalesReport()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(report)
	}
}
