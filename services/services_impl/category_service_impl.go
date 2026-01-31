package services_impl

import (
	"kasir-api/models"
	"kasir-api/models/dto"
	"kasir-api/repositories"
	"kasir-api/services"
)

type CategoryServiceImpl struct {
	repo repositories.CategoryRepository
}

func NewCategoryServiceImpl(repo repositories.CategoryRepository) services.CategoryService {
	return &CategoryServiceImpl{repo: repo}
}

func (s *CategoryServiceImpl) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryServiceImpl) Create(kategori *models.Category) error {
	return s.repo.Create(kategori)
}

func (s *CategoryServiceImpl) GetByID(id int) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryServiceImpl) Update(kategori *models.Category) error {
	return s.repo.Update(kategori)
}

func (s *CategoryServiceImpl) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *CategoryServiceImpl) GetProductListByCategoryID(categoryID int) (*dto.CategoryDetailDto, error) {
	return s.repo.GetProductListByCategoryID(categoryID)
}
