package service

import (
	"financial_literacy/internal/repository"
	"financial_literacy/models"
)

type Service struct {
	Category
	Transaction
}

type Category interface {
	CreateCategory(*models.Category) error
	ReadCategory(id int) (*models.Category, error)
	UpdateCategory(*models.Category) error
	DeleteCategory(int) error
}

type Transaction interface {
	CreateTransaction(*models.Transaction) error
	UpdateTransaction(*models.Transaction) error
	ReadTransaction(id int) (*models.Transaction, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Category:    NewCategoryService(repo),
		Transaction: NewTransactionService(repo)}
}
