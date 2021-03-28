package routes

import (
	"fmt"
	"net/http"
	"tech-test/handlers"
	"tech-test/middleware"

	"github.com/gorilla/mux"
)

// Initiation ...
func Initiation() *mux.Router {

	r := mux.NewRouter()

	// for check health purpose
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Pong!")
	}).Methods("GET")

	// Versioning API
	r1 := r.PathPrefix("/v1").Subrouter()

	// Middleware
	r1.Use(middleware.CORS)

	r1.HandleFunc("/products", handlers.NewProductsHandlers().Create).Methods("POST")
	r1.HandleFunc("/products", handlers.NewProductsHandlers().GetAll).Methods("GET")
	r1.HandleFunc("/products/{id}", handlers.NewProductsHandlers().GetByID).Methods("GET")
	r1.HandleFunc("/products/{id}", handlers.NewProductsHandlers().UpdateByID).Methods("PUT")
	r1.HandleFunc("/products/{id}", handlers.NewProductsHandlers().DeleteByID).Methods("DELETE")

	return r
}
