package app

import "errors"

var (
	ErrCanNotJoinRoom = errors.New("can not join room")
)

type WsRoom struct {
	ID        string
	UserA     string
	UserB     string
	clientMap *wsClientMap
}

func (r WsRoom) key() string {
	if r.UserA > r.UserB {
		return r.UserB + "-" + r.UserA
	}
	return r.UserA + "-" + r.UserB
}

func (r *WsRoom) addClient(c *WsClient) error {
	if r.UserA == c.ID || r.UserB == c.ID {
		if _, ok := r.clientMap.findByID(c.ID); !ok {
			r.clientMap.add(c)
		}
		return nil
	}

	return ErrCanNotJoinRoom
}

func (r *WsRoom) sendMessageToAll(msgB []byte) {
	r.clientMap.walkLock(func(c *WsClient) {
		c.send(msgB)
	})
}
