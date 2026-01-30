package services

import "kasir-api/models"

type ProductService interface {
	GetAll() ([]models.Product, error)
	Create(product *models.Product) error
	GetById(id int) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id int) error
}
