package services

import (
	"database/sql"
	"kasir-api/internal/models"
	"kasir-api/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

// Get all categories
func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

// Get category by ID
func (s *CategoryService) GetByID(id int) (models.Category, error) {
	category, err := s.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Category{}, err
		}
		return models.Category{}, err
	}
	return category, nil
}

// Create new category
func (s *CategoryService) Create(category models.Category) (models.Category, error) {
	// validasi sederhana
	if category.Name == "" {
		return models.Category{}, sql.ErrNoRows
	}

	return s.repo.Create(category)
}

// Update category
func (s *CategoryService) Update(id int, category models.Category) (models.Category, error) {
	if category.Name == "" {
		return models.Category{}, sql.ErrNoRows
	}

	return s.repo.Update(id, category)
}

// Delete category
func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
