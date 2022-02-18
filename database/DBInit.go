package database

import (
	"database/sql"
)

// DBInit initalizes database or otherwise returns nil and error
func DBInit() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "transactions.db")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if !IsEmpty(db) {
		if err := CreateTables(db); err != nil {
			return nil, err
		}
	}

	return db, nil
}
