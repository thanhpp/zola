package controller

import (
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	readBufferSize  = 4096
	writeBufferSize = 4096
)

func NewWsController() *WsCtrl {
	wsMan := NewWsManager()
	go wsMan.Run()
	return &WsCtrl{
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  readBufferSize,
			WriteBufferSize: writeBufferSize,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		wsManager: wsMan,
	}
}
