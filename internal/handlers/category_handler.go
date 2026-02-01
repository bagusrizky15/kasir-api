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

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

// GetCategories godoc
// @Summary      Get all categories
// @Description  Ambil semua data category
// @Tags         Categories
// @Produce      json
// @Success      200 {array} models.Category
// @Failure      500 {object} map[string]string
// @Router       /categories [get]
func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

// CreateCategory godoc
// @Summary      Create new category
// @Description  Tambah category baru
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        category body models.Category true "Create category payload"
// @Success      201 {object} models.Category
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /categories [post]
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var payload models.Category
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	category, err := h.service.Create(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

func getCategoryId(path string) (int, error) {
	idStr := strings.TrimPrefix(path, "/api/v1/categories/")
	return strconv.Atoi(idStr)
}

// GetCategoryByID godoc
// @Summary      Get category by ID
// @Description  Ambil detail category berdasarkan ID
// @Tags         Categories
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Success      200  {object}  models.Category
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := getCategoryId(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

// UpdateCategoryByID godoc
// @Summary      Update category
// @Description  Update data category berdasarkan ID
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        id       path int true "Category ID"
// @Param        category body models.Category true "Update category payload"
// @Success      200 {object} models.Category
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /categories/{id} [put]
func (h *CategoryHandler) UpdateCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := getCategoryId(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var payload models.Category
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedCategory, err := h.service.Update(id, payload)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCategory)
}

// DeleteCategoryByID godoc
// @Summary      Delete category
// @Description  Hapus category berdasarkan ID
// @Tags         Categories
// @Produce      json
// @Param        id path int true "Category ID"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := getCategoryId(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(id); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Category not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Category deleted successfully",
	})
}
