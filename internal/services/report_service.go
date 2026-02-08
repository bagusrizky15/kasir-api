package services

import (
	"kasir-api/internal/models"
	"kasir-api/internal/repository"
)

type ReportService struct {
	repo *repository.ReportRepository
}

func NewReportService(repo *repository.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetTodaySalesReport() (*models.SalesReport, error) {
	totalRevenue, totalTransaksi, err := s.repo.GetTodaySummary()
	if err != nil {
		return nil, err
	}

	nama, qty, err := s.repo.GetTodayBestSeller()
	if err != nil {
		return nil, err
	}

	return &models.SalesReport{
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransaksi,
		ProdukTerlaris: models.BestSeller{
			Nama:       nama,
			QtyTerjual: qty,
		},
	}, nil
}
