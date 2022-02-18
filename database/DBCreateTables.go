package database

import (
	"database/sql"
	"financial_literacy/config"
	"fmt"
	"io/ioutil"
)

// CreateTables initializes tables inside db
// otherwise throws error
func CreateTables(db *sql.DB) error {
	files, err := ioutil.ReadDir(config.SchemaDIRUP)
	if err != nil {
		return err
	}
	for _, fileInfo := range files {
		query, err := ioutil.ReadFile(config.SchemaDIRUP + fileInfo.Name())
		if err != nil {
			return err
		}
		_, err = db.Exec(string(query))
		if err != nil {
			fmt.Println(fileInfo.Name())
			return err
		}
	}
	return nil
}
