package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (handler *Handler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID := strings.TrimPrefix(r.URL.Path, "/delete/transaction/")
	id, err := strconv.Atoi(transactionID)
	if err != nil {
		log.Println(err)
		return
	}
	err = handler.service.DeleteTransaction(id)
}
