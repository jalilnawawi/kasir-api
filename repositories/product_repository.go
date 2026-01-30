package repositories

import "kasir-api/models"

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	Create(produk *models.Product) error
	GetByID(id int) (*models.Product, error)
	Update(produk *models.Product) error
	Delete(id int) error
}
