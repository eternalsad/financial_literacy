package service

import (
	"financial_literacy/internal/repository"
	"financial_literacy/models"
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
