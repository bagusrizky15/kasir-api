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
