package repository

import (
	"database/sql"
	"kasir-api/internal/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

// ===== GET ALL =====
func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	rows, err := r.db.Query(`
		SELECT id, name, description
		FROM categories
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var c models.Category
		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Description,
		); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

// ===== GET BY ID =====
func (r *CategoryRepository) GetByID(id int) (models.Category, error) {
	var c models.Category

	err := r.db.QueryRow(`
		SELECT id, name, description
		FROM categories
		WHERE id = $1
	`, id).Scan(
		&c.ID,
		&c.Name,
		&c.Description,
	)

	if err != nil {
		return models.Category{}, err
	}

	return c, nil
}

// ===== CREATE =====
func (r *CategoryRepository) Create(category models.Category) (models.Category, error) {
	err := r.db.QueryRow(`
		INSERT INTO categories (name, description)
		VALUES ($1, $2)
		RETURNING id
	`,
		category.Name,
		category.Description,
	).Scan(&category.ID)

	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

// ===== UPDATE =====
func (r *CategoryRepository) Update(id int, updated models.Category) (models.Category, error) {
	err := r.db.QueryRow(`
		UPDATE categories
		SET name = $1, description = $2
		WHERE id = $3
		RETURNING id
	`,
		updated.Name,
		updated.Description,
		id,
	).Scan(&updated.ID)

	if err != nil {
		return models.Category{}, err
	}

	return updated, nil
}

// ===== DELETE =====
func (r *CategoryRepository) Delete(id int) error {
	result, err := r.db.Exec(`
		DELETE FROM categories
		WHERE id = $1
	`, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
