package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/thanhpp/zola/internal/auco/app"
	"github.com/thanhpp/zola/pkg/logger"
)

type WsCtrl struct {
	upgrader  *websocket.Upgrader // upgrader is used to upgrade the HTTP server connection to the WebSocket protocol.
	wsManager *app.WsManager
}

func (ctrl WsCtrl) ServeWebsocket(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	conn, err := ctrl.upgrader.Upgrade(c.Writer, c.Request, nil) // the return value of this function is a WebSocket connection
	if err != nil {
		logger.Errorf("upgrade connection error: %v", err)
		return
	}

	ctrl.wsManager.CreateClient(id, name, conn)

	logger.Debugf("new client connected with id %s, name %s", id, name)
}
