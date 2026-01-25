package repository

import "kasir-api/internal/models"

type CategoryRepository struct {
	categories []models.Category
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		categories: []models.Category{
			{ID: 1, Name: "Makanan", Description: "Makanan Instan"},
			{ID: 2, Name: "Minuman", Description: "Minuman Dingin"},
		},
	}
}

func (r *CategoryRepository) GetAll() []models.Category {
	return r.categories
}

func (r *CategoryRepository) GetByID(id int) (models.Category, bool) {
	for _, p := range r.categories {
		if p.ID == id {
			return p, true
		}
	}
	return models.Category{}, false
}

func (r *CategoryRepository) Create(category models.Category) models.Category {
	category.ID = len(r.categories) + 1
	r.categories = append(r.categories, category)
	return category
}

func (r *CategoryRepository) Update(id int, updated models.Category) (models.Category, bool) {
	for i, category := range r.categories {
		if category.ID == id {
			updated.ID = id
			r.categories[i] = updated
			return updated, true
		}
	}
	return models.Category{}, false
}

func (r *CategoryRepository) Delete(id int) bool {
	for i, category := range r.categories {
		if category.ID == id {
			r.categories = append(r.categories[:i], r.categories[i+1:]...)
			return true
		}
	}
	return false
}
