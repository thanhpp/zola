package app

import (
	"github.com/gorilla/websocket"
	"github.com/thanhpp/zola/pkg/logger"
)

type WsManager struct {
	fac       *WsFactory
	clientMap *wsClientMap
	roomMap   *wsRoomMap
	roomRepo  RoomRepository
	msgRepo   MessageRepository
}

func NewWsManager(roomRepo RoomRepository, msgRepo MessageRepository) *WsManager {
	return &WsManager{
		fac:       &WsFactory{},
		clientMap: newWsClientMap(),
		roomMap:   newWsRoomMap(),
		roomRepo:  roomRepo,
		msgRepo:   msgRepo,
	}
}

func (wm *WsManager) findRoom(userA, userB string) (*WsRoom, bool) {
	// check in the cache
	room, ok := wm.roomMap.findByUserIDs(userA, userB)
	if ok {
		return room, ok
	}

	// check in the repository
	var err error
	room, err = wm.roomRepo.FindRoomBetween(userA, userB)
	if err == nil {
		// add to cache
		wm.roomMap.add(room)
		return room, ok
	}

	logger.Errorf("WsManager: findRoom error %v", err)
	return nil, false
}

func (wm *WsManager) createRoom(userA, userB string) (*WsRoom, error) {
	room, err := wm.fac.NewRoom(userA, userB)
	if err != nil {
		return nil, err
	}

	// add to repository
	err = wm.roomRepo.CreateRoom(room)
	if err != nil {
		return nil, err
	}

	// add to cache
	wm.roomMap.add(room)

	return room, nil
}

func (wm *WsManager) CreateClient(id, name string, conn *websocket.Conn) (*WsClient, error) {
	client, err := wm.fac.NewClient(id, name, conn, wm)
	if err != nil {
		return nil, err
	}

	go client.writePump()
	go client.readPump()

	wm.clientMap.add(client)

	return client, nil
}
