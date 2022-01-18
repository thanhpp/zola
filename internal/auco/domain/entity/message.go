package entity

import (
	"time"
)

type MessageStatus string

func (ms *MessageStatus) String() string {
	return string(*ms)
}

type Message struct {
	RoomID    string
	Sender    string
	Receiver  string
	Content   string
	Status    MessageStatus
	CreatedAt time.Time
}

func (m Message) GetRoomID() string {
	return m.RoomID
}

func (m Message) GetSender() string {
	return m.Sender
}

func (m Message) GetReceiver() string {
	return m.Receiver
}

func (m Message) GetContent() string {
	return m.Content
}

func (m Message) GetStatus() string {
	return m.Status.String()
}

func (m Message) GetCreatedAt() time.Time {
	return m.CreatedAt
}
