package handlers

import (
	"encoding/json"
	"kasir-api/internal/models"
	"kasir-api/internal/repository"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	productRepo *repository.ProductRepository
}

func NewProductHandler(productRepo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{
		productRepo: productRepo,
	}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.productRepo.GetAll())
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var payload models.Product
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product := h.productRepo.Create(payload)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func getProductId(path string) (int, error) {
	idStr := strings.TrimPrefix(path, "/api/v1/products/")
	return strconv.Atoi(idStr)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := getProductId(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, found := h.productRepo.GetByID(id)
	if !found {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := getProductId(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var payload models.Product
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updated, found := h.productRepo.Update(id, payload)
	if !found {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (h *ProductHandler) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := getProductId(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid product ID",
		})
		return
	}

	if !h.productRepo.Delete(id) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Product not found",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product deleted successfully",
	})
}
