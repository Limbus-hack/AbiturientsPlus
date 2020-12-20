package model

type User struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	LastName   string `json:"last_name" db:"last_name"`
	Region     int    `json:"region" db:"region"`
	Prediction int    `json:"prediction" db:"prediction"`
	Status     string `json:"status" db:"status"`
}
