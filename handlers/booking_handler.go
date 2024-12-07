package handlers

import (
	"encoding/json"
	"net/http"
	"rental-system/models"
)

var bookings = []models.Booking{}

func GetAllBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	booking.ID = len(bookings) + 1
	bookings = append(bookings, booking)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
}
