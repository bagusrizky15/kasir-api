package handlers

import (
	"database/sql"
	"encoding/json"
	"kasir-api/internal/models"
	"kasir-api/internal/services"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

// GetProducts godoc
// @Summary      Get all products
// @Description  Ambil semua data product
// @Tags         Products
// @Produce      json
// @Success      200 {array} models.Product
// @Failure      500 {object} map[string]string
// @Router       /products [get]
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	products, err := h.service.GetAll(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
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
// @Failure      500 {object} map[string]string
// @Router       /products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload models.Product
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product, err := h.service.Create(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
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
// @Failure      500  {object}  map[string]string
// @Router       /products/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := getProductId(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
// @Failure      500 {object} map[string]string
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

	updated, err := h.service.Update(id, payload)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
// @Failure      500 {object} map[string]string
// @Router       /products/{id} [delete]
func (h *ProductHandler) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := getProductId(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(id); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product deleted successfully",
	})
}

// helper
func getProductId(path string) (int, error) {
	idStr := strings.TrimPrefix(path, "/api/v1/products/")
	return strconv.Atoi(idStr)
}
