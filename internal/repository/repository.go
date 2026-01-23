package repository

import "kasir-api/internal/models"

type ProductRepository struct {
	products []models.Product
}

type CategoryRepository struct {
	categories []models.Category
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: []models.Product{
			{ID: 1, Name: "Indomie Goreng", Price: 3000, Stok: 10},
			{ID: 2, Name: "Indomie Ayam Bawang", Price: 3000, Stok: 10},
			{ID: 3, Name: "Indomie Kari Ayam", Price: 3000, Stok: 10},
			{ID: 4, Name: "Indomie Ayam Geprek", Price: 3000, Stok: 10},
			{ID: 5, Name: "Indomie Kari Special", Price: 3000, Stok: 10},
		},
	}
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		categories: []models.Category{
			{ID: 1, Name: "Makanan", Description: "Makanan Instan"},
			{ID: 2, Name: "Minuman", Description: "Minuman Dingin"},
		},
	}
}
