package database

import (
	"database/sql"
)

// IsEmpty checks if a database is empty
func IsEmpty(db *sql.DB) bool {
	_, err := db.Query("SELECT * FROM transactions;")
	if err != nil {
		return false
	}
	// for rows.Next() {
	// 	var id int
	// 	if err := rows.Scan(&id); err != nil {
	// 		return false
	// 	}
	// 	if id == 0 {
	// 		return false
	// 	}
	// }
	return true
}
