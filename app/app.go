package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	fmt.Println("Starting server on port 8081")

	router.HandleFunc("/greet", Greet).Methods("GET")
	router.HandleFunc("/customers", GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers", createCustomer).Methods("POST")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", GetCustomer).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:8081", router))
}
