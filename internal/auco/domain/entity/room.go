package entity

import (
	"time"
)

type Room struct {
	ID        string
	UserA     string
	UserB     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
