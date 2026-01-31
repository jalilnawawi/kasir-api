package services

import (
	"kasir-api/models"
	"kasir-api/models/dto"
)

type CategoryService interface {
	GetAll() ([]models.Category, error)
	Create(kategori *models.Category) error
	GetByID(id int) (*models.Category, error)
	Update(kategori *models.Category) error
	Delete(id int) error
	GetProductListByCategoryID(categoryID int) (*dto.CategoryDetailDto, error)
}
