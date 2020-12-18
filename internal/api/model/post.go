package model

import (
	"time"
)

type Post struct {
	ID          int    `bson:"_id" json:"id,omitempty"`
	Title       string `json:"title"`
	Description string
	UserID      int `bson:"user_id" json:"user_id,omitempty"`
	CreatedAt   time.Time
}
