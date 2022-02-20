package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (handler *Handler) ReadTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID := strings.TrimPrefix(r.URL.Path, "/read/transaction/")
	id, err := strconv.Atoi(transactionID)
	if err != nil {
		log.Println(err)
		return
	}
	transaction, err := handler.service.ReadTransaction(id)
	if err != nil {
		log.Println(err)
		return
	}
	response, err := json.Marshal(transaction)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(response)
}
