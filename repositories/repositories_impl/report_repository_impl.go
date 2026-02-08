package repositories_impl

import (
	"database/sql"
	"fmt"
	"kasir-api/models/dto"
	"kasir-api/repositories"
	"time"
)

type ReportRepositoryImpl struct {
	db *sql.DB
}

func NewReportRepositoryImpl(db *sql.DB) repositories.ReportRepository {
	return &ReportRepositoryImpl{db: db}
}

func (r *ReportRepositoryImpl) GetReportToday() (*dto.ReportDto, error) {
	query := `
				WITH stats AS (
					SELECT 
						COALESCE(SUM(t.total_amount), 0) as total_revenue,
						COUNT(t.id) as total_transaksi
					FROM transactions t
					WHERE DATE(t.created_at) = DATE(NOW())
				),
				best_product AS (
					SELECT 
						p.name,
						SUM(td.quantity) as qty_terjual
					FROM transaction_details td
					JOIN products p ON p.id = td.product_id
					WHERE DATE(td.created_at) = DATE(NOW())
					GROUP BY p.id, p.name
					ORDER BY qty_terjual DESC
					LIMIT 1
				)
				SELECT 
					s.total_revenue,
					s.total_transaksi,
					COALESCE(bp.name, '-') as product_name,
					COALESCE(bp.qty_terjual, 0) as qty_sold
				FROM stats s
				JOIN best_product bp ON true
             `

	report := &dto.ReportDto{}
	err := r.db.QueryRow(query).Scan(
		&report.TotalRevenue,
		&report.TotalTransaksi,
		&report.ProdukTerlaris.ProductName,
		&report.ProdukTerlaris.QtyTerjual,
	)
	if err != nil {
		return nil, err
	}

	date := time.Now().Format("2006-01-02")
	report.Periode = fmt.Sprintf("%s", date)

	return report, nil
}

func (r *ReportRepositoryImpl) GetReportByDate(startDate, endDate time.Time) (*dto.ReportDto, error) {
	query := `
				WITH stats AS (
					SELECT 
						COALESCE(SUM(t.total_amount), 0) as total_revenue,
						COUNT(t.id) as total_transaksi
					FROM transactions t
					WHERE t.created_at::date BETWEEN $1 AND $2
				),
				best_product AS (
					SELECT 
						p.name,
						SUM(td.quantity) as qty_terjual
					FROM transaction_details td
					JOIN products p ON p.id = td.product_id
					WHERE td.created_at::date BETWEEN $1 AND $2
					GROUP BY p.id, p.name
					ORDER BY qty_terjual DESC
					LIMIT 1
				)
				SELECT 
					s.total_revenue,
					s.total_transaksi,
					COALESCE(bp.name, '-') as product_name,
					COALESCE(bp.qty_terjual, 0) as qty_sold
				FROM stats s
				LEFT JOIN best_product bp ON true
             `

	report := &dto.ReportDto{}
	err := r.db.QueryRow(query, startDate, endDate).Scan(
		&report.TotalRevenue,
		&report.TotalTransaksi,
		&report.ProdukTerlaris.ProductName,
		&report.ProdukTerlaris.QtyTerjual,
	)
	if err != nil {
		return nil, err
	}

	if startDate == endDate {
		report.Periode = startDate.Format("2006-01-02")
	} else {
		report.Periode = fmt.Sprintf("%s - %s",
			startDate.Format("2006-01-02"),
			endDate.Format("2006-01-02"))
	}

	return report, nil
}
