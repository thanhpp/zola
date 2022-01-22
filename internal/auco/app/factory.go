package app

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WsFactory struct {
}

func (fac WsFactory) NewRoom(userA, userB string) (*WsRoom, error) {
	return &WsRoom{
		ID:        uuid.NewString(),
		UserA:     userA,
		UserB:     userB,
		clientMap: newWsClientMap(),
	}, nil
}

func (fac WsFactory) NewRoomFromDB(id, userA, userB string) (*WsRoom, error) {
	return &WsRoom{
		ID:        id,
		UserA:     userA,
		UserB:     userB,
		clientMap: newWsClientMap(),
	}, nil
}

func (fac WsFactory) NewClient(id, name string, conn *websocket.Conn, wm *WsManager) (*WsClient, error) {
	return &WsClient{
		ID:        id,
		Name:      name,
		conn:      conn,
		wsManager: wm,
		sendC:     make(chan []byte),
	}, nil
}

func (fac WsFactory) NewMessage(roomID, senderID, receiverID, createdAt, content string, seen bool) (*WsMessage, error) {
	// time check
	createdAtInt64, err := strconv.ParseInt(createdAt, 10, 64)
	if err != nil {
		return nil, err
	}

	// form messageID
	msgID := fmt.Sprintf("%s|%s", roomID, createdAt)

	return &WsMessage{
		MsgID:      msgID,
		Created:    createdAt,
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		roomID:     roomID,
		createdAt:  createdAtInt64,
		seen:       seen,
	}, nil
}
