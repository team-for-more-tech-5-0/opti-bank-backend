package models

type Customer struct {
	ID        int64   `json:"customer_id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
