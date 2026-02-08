package services

import (
	"kasir-api/models/dto"
	"time"
)

type ReportService interface {
	GetReportToday() (*dto.ReportDto, error)
	GetReportByDate(startDate, endDate time.Time) (*dto.ReportDto, error)
}
