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

// GetProducts godoc
// @Summary      Get all products
// @Description  Ambil semua data product
// @Tags         Products
// @Produce      json
// @Success      200 {array} models.Product
// @Router       /products [get]
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.productRepo.GetAll())
}

// CreateProduct godoc
// @Summary      Create new product
// @Description  Tambah product baru
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product body models.Product true "Create product payload"
// @Success      201 {object} models.Product
// @Failure      400 {object} map[string]string
// @Router       /products [post]
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

// GetProductByID godoc
// @Summary      Get product by ID
// @Description  Ambil detail product berdasarkan ID
// @Tags         Products
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  models.Product
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /products/{id} [get]
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

// UpdateProductByID godoc
// @Summary      Update product
// @Description  Update data product berdasarkan ID
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id      path int true "Product ID"
// @Param        product body models.Product true "Update product payload"
// @Success      200 {object} models.Product
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /products/{id} [put]
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

// DeleteProductByID godoc
// @Summary      Delete product
// @Description  Hapus product berdasarkan ID
// @Tags         Products
// @Produce      json
// @Param        id path int true "Product ID"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /products/{id} [delete]
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
