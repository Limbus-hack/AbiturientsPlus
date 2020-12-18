package model

import (
	"time"
)

type User struct {
	ID        int `bson:"_id" json:"id,omitempty"`
	Name      string
	Surname   string
	NickName  string
	Email     string
	CreatedAt time.Time
}
