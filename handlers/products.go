package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"tech-test/helpers"
	"tech-test/models"
	"tech-test/usecases"

	"github.com/gorilla/mux"
)

// ProductsHandlersInterface ...
type ProductsHandlersInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
	DeleteByID(w http.ResponseWriter, r *http.Request)
}

// Products Handlers
type Products struct{}

// NewProductsHandlers ...
func NewProductsHandlers() ProductsHandlersInterface {
	return Products{}
}

// Create ...
func (v Products) Create(w http.ResponseWriter, r *http.Request) {
	res := helpers.Response{}
	defer res.ServeJSON(w, r)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res.Err = fmt.Errorf("handlers: could not read request body: %s", err.Error())
		return
	}
	defer r.Body.Close()

	var p models.Product
	err = json.Unmarshal(b, &p)
	if err != nil {
		res.Err = fmt.Errorf("handlers: could not unmarshal request body to model: %s", err.Error())
		return
	}

	err = usecases.NewProductsUsecase().Create(&p)
	if err != nil {
		res.Err = fmt.Errorf("handlers: %s", err.Error())
		return
	}

	res.Body.Payload = p
}

// GetByID ...
func (v Products) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	res := helpers.Response{}
	defer res.ServeJSON(w, r)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.Err = fmt.Errorf("handlers: could not convert string to int %s", err.Error())
		return
	}

	var p models.Product
	err = usecases.NewProductsUsecase().GetByID(id, &p)
	if err != nil {
		res.Err = fmt.Errorf("handlers: %s", err.Error())
		return
	}
	res.Body.Payload = p
}

// GetAll ...
func (v Products) GetAll(w http.ResponseWriter, r *http.Request) {

	queryParam := r.URL.Query()
	// Set default query
	limit, lastID := "2", "0"
	if v := queryParam.Get("limit"); v != "" {
		limit = queryParam.Get("limit")
	}
	if v := queryParam.Get("lastID"); v != "" {
		lastID = queryParam.Get("lastID")
	}

	res := helpers.Response{}
	defer res.ServeJSON(w, r)

	res.Body.Payload, res.Err = usecases.NewProductsUsecase().GetAll(limit, lastID)

}

// UpdateByID ...
func (v Products) UpdateByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	res := helpers.Response{}
	defer res.ServeJSON(w, r)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.Err = fmt.Errorf("handlers: could not convert string to int %s", err.Error())
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res.Err = fmt.Errorf("handlers: could not read request body: %s", err.Error())
		return
	}
	defer r.Body.Close()

	var p models.Product
	err = json.Unmarshal(b, &p)
	if err != nil {
		res.Err = fmt.Errorf("handlers: could not unmarshal request body to model: %s", err.Error())
		return
	}

	ra, err := usecases.NewProductsUsecase().UpdateByID(id, &p)
	if err != nil {
		res.Err = fmt.Errorf("handlers: %s", err.Error())
		return
	}
	res.Body.Payload = fmt.Sprintf("row affected: %d", ra)
}

// DeleteByID ...
func (v Products) DeleteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	res := helpers.Response{}
	defer res.ServeJSON(w, r)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.Err = fmt.Errorf("handlers: could not convert string to int %s", err.Error())
		return
	}
	defer r.Body.Close()

	ra, err := usecases.NewProductsUsecase().DeleteByID(id)
	if err != nil {
		res.Err = fmt.Errorf("handlers: %s", err.Error())
		return
	}
	res.Body.Payload = fmt.Sprintf("row affected: %d", ra)
}
