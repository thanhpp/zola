package controller

import "github.com/thanhpp/zola/pkg/logger"

type WebSocketManager struct {
	clients     map[*Client]struct{}
	registerC   chan *Client
	unregisterC chan *Client
	broadcast   chan []byte
}

func NewWsManager() *WebSocketManager {
	return &WebSocketManager{
		clients:     make(map[*Client]struct{}),
		registerC:   make(chan *Client),
		unregisterC: make(chan *Client),
		broadcast:   make(chan []byte),
	}
}

func (man *WebSocketManager) Run() {
	for {
		select {
		case client := <-man.registerC:
			logger.Debugf("WsManager: new client %v", client)
			man.registerClient(client)
		case client := <-man.unregisterC:
			man.unregisterClient(client)
		}
	}
}

func (man *WebSocketManager) registerClient(client *Client) {
	man.clients[client] = struct{}{}
}

func (man *WebSocketManager) unregisterClient(client *Client) {
	if _, ok := man.clients[client]; ok {
		delete(man.clients, client)
	}
}

func (man *WebSocketManager) broadcastToClients(msg []byte) {
	for client := range man.clients {
		client.send <- msg
		logger.Debugf("WsMan - send msg %s to client %s", msg, client)
	}
}
