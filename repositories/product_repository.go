package repositories

import (
	"kasir-api/models"
	"kasir-api/models/dto"
)

type ProductRepository interface {
	GetAll(nameFilter string) ([]models.Product, error)
	Create(produk *models.Product) error
	GetByID(id int) (*dto.ProductDetailDto, error)
	Update(produk *models.Product) error
	Delete(id int) error
}
