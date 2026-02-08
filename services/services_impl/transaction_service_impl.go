package services_impl

import (
	"kasir-api/models"
	"kasir-api/models/dto"
	"kasir-api/repositories"
	"kasir-api/services"
)

type TransactionServiceImpl struct {
	repository repositories.TransactionRepository
}

func NewTransactionServiceImpl(repository repositories.TransactionRepository) services.TransactionService {
	return &TransactionServiceImpl{repository: repository}
}

func (s *TransactionServiceImpl) Checkout(items []dto.CheckoutItem) (*models.Transaction, error) {
	return s.repository.CreateTransaction(items)
}

func (s *TransactionServiceImpl) GetAllTransaction() (*[]dto.TransactionDto, error) {
	return s.repository.GetAllTransaction()
}
