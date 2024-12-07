package models

type Booking struct {
	ID         int    `json:"id"`
	CarID      int    `json:"car_id"`
	CustomerID int    `json:"customer_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	TotalPrice int    `json:"total_price"`
}
