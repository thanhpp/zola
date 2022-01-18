package entity

import (
	"time"
)

type Message struct {
	RoomID    string
	Sender    string
	Receiver  string
	Content   string
	CreatedAt time.Time
}
