package services_impl

import (
	"kasir-api/models"
	"kasir-api/repositories"
	"kasir-api/services"
)

type ProductServiceImpl struct {
	repo repositories.ProductRepository
}

func NewProductServiceImpl(repo repositories.ProductRepository) services.ProductService {
	return &ProductServiceImpl{repo: repo}
}

func (s *ProductServiceImpl) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductServiceImpl) Create(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductServiceImpl) GetById(id int) (*models.Product, error) {
	return s.GetById(id)
}

func (s *ProductServiceImpl) Update(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductServiceImpl) Delete(id int) error {
	return s.repo.Delete(id)
}
