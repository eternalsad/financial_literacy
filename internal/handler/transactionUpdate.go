package handler

import (
	"encoding/json"
	"financial_literacy/models"
	"io/ioutil"
	"log"
	"net/http"
)

func (handler *Handler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
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
	err = handler.service.UpdateTransaction(transaction)
	if err != nil {
		log.Println(err)
		return
	}
}
