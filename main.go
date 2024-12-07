package main

import (
	"log"
	"net/http"
	"rental-system/handlers"
	"rental-system/middleware"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Middleware
	r.Use(middleware.CORS)

	// Auth routes
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/register", handlers.Register).Methods("POST")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTAuth)

	// Car routes
	api.HandleFunc("/cars", handlers.GetAllCars).Methods("GET")
	api.HandleFunc("/cars", handlers.CreateCar).Methods("POST")
	api.HandleFunc("/cars/{id}", handlers.UpdateCar).Methods("PUT")
	api.HandleFunc("/cars/{id}", handlers.DeleteCar).Methods("DELETE")

	// Customer routes
	api.HandleFunc("/customers", handlers.GetAllCustomers).Methods("GET")
	api.HandleFunc("/customers", handlers.CreateCustomer).Methods("POST")
	api.HandleFunc("/customers/{id}", handlers.UpdateCustomer).Methods("PUT")
	api.HandleFunc("/customers/{id}", handlers.DeleteCustomer).Methods("DELETE")

	// Booking routes
	api.HandleFunc("/bookings", handlers.GetAllBookings).Methods("GET")
	api.HandleFunc("/bookings", handlers.CreateBooking).Methods("POST")
	api.HandleFunc("/bookings/{id}", handlers.UpdateBooking).Methods("PUT")
	api.HandleFunc("/bookings/{id}", handlers.DeleteBooking).Methods("DELETE")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
