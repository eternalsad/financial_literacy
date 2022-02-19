package repository

import (
	"database/sql"
	"financial_literacy/models"
)

// Repository struct to bind transaction
// and category interfaces
// will be used as interface with which service
// can communicate in order to have access to db
type Repository struct {
	Transaction
	Category
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Transaction: NewTransaction(db), Category: NewCategoryRepo(db)}
}

type Transaction interface {
	CreateTransaction(*models.Transaction) error
	ReadTransaction() ([]*models.Transaction, error)
	UpdateTransaction(*models.Transaction) error
	DeleteTransaction(int) error
}

type Category interface {
	CreateCategory(*models.Category) error
	ReadCategory(id int) (*models.Category, error)
	UpdateCategory(*models.Category) error
	DeleteCategory(int) error
}
