package repository

import "kasir-api/internal/models"

type ProductRepository struct {
	products []models.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: []models.Product{
			{ID: 1, Name: "Indomie Goreng", Price: 3000, Stock: 10},
			{ID: 2, Name: "Indomie Ayam Bawang", Price: 3000, Stock: 10},
		},
	}
}

func (r *ProductRepository) GetAll() []models.Product {
	return r.products
}

func (r *ProductRepository) GetByID(id int) (models.Product, bool) {
	for _, p := range r.products {
		if p.ID == id {
			return p, true
		}
	}
	return models.Product{}, false
}

func (r *ProductRepository) Create(product models.Product) models.Product {
	product.ID = len(r.products) + 1
	r.products = append(r.products, product)
	return product
}

func (r *ProductRepository) Update(id int, updated models.Product) (models.Product, bool) {
	for i, p := range r.products {
		if p.ID == id {
			updated.ID = id
			r.products[i] = updated
			return updated, true
		}
	}
	return models.Product{}, false
}

func (r *ProductRepository) Delete(id int) bool {
	for i, p := range r.products {
		if p.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return true
		}
	}
	return false
}
