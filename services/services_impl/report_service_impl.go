package services_impl

import (
	"kasir-api/models/dto"
	"kasir-api/repositories"
	"kasir-api/services"
	"time"
)

type ReportServiceImpl struct {
	repository repositories.ReportRepository
}

func NewReportServiceImpl(repository repositories.ReportRepository) services.ReportService {
	return &ReportServiceImpl{repository: repository}
}

func (s *ReportServiceImpl) GetReportToday() (*dto.ReportDto, error) {
	return s.repository.GetReportToday()
}

func (s *ReportServiceImpl) GetReportByDate(startDate, endDate time.Time) (*dto.ReportDto, error) {
	return s.repository.GetReportByDate(startDate, endDate)
}
