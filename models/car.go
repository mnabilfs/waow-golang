package models

type Car struct {
	ID         int    `json:"id"`
	Model      string `json:"model"`
	Plate      string `json:"plate_number"`
	DailyRate  int    `json:"daily_rate"`
	Status     string `json:"status"` // available, booked, etc.
}
