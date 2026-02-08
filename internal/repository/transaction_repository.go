package repository

import (
	"database/sql"
	"fmt"
	"kasir-api/internal/models"
	"time"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	stmtDetail, err := tx.Prepare(`
	INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal)
	VALUES ($1, $2, $3, $4)
	RETURNING id
`)
	if err != nil {
		return nil, err
	}
	defer stmtDetail.Close()

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow(`
			SELECT name, price, stock
			FROM products
			WHERE id = $1
			FOR UPDATE
		`, item.ProductID).Scan(&productName, &productPrice, &stock)

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		if stock < item.Quantity {
			return nil, fmt.Errorf("stock product %s not enough", productName)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec(`
			UPDATE products
			SET stock = stock - $1
			WHERE id = $2
		`, item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	var createdAt time.Time
	err = tx.QueryRow(`
	INSERT INTO transactions (total_amount)
	VALUES ($1)
	RETURNING id, created_at
`, totalAmount).Scan(&transactionID, &createdAt)

	if err != nil {
		return nil, err
	}

	for i := range details {
		details[i].TransactionID = transactionID

		err = stmtDetail.QueryRow(
			transactionID,
			details[i].ProductID,
			details[i].Quantity,
			details[i].Subtotal,
		).Scan(&details[i].ID)

		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}
