package controller

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/thanhpp/zola/pkg/logger"
)

const (
	writeWait      = 10 * time.Second    // Max wait time when writing message to peer
	pongWait       = 60 * time.Second    // Max time till next pong from peer
	pingPeriod     = (pongWait * 9) / 10 // Send ping interval, must be less then pong wait time
	maxMessageSize = 10000               // Maximum message size allowed from peer.
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	ID        string
	UUID      uuid.UUID
	conn      *websocket.Conn
	wsManager *WebSocketManager
	send      chan []byte
	rooms     map[*WsRoom]struct{}
}

func (c *Client) readPump() {
	defer func() {
		// TODO: disconnect
		c.disconnect()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// Start endless read loop, waiting for messages from client
	for {
		_, jsonMessage, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}

		logger.Debugf("received message: %s", jsonMessage)
		c.handleMessage(jsonMessage)
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := c.conn.Close(); err != nil {
			logger.Errorf("close websocket connection err: %v", err)
		}
	}()

	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The WsManager closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				logger.Errorf("write websocket message err: %v", err)
				return
			}
			w.Write(msg)

			// Attach queued chat messages to the current websocket message.
			// write whatever is in the channel
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			// keep alive
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (client *Client) disconnect() {
	client.wsManager.unregisterC <- client
	for room := range client.rooms {
		room.UnregisterC <- client
	}
}

func (c *Client) handleMessage(msgB []byte) {
	var newMsg = new(WsMessage)
	if err := json.Unmarshal(msgB, newMsg); err != nil {
		logger.Errorf("handleMessage: decode error: %v", err)
		return
	}

	// attach the sender to the message
	newMsg.Sender = c

	switch newMsg.Action {
	case MessageActionSend:
		// get the room name
		roomName := newMsg.Target

		if room := c.wsManager.findRoomByName(roomName.ID); room != nil {
			room.Broadcast <- []byte(newMsg.Message)
		}

	case MessageActionJoin:
		logger.Debug("Client: handleJoinRoom")
		c.handleJoinRoom(newMsg)

	case MessageActionLeave:
		c.handleLeaveRoom(newMsg)
	}
}

func (c *Client) handleJoinRoom(msg *WsMessage) {
	roomName := msg.Message

	c.joinRoom(roomName, nil)
}

func (c *Client) handleLeaveRoom(msg *WsMessage) {
	room := c.wsManager.findRoomByName(msg.Message)

	if room != nil {
		return
	}

	if _, ok := c.rooms[room]; ok {
		delete(c.rooms, room)
		room.UnregisterC <- c
	}
}

func (c *Client) handleJoinRoomPrivateMessage(msg WsMessage) {
	target := c.wsManager.findClientByUUID(msg.Message)
	if target == nil {
		return
	}

	roomName := msg.Message + c.UUID.String()

	c.joinRoom(roomName, target)
	target.joinRoom(roomName, c)
}

func (c *Client) joinRoom(roomName string, sender *Client) {
	room := c.wsManager.findRoomByName(roomName)
	if room == nil {
		room = c.wsManager.createRoom(roomName, sender != nil)
	}

	if sender == nil && room.Private {
		return
	}

	if !c.isInRoom(room) {
		c.rooms[room] = struct{}{}
		room.RegisterC <- c
		c.notifyRoomJoined(room, sender)
	}
}

func (c *Client) isInRoom(room *WsRoom) bool {
	if _, ok := c.rooms[room]; ok {
		return true
	}
	return false
}

func (c *Client) notifyRoomJoined(room *WsRoom, sender *Client) {
	message := &WsMessage{
		Action: MessageActionJoin,
		Target: room,
		Sender: sender,
	}
	c.send <- message.Encode()
}
