package service

import (
	"financial_literacy/internal/repository"
	"financial_literacy/models"
	"fmt"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo *repository.Repository) *CategoryService {
	return &CategoryService{repo}
}

func (service *CategoryService) CreateCategory(categoryInfo *models.Category) error {
	return service.repo.CreateCategory(categoryInfo)
}

func (service *CategoryService) ReadCategory(id int) (*models.Category, error) {
	return service.repo.ReadCategory(id)
}

func (service *CategoryService) UpdateCategory(categoryInfo *models.Category) error {
	if categoryInfo.ID < 0 {
		return fmt.Errorf("invalid category id")
	}
	return service.repo.UpdateCategory(categoryInfo)
}
func (service *CategoryService) DeleteCategory(id int) error {
	if id < 0 {
		return fmt.Errorf("invalid category id")
	}
	return service.repo.DeleteCategory(id)
}
