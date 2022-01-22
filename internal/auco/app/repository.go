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
	GetRoomByID(ctx context.Context, roomID string) (*WsRoom, error)
	GetListRoom(ctx context.Context, userID string, offset, limit int) ([]*WsRoom, error)
	FindRoomBetween(userA, userB string) (*WsRoom, error)

	// write
	CreateRoom(room *WsRoom) error
}

type MessageRepository interface {
	// read
	GetLastMessageByRoomID(ctx context.Context, roomID string) (*WsMessage, error)
	GetMessagesByRoomID(ctx context.Context, roomID string, offset, limit int) ([]*WsMessage, error)

	// write
	CreateMessage(msg *WsMessage) error
	DeleteByRoomID(ctx context.Context, roomID string) error
}
