package service

import (
	"financial_literacy/internal/repository"
	"financial_literacy/models"
	"fmt"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo *repository.Repository) *TransactionService {
	return &TransactionService{repo}
}

func (service *TransactionService) CreateTransaction(transactionInfo *models.Transaction) error {
	return service.repo.CreateTransaction(transactionInfo)
}

func (service *TransactionService) ReadTransaction(id int) (*models.Transaction, error) {
	if id < 0 {
		return nil, fmt.Errorf("invalid transaction id")
	}
	return service.repo.ReadTransaction(id)
}

func (service *TransactionService) UpdateTransaction(transactionInfo *models.Transaction) error {
	return service.repo.UpdateTransaction(transactionInfo)
}
