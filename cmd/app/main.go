package main

import (
	"financial_literacy/database"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_, err := database.DBInit()
	if err != nil {
		log.Println(err)
		return
	}

}
