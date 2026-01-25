package repository

import "kasir-api/internal/models"

type ProductRepository struct {
	products []models.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: []models.Product{
			{ID: 1, Name: "Indomie Goreng", Price: 3000, Stock: 10, CategoryID: 1},
			{ID: 2, Name: "Indomie Ayam Bawang", Price: 3000, Stock: 10, CategoryID: 1},
		},
	}
}

func (r *ProductRepository) GetAllProduct() []models.Product {
	return r.products
}
