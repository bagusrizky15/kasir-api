package handlers

import (
	"encoding/json"
	"net/http"

	"kasir-api/internal/models"
	"kasir-api/internal/services"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Checkout godoc
// @Summary      Create checkout
// @Description  Melakukan checkout dan membuat transaksi baru
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body models.CheckoutRequest true "Checkout items"
// @Success      200 {object} models.Transaction
// @Failure      400 {object} map[string]string "Invalid request body"
// @Failure      500 {object} map[string]string "Internal server error"
// @Router       /checkout [post]
func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
