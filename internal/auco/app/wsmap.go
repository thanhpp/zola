package app

import "sync"

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
	cm.m[client.ID] = client
	cm.rw.Unlock()
}

func (cm *wsClientMap) delete(client *WsClient) {
	cm.rw.Lock()
	if _, ok := cm.m[client.ID]; ok {
		delete(cm.m, client.ID)
	}
	cm.rw.Unlock()
}

func (cm *wsClientMap) findByID(id string) (*WsClient, bool) {
	cm.rw.RLock()
	client, ok := cm.m[id]
	cm.rw.RUnlock()
	return client, ok
}

func (cm *wsClientMap) walkLock(fn func(*WsClient)) {
	cm.rw.Lock()
	for _, client := range cm.m {
		fn(client)
	}
	cm.rw.Unlock()
}
