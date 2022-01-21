package controller

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/thanhpp/zola/internal/auco/app"
)

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
