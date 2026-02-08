package repository

import (
	"database/sql"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetTodaySummary() (totalRevenue int, totalTransaksi int, err error) {
	err = r.db.QueryRow(`
		SELECT
			COALESCE(SUM(total_amount), 0),
			COUNT(*)
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`).Scan(&totalRevenue, &totalTransaksi)

	return
}

func (r *ReportRepository) GetTodayBestSeller() (nama string, qty int, err error) {
	err = r.db.QueryRow(`
		SELECT
			p.name,
			SUM(td.quantity) AS qty_terjual
		FROM transaction_details td
		JOIN transactions t ON t.id = td.transaction_id
		JOIN products p ON p.id = td.product_id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`).Scan(&nama, &qty)

	if err == sql.ErrNoRows {
		return "", 0, nil
	}

	return
}
