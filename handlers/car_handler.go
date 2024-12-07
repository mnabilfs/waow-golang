package handlers

import (
	"encoding/json"
	"net/http"
	"rental-system/models"
)

var cars = []models.Car{
	{ID: 1, Model: "Toyota Avanza", Plate: "B1234ABC", DailyRate: 500000, Status: "available"},
}

func GetAllCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	car.ID = len(cars) + 1
	cars = append(cars, car)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
}
