package handlers

import (
	"encoding/json"
	"net/http"
	"rental-system/models"
)

var customers = []models.Customer{
	{ID: 1, Name: "John Doe", Phone: "08123456789", Address: "Jl. Example No. 1"},
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	customer.ID = len(customers) + 1
	customers = append(customers, customer)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
}
