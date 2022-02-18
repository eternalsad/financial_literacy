package repository

import (
	"database/sql"
	"financial_literacy/models"
)

type CategoryRepo struct {
	db *sql.DB
}

// type category interface {
// 	CreateCategory(*models.Category) error
// 	ReadCategory() ([]*models.Category, error)
// 	UpdateCategory() error
// 	DeleteCategory(int) error
// }

func NewCategoryRepo(db *sql.DB) *CategoryRepo {
	return &CategoryRepo{db}
}

func (categoryRepo *CategoryRepo) CreateCategory(categoryInfo *models.Category) error {
	query := "INSERT into category (name) VALUES (?)"
	stmt, err := categoryRepo.db.Prepare(query)
	if err != nil {
		return err
	}
	stmt.Exec(categoryInfo.Name)
	return nil
}

func (categoryRepo *CategoryRepo) ReadCategory(id int) (*models.Category, error) {
	query := "SELECT * FROM category WHERE id = ?"
	stmt, err := categoryRepo.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	category := &models.Category{}
	for res.Next() {
		var name string
		err = res.Scan(&name)
		if err != nil {
			return nil, err
		}
		category.Name = name
	}
	return category, nil
}

func (category *CategoryRepo) UpdateCategory() error {
	return nil
}

func (category *CategoryRepo) DeleteCategory(id int) error {
	return nil
}
