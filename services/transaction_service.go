package services

import (
	"kasir-api/models"
	"kasir-api/models/dto"
)

type TransactionService interface {
	Checkout(items []dto.CheckoutItem) (*models.Transaction, error)
	GetAllTransaction() (*[]dto.TransactionDto, error)
}
