package handlers

import (
	"encoding/json"
	"kasir-api/internal/models"
	"kasir-api/internal/repository"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	categoryRepo *repository.CategoryRepository
}

func NewCategoryHandler(categoryRepo *repository.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{
		categoryRepo: categoryRepo,
	}
}

// GetCategories godoc
// @Summary      Get all categories
// @Description  Ambil semua data category
// @Tags         Categories
// @Produce      json
// @Success      200 {array} models.Category
// @Router       /categories [get]
func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.categoryRepo.GetAll())
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
// @Router       /categories [post]
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var payload models.Category
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	category := h.categoryRepo.Create(payload)
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
// @Router       /categories/{id} [get]
func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := getCategoryId(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	category, found := h.categoryRepo.GetByID(id)
	if !found {
		http.Error(w, "Category not found", http.StatusNotFound)
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

	updatedCategory, found := h.categoryRepo.Update(id, payload)
	if !found {
		http.Error(w, "Category not found", http.StatusNotFound)
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
// @Router       /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := getCategoryId(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid category ID",
		})
		return
	}

	if !h.categoryRepo.Delete(id) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Category not found",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Category deleted successfully",
	})
}
