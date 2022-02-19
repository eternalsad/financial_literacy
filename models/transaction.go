package models

type Transaction struct {
	ID              int    `json:"id"`
	TransactionName string `json:"transaction_name"`
	Amount          int    `json:"amount"`
	TransactionDate int    `json:"transaction_date"`
	IsIncome        bool   `json:"isIncome"`
	Comment         string `json:"comment"`
	CategoryID      int    `json:"category_id"`
}
