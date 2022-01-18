package controller

import (
	"log"
	"time"

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
	conn      *websocket.Conn
	wsManager *WebSocketManager
	send      chan []byte
}

func (c *Client) readPump() {
	defer func() {
		// TODO: disconnect
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
		c.wsManager.broadcastToClients(jsonMessage)
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
