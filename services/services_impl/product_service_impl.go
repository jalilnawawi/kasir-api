package services_impl

import (
	"kasir-api/models"
	"kasir-api/models/dto"
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

func (s *ProductServiceImpl) GetById(id int) (*dto.ProductDetailDto, error) {
	return s.repo.GetByID(id)
}

func (s *ProductServiceImpl) Update(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductServiceImpl) Delete(id int) error {
	return s.repo.Delete(id)
}
