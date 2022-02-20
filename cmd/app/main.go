package main

import (
	"financial_literacy/database"
	"financial_literacy/internal/handler"
	"financial_literacy/internal/repository"
	"financial_literacy/internal/service"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := database.DBInit()
	if err != nil {
		log.Println(err)
		return
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	mux := http.NewServeMux()
	mux.HandleFunc("/add/category", handler.CreateCategory)
	mux.HandleFunc("/read/category/", handler.ReadCategory)
	mux.HandleFunc("/update/category", handler.UpdateCategory)
	mux.HandleFunc("/delete/category/", handler.DeleteCategory)
	mux.HandleFunc("/add/transaction", handler.CreateTransaction)
	mux.HandleFunc("/update/transaction", handler.UpdateTransaction)
	mux.HandleFunc("/read/transaction/", handler.ReadTransaction)
	log.Println("starting server at localhost:9090")
	err = http.ListenAndServe("localhost:9090", mux)
	if err != nil {
		log.Println(err)
		return
	}
}
