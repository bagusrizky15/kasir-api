package handlers

import (
	"encoding/json"
	"kasir-api/internal/repository"
	"net/http"
)

type ProductHandler struct {
	productRepo *repository.ProductRepository
}

func NewProductHandler(productRepo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		productRepo: productRepo,
	}
}

func (h *ProductHandler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	products := h.productRepo.GetAllProduct()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
