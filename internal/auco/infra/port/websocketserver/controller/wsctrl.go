package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/thanhpp/zola/pkg/logger"
)

type WsCtrl struct {
	upgrader  *websocket.Upgrader // upgrader is used to upgrade the HTTP server connection to the WebSocket protocol.
	wsManager *WebSocketManager
}

func (ctrl WsCtrl) ServeWebsocket(c *gin.Context) {
	conn, err := ctrl.upgrader.Upgrade(c.Writer, c.Request, nil) // the return value of this function is a WebSocket connection
	if err != nil {
		log.Println(err)
		return
	}

	name, ok := c.GetQuery("name")
	if !ok || len(name) == 0 {
		logger.Errorf("WsCtrl: invalid name")
		return
	}

	client := ctrl.newClient(conn, name)

	go client.writePump()
	go client.readPump()

	ctrl.wsManager.registerC <- client

	logger.Infof("new client %v", client)
}

func (ctrl WsCtrl) newClient(conn *websocket.Conn, name string) *Client {
	return &Client{
		ID:        name,
		UUID:      uuid.New(),
		conn:      conn,
		wsManager: ctrl.wsManager,
		send:      make(chan []byte),
		rooms:     make(map[*WsRoom]struct{}),
	}
}
