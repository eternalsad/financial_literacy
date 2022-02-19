package service

import (
	"financial_literacy/internal/repository"
	"financial_literacy/models"
)

type Service struct {
	Category
}

type Category interface {
	CreateCategory(*models.Category) error
	ReadCategory(id int) (*models.Category, error)
	UpdateCategory(*models.Category) error
	DeleteCategory(int) error
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Category: NewCategoryService(repo),
	}
}
