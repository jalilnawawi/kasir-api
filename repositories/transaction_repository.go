package repositories

import (
	"kasir-api/models"
	"kasir-api/models/dto"
)

type TransactionRepository interface {
	CreateTransaction(items []dto.CheckoutItem) (*models.Transaction, error)
	GetAllTransaction() (*[]dto.TransactionDto, error)
}
