package app

import (
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
