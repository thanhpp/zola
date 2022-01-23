package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/thanhpp/zola/internal/auco/app"
	"github.com/thanhpp/zola/pkg/logger"
)

type WsCtrl struct {
	upgrader  *websocket.Upgrader // upgrader is used to upgrade the HTTP server connection to the WebSocket protocol.
	wsManager *app.WsManager
}

const (
	readBufferSize  = 4096
	writeBufferSize = 4096
)

func NewWsController(wm *app.WsManager) *WsCtrl {
	return &WsCtrl{
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  readBufferSize,
			WriteBufferSize: writeBufferSize,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		wsManager: wm,
	}
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

	logger.Infof("WsCtrl - new client connected with id %s, name %s", id, name)
}
