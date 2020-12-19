package model

type User struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	LastName   string `db:"last_name"`
	Region     int    `db:"region"`
	Prediction int    `db:"prediction"`
	Status     string `db:"status"`
}
