package app

import "errors"

var (
	ErrRoomNotFound = errors.New("room not found")
)

type RoomRepository interface {
	FindRoomBetween(userA, userB string) (*WsRoom, error)
	CreateRoom(room *WsRoom) error
}

type MessageRepository interface {
	CreateMessage(msg *WsMessage) error
}
