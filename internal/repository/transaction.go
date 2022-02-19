package repository

import (
	"database/sql"
	"financial_literacy/models"
)

// TransactionRepo ...
type TransactionRepo struct {
	db *sql.DB
}

// NewTransaction create new instance of
// TransactionRepo
func NewTransaction(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{db}
}

// type transaction interface {
// 	CreateTransaction(*models.Transaction) error
// 	ReadTransaction() ([]*models.Transaction, error)
// 	UpdateTransaction() error
// 	DeleteTransaction(int) error
// }

func (transactionRepo *TransactionRepo) CreateTransaction(transactionInfo *models.Transaction) error {
	query := `INSERT INTO transactions (transaction_name, amount, transaction_date, 
	isIncome, comment, category_id) VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := transactionRepo.db.Prepare(query)
	if err != nil {
		return err
	}
	// t := transactionInfo
	_, err = stmt.Exec(transactionInfo.TransactionName, transactionInfo.Amount, transactionInfo.TransactionDate, transactionInfo.IsIncome, transactionInfo.Comment, transactionInfo.CategoryID)
	return err
}

func (transactionRepo *TransactionRepo) ReadTransaction() ([]*models.Transaction, error) {
	return nil, nil
}

func (transactionRepo *TransactionRepo) UpdateTransaction(*models.Transaction) error {
	return nil
}

func (transactionRepo *TransactionRepo) DeleteTransaction(id int) error {
	return nil
}
