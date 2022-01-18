package entity

import "time"

type Factory interface {
	NewRoom(userA, userB string) (*Room, error)
	NewMessage(sender, receiver, content string, room *Room) (*Message, error)
}

type implFactory struct {
}

func (fac implFactory) NewMessage(sender, receiver, content string, room *Room) (*Message, error) {
	return &Message{
		Sender:    sender,
		Receiver:  receiver,
		Content:   content,
		RoomID:    room.ID,
		CreatedAt: time.Now(),
	}, nil
}
