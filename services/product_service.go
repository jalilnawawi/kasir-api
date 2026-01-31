package services

import (
	"kasir-api/models"
	"kasir-api/models/dto"
)

type ProductService interface {
	GetAll() ([]models.Product, error)
	Create(product *models.Product) error
	GetById(id int) (*dto.ProductDetailDto, error)
	Update(product *models.Product) error
	Delete(id int) error
}
