package app

import "github.com/thanhpp/zola/pkg/logger"

func (c *WsClient) handleNewMessage(msgB []byte) {
	var newMsg = new(WsMessage)
	if err := newMsg.Decode(msgB); err != nil {
		logger.Errorf("WsClient %s: decode message error: %v", c.ID, err)
	}

	switch newMsg.Event {
	case MsgEventJoin:
		c.handleJoin(newMsg)

	case MsgEventReconnect:

	case MsgEventAvaliable:

	case MsgEventDisconnect:

	case MsgEventDeleteMessage:

	case MsgEventSend:
		c.handleSend(newMsg)

	default:
		logger.Errorf("WsClient %s: unknown event: %s", c.ID, newMsg.Event)
		return
	}
}

func (c *WsClient) handleJoin(msg *WsMessage) {
	// find the room
	var err error
	room, ok := c.wsManager.findRoom(msg.SenderID, msg.ReceiverID)
	if !ok {
		room, err = c.wsManager.createRoom(msg.SenderID, msg.ReceiverID)
		if err != nil {
			logger.Errorf("WsClient %s: create room error: %v", c.ID, err)
			return
		}
	}

	// add the client to the room
	if err = room.addClient(c); err != nil {
		logger.Errorf("WsClient %s: add client to room error: %v", c.ID, err)
		return
	}
}

func (c *WsClient) handleSend(msg *WsMessage) {
	// find the room
	room, ok := c.wsManager.findRoom(msg.SenderID, msg.ReceiverID)
	if !ok {
		logger.Errorf("WsClient %s: find room error: %s", c.ID, msg.ReceiverID)
		return
	}

	// send message to the room
	room.sendMessageToAll([]byte(msg.Content))
}
