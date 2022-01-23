package app

import (
	"sync"

	"github.com/thanhpp/zola/pkg/logger"
)

// ----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- ROOM MAP ----------------------------------------------------------

type wsRoomMap struct {
	m  map[string]*WsRoom
	rw sync.RWMutex
}

func newWsRoomMap() *wsRoomMap {
	return &wsRoomMap{
		m: make(map[string]*WsRoom),
	}
}

func (rm *wsRoomMap) add(room *WsRoom) {
	rm.rw.Lock()
	logger.Debugf("WsRoomMap: added room %s", room.key())
	rm.m[room.key()] = room
	rm.rw.Unlock()
}

func (rm *wsRoomMap) delete(room *WsRoom) {
	rm.rw.Lock()
	if _, ok := rm.m[room.key()]; ok {
		delete(rm.m, room.key())
	}
	rm.rw.Unlock()
}

func (rm *wsRoomMap) findByUserIDs(userA, userB string) (*WsRoom, bool) {
	rm.rw.RLock()
	if userA > userB {
		userA, userB = userB, userA
	}
	room, ok := rm.m[userA+"-"+userB]
	rm.rw.RUnlock()
	return room, ok
}

func (rm *wsRoomMap) walkRLock(fn func(*WsRoom)) {
	rm.rw.RLock()
	for _, room := range rm.m {
		fn(room)
	}
	rm.rw.RUnlock()
}

func (rm *wsRoomMap) walkLock(fn func(*WsRoom)) {
	rm.rw.Lock()
	for _, room := range rm.m {
		fn(room)
	}
	rm.rw.Unlock()
}

// ------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- CLIENT MAP ----------------------------------------------------------

type wsClientMap struct {
	m  map[string]*WsClient
	rw sync.RWMutex
}

func newWsClientMap() *wsClientMap {
	return &wsClientMap{
		m: make(map[string]*WsClient),
	}
}

func (cm *wsClientMap) add(client *WsClient) {
	cm.rw.Lock()
	// logger.Warnf("WsClientMap: add client %s - lock", client.ID)
	cm.m[client.ID] = client
	cm.rw.Unlock()
	// logger.Warnf("WsClientMap: add client %s - unlock", client.ID)
}

func (cm *wsClientMap) delete(client *WsClient) {
	cm.rw.Lock()
	// logger.Warnf("WsClientMap: delete client %s - lock", client.ID)
	if _, ok := cm.m[client.ID]; ok {
		delete(cm.m, client.ID)
	}
	cm.rw.Unlock()
	// logger.Warnf("WsClientMap: delete client %s - unlock", client.ID)
}

func (cm *wsClientMap) findByID(id string) (*WsClient, bool) {
	// logger.Warnf("WsClientMap: find client %s - lock", id)
	client, ok := cm.m[id]
	// logger.Warnf("WsClientMap: find client %s - unlock", id)
	return client, ok
}

func (cm *wsClientMap) walkLock(fn func(*WsClient)) {
	cm.rw.Lock()
	// logger.Debugf("WsClientMap: walkLock - lock")
	for _, client := range cm.m {
		fn(client)
	}
	cm.rw.Unlock()
	// logger.Debugf("WsClientMap: walkLock - unlock")
}
