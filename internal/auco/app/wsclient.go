package app

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/thanhpp/zola/pkg/logger"
)

const (
	writeWait      = 10 * time.Second      // Max wait time when writing message to peer
	pongWait       = 60 * time.Second      // Max time till next pong from peer
	pingPeriod     = (pongWait * 9) / 10   // Send ping interval, must be less then pong wait time
	maxMessageSize = 10000                 // Maximum message size allowed from peer.
	readPeriod     = 10 * time.Millisecond // Prevent CPU overload
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type WsClient struct {
	ID        string // user id
	Name      string // username
	conn      *websocket.Conn
	wsManager *WsManager
	sendC     chan []byte
}

func (c *WsClient) readPump() {
	defer func() {
		logger.Infof("WsClient %s: read pump stopped - 1", c.ID)
		c.disconnect()
	}()

	// setup
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	// Start endless read loop, waiting for messages from client
	for {
		_, jsonMessage, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Errorf("WsClient %s: unexpected close error: %v", c.ID, err)
			}
			break
		}

		logger.Debugf("WsClient %s: received message: %s", c.ID, string(jsonMessage))
		c.handleNewMessage(jsonMessage)
	}

}

func (c *WsClient) writePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		logger.Infof("WsClient %s: write pump stopped - 1", c.ID)
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.sendC:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The WsManager closed the channel.
				err := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				logger.Errorf("WsClient %s: close message %v", c.ID, err)
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				logger.Errorf("WsClient %s: create writer err: %v", c.ID, err)
				return
			}
			if _, err := w.Write(msg); err != nil {
				logger.Errorf("WsClient %s: write message 1 err: %v", c.ID, err)

			}

			// Write queued chat messages to the current websocket message.
			n := len(c.sendC)
			for i := 0; i < n; i++ {
				if _, err := w.Write(newline); err != nil {

					logger.Errorf("WsClient %s: write message 2 err: %v", c.ID, err)

				}
				if _, err := w.Write(<-c.sendC); err != nil {

					logger.Errorf("WsClient %s: write message 3 err: %v", c.ID, err)

				}
			}

			if err := w.Close(); err != nil {
				logger.Errorf("WsClient %s: close writer err: %v", c.ID, err)

			}
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				logger.Errorf("WsClient %s: set write deadline err: %v", c.ID, err)

			}

		case <-ticker.C:
			// keep alive. server -> user
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				logger.Errorf("WsClient %s: set write deadline err: %v", c.ID, err)
				return
			}
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logger.Errorf("WsClient %s: ping message err: %v", c.ID, err)
				return
			}
		}
	}

}

func (c *WsClient) send(msg []byte) {
	select {
	case c.sendC <- msg:
	default:
		logger.Errorf("WsClient %s: send message failed", c.ID)
	}
}

func (c *WsClient) disconnect() {
	// logger.Warnf("WsClient %s: disconnecting", c.ID)
	c.wsManager.clientMap.delete(c)
	// logger.Warnf("WsClient %s: disconnecting - CP1", c.ID)
	c.wsManager.roomMap.walkLock(func(wr *WsRoom) {
		wr.clientMap.delete(c)
	})
	close(c.sendC)
	// logger.Warnf("WsClient %s: disconnecting - CP2", c.ID)
	if err := c.conn.Close(); err != nil {
		logger.Errorf("WsClient %s: close websocket connection err: %v", c.ID, err)
	}
	// logger.Warnf("WsClient %s: disconnecting - CP3", c.ID)
	logger.Infof("WsClient %s: disconnected", c.ID)
}
