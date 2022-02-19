package handler

import (
	"encoding/json"
	"financial_literacy/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (handler *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	transaction := &models.Transaction{}
	err = json.Unmarshal(body, &transaction)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(transaction)
	err = handler.service.CreateTransaction(transaction)
	if err != nil {
		log.Println(err)
		return
	}
}
