package app

import (
	"context"
	"errors"
)

var (
	ErrRoomNotFound = errors.New("room not found")
)

type RoomRepository interface {
	// read
	GetListRoom(ctx context.Context, userID string, offset, limit int) ([]*WsRoom, error)
	FindRoomBetween(userA, userB string) (*WsRoom, error)

	// write
	CreateRoom(room *WsRoom) error
}

type MessageRepository interface {
	// read
	GetLastMessageByRoomID(ctx context.Context, roomID string) (*WsMessage, error)

	// write
	CreateMessage(msg *WsMessage) error
}
