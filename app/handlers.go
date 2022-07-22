package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
	customer := []Customer{
		{Name: "John", City: "New York", ZipCode: "10001"},
		{Name: "Jane", City: "New York", ZipCode: "10001"},
		{Name: "Joe", City: "New York", ZipCode: "10001"},
	}
	json.NewEncoder(w).Encode(customer)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "Customer: ", vars["id"])
}
