package repository

import (
	"database/sql"
	"financial_literacy/models"
	"fmt"
	"reflect"
	"strconv"
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

func (transactionRepo *TransactionRepo) ReadTransaction(id int) (*models.Transaction, error) {
	query := "SELECT * from transactions WHERE id=?"
	stmt, err := transactionRepo.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	transaction := &models.Transaction{}
	for res.Next() {
		var name string
		var amount float32
		var date int
		var isIncome bool
		var comment string
		var categoryID int
		err = res.Scan(&id, &name, &amount, &date, &isIncome, &comment, &categoryID)
		if err != nil {
			return nil, err
		}
		transaction.TransactionName = name
		transaction.Amount = amount
		transaction.TransactionDate = date
		transaction.IsIncome = isIncome
		transaction.Comment = comment
		transaction.CategoryID = categoryID
	}
	return transaction, nil
}

func (transactionRepo *TransactionRepo) UpdateTransaction(transactionInfo *models.Transaction) error {
	query := "UPDATE transactions "
	columns := []string{}
	var values []interface{}
	e := reflect.ValueOf(transactionInfo).Elem()
	fmt.Println(e)
	for i := 0; i < e.NumField(); i++ {
		if !e.Field(i).IsZero() {
			columns = append(columns, e.Type().Field(i).Tag.Get("json"))
			values = append(values, e.Field(i).Interface())
		}
	}
	vals := "SET "
	for i, c := range columns {
		vals += c
		vals += " = "
		v := fmt.Sprintf("%v", values[i])
		if _, err := strconv.Atoi(v); err != nil {
			vals += `"` + v + `"`
		} else {
			vals += v
		}
		if i < len(columns)-1 {
			vals += ", "
		}
	}
	query += vals
	query += fmt.Sprintf(" WHERE id=%v", transactionInfo.ID)
	fmt.Println(query)
	stmt, err := transactionRepo.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}

func (transactionRepo *TransactionRepo) DeleteTransaction(id int) error {
	query := "DELETE FROM transactions WHERE id = ?"
	stmt, err := transactionRepo.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
