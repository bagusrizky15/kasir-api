package services

import (
	"database/sql"
	"kasir-api/internal/models"
	"kasir-api/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

// Get all products
func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

// Get product by ID
func (s *ProductService) GetByID(id int) (models.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Product{}, err
		}
		return models.Product{}, err
	}
	return product, nil
}

// Create new product
func (s *ProductService) Create(product models.Product) (models.Product, error) {
	// (opsional) validasi sederhana
	if product.Name == "" {
		return models.Product{}, sql.ErrNoRows
	}

	return s.repo.Create(product)
}

// Update product
func (s *ProductService) Update(id int, product models.Product) (models.Product, error) {
	// (opsional) validasi
	if product.Name == "" {
		return models.Product{}, sql.ErrNoRows
	}

	return s.repo.Update(id, product)
}

// Delete product
func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
