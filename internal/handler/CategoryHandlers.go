package handler

import (
	"encoding/json"
	"financial_literacy/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (handler *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	category := &models.Category{}
	err = json.Unmarshal(body, &category)
	if err != nil {
		log.Println(err)
		return
	}
	err = handler.service.CreateCategory(category)
	if err != nil {
		log.Println(err)
		return
	}
}

func (handler *Handler) ReadCategory(w http.ResponseWriter, r *http.Request) {
	categoryID := strings.TrimPrefix(r.URL.Path, "/read/category/")
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		log.Println(err)
		return
	}
	category, err := handler.service.ReadCategory(id)
	if err != nil {
		log.Println(err)
		return
	}
	response, err := json.Marshal(category)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(response)
}

func (handler *Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	category := &models.Category{}
	err = json.Unmarshal(body, &category)
	if err != nil {
		log.Println(err)
		return
	}
	err = handler.service.UpdateCategory(category)
	if err != nil {
		log.Println(err)
		return
	}
}

func (handler *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryID := strings.TrimPrefix(r.URL.Path, "/delete/category/")
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		log.Println(err)
		return
	}
	err = handler.service.DeleteCategory(id)
	if err != nil {
		log.Println(err)
		return
	}
}
