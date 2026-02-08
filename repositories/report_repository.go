package repositories

import (
	"kasir-api/models/dto"
	"time"
)

type ReportRepository interface {
	GetReportToday() (*dto.ReportDto, error)
	GetReportByDate(startDate, endDate time.Time) (*dto.ReportDto, error)
}
