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

func (r Room) GetID() string {
	return r.ID
}

func (r Room) GetUserA() string {
	return r.UserA
}

func (r Room) GetUserB() string {
	return r.UserB
}

func (r Room) GetCreatedAt() time.Time {
	return r.CreatedAt
}

func (r Room) GetUpdatedAt() time.Time {
	return r.UpdatedAt
}
