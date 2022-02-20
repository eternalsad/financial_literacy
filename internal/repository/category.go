package repository

import (
	"database/sql"
	"financial_literacy/models"
	"fmt"
)

type CategoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) *CategoryRepo {
	return &CategoryRepo{db}
}

// type Category interface {
// 	CreateCategory(*models.Category) error
// 	ReadCategory(id int) (*models.Category, error)
// 	UpdateCategory(*models.Category) error
// 	DeleteCategory(int) error
// }

func (categoryRepo *CategoryRepo) CreateCategory(categoryInfo *models.Category) error {
	fmt.Println(categoryInfo)
	query := "INSERT into categories (category_name) VALUES (?)"
	stmt, err := categoryRepo.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(categoryInfo.Name)
	return err
}

func (categoryRepo *CategoryRepo) ReadCategory(id int) (*models.Category, error) {
	query := "SELECT * FROM categories WHERE id = ?"
	stmt, err := categoryRepo.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	category := &models.Category{}
	for res.Next() {
		var name string
		var id int
		err = res.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		category.ID = id
		category.Name = name
	}
	fmt.Println(category.Name)
	return category, nil
}

func (category *CategoryRepo) UpdateCategory(categoryInfo *models.Category) error {
	if categoryInfo.ID == 0 {
		fmt.Errorf("category id is nil")
	}
	if categoryInfo.Name == "" {
		fmt.Errorf("category category_name is nil")
	}
	query := "UPDATE categories SET category_name = ? WHERE id = ?"
	stmt, err := category.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(categoryInfo.Name, categoryInfo.ID)
	if err != nil {
		return err
	}
	return err
}

func (category *CategoryRepo) DeleteCategory(id int) error {
	query := "DELETE from categories WHERE id = ?"
	stmt, err := category.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
