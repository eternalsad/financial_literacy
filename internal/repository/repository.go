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
	transaction
	category
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{transaction: NewTransaction(db), category: NewCategoryRepo(db)}
}

type transaction interface {
	CreateTransaction(*models.Transaction) error
	ReadTransaction() ([]*models.Transaction, error)
	UpdateTransaction() error
	DeleteTransaction(int) error
}

type category interface {
	CreateCategory(*models.Category) error
	ReadCategory(id int) (*models.Category, error)
	UpdateCategory() error
	DeleteCategory(int) error
}
