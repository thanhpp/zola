package controller

import "github.com/thanhpp/zola/pkg/logger"

type WebSocketManager struct {
	clients     map[*Client]struct{}
	rooms       map[*WsRoom]struct{}
	registerC   chan *Client
	unregisterC chan *Client
	broadcast   chan []byte
}

func NewWsManager() *WebSocketManager {
	return &WebSocketManager{
		clients:     make(map[*Client]struct{}),
		rooms:       make(map[*WsRoom]struct{}),
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
	man.notifyClientJoined(client)
	man.listOnlineCLients(client)
}

func (man *WebSocketManager) unregisterClient(client *Client) {
	if _, ok := man.clients[client]; ok {
		man.notifyClientLeft(client)
		delete(man.clients, client)
	}
}

func (man *WebSocketManager) broadcastToClients(msg []byte) {
	for client := range man.clients {
		client.send <- msg
	}
}

func (man *WebSocketManager) findRoomByName(name string) *WsRoom {
	for room := range man.rooms {
		if room.ID == name {
			return room
		}
	}

	return nil
}

func (man *WebSocketManager) createRoom(name string, private bool) *WsRoom {
	room := NewRoom(name, private)
	go room.Run()
	man.rooms[room] = struct{}{}

	return room
}

func (man *WebSocketManager) notifyClientJoined(client *Client) {
	message := &WsMessage{
		Action: MessageActionUserJoined,
		Sender: client,
	}

	man.broadcastToClients(message.Encode())
}

func (man *WebSocketManager) notifyClientLeft(client *Client) {
	message := &WsMessage{
		Action: MessageActionUserLeft,
		Sender: client,
	}

	man.broadcastToClients(message.Encode())
}

func (man *WebSocketManager) listOnlineCLients(client *Client) {
	for existingClient := range man.clients {
		message := &WsMessage{
			Action: MessageActionUserJoined,
			Sender: existingClient,
		}
		client.send <- message.Encode()
	}
}

func (man *WebSocketManager) findClientByUUID(in string) *Client {
	for client := range man.clients {
		if client.UUID.String() == in {
			return client
		}
	}

	return nil
}

func (man *WebSocketManager) findRoomByUUID(in string) *WsRoom {
	for room := range man.rooms {
		if room.UUID.String() == in {
			return room
		}
	}

	return nil
}
