package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/greet", Greet)
	router.HandleFunc("/customers", GetAllCustomers)
	router.HandleFunc("/customers/{id}", GetCustomer)

	log.Fatal(http.ListenAndServe("localhost:8081", router))
}
