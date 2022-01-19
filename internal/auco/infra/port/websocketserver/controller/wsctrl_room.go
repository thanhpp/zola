package controller

import (
	"fmt"

	"github.com/google/uuid"
)

type WsRoom struct {
	ID          string
	UUID        uuid.UUID
	Private     bool
	Clients     map[*Client]struct{}
	RegisterC   chan *Client
	UnregisterC chan *Client
	Broadcast   chan []byte
}

func NewRoom(id string, private bool) *WsRoom {
	return &WsRoom{
		ID:          id,
		UUID:        uuid.New(),
		Private:     private,
		Clients:     make(map[*Client]struct{}),
		RegisterC:   make(chan *Client),
		UnregisterC: make(chan *Client),
		Broadcast:   make(chan []byte),
	}
}

func (wr *WsRoom) Run() {
	for {
		select {
		case client := <-wr.RegisterC:
			wr.Register(client)

		case client := <-wr.UnregisterC:
			wr.Unregister(client)

		case msg := <-wr.Broadcast:
			wr.BroadcastMessage(msg)
		}
	}
}

func (wr *WsRoom) Register(client *Client) {
	wr.Clients[client] = struct{}{}
}

func (wr *WsRoom) Unregister(client *Client) {
	if _, ok := wr.Clients[client]; ok {
		delete(wr.Clients, client)
	}
}

func (wr *WsRoom) BroadcastMessage(message []byte) {
	for client := range wr.Clients {
		client.send <- message
	}
}

const (
	welcomeMsg = "%s joined the room"
)

func (wr *WsRoom) NotifyClients(newClient *Client) {
	newMsg := &WsMessage{
		Action:  MessageActionSend,
		Target:  wr,
		Message: fmt.Sprintf(welcomeMsg, newClient.ID),
	}

	wr.BroadcastMessage(newMsg.Encode())
}

func (wr *WsRoom) registerClientInRoom(client *Client) {
	if !wr.Private {
		wr.NotifyClients(client)
	}

	wr.Clients[client] = struct{}{}
}
