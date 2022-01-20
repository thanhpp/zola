package controller

import (
	"fmt"

	"github.com/google/uuid"
)

type WsRoom struct {
	ID          string    `json:"name"`
	UUID        uuid.UUID `json:"id"`
	Private     bool
	clients     map[*Client]struct{} `json:"-"`
	registerC   chan *Client         `json:"-"`
	unregisterC chan *Client         `json:"-"`
	broadcast   chan []byte          `json:"-"`
}

func NewRoom(id string, private bool) *WsRoom {
	return &WsRoom{
		ID:          id,
		UUID:        uuid.New(),
		Private:     private,
		clients:     make(map[*Client]struct{}),
		registerC:   make(chan *Client),
		unregisterC: make(chan *Client),
		broadcast:   make(chan []byte),
	}
}

func (wr *WsRoom) Run() {
	for {
		select {
		case client := <-wr.registerC:
			wr.Register(client)

		case client := <-wr.unregisterC:
			wr.Unregister(client)

		case msg := <-wr.broadcast:
			wr.BroadcastMessage(msg)
		}
	}
}

func (wr *WsRoom) Register(client *Client) {
	wr.clients[client] = struct{}{}
}

func (wr *WsRoom) Unregister(client *Client) {
	if _, ok := wr.clients[client]; ok {
		delete(wr.clients, client)
	}
}

func (wr *WsRoom) BroadcastMessage(message []byte) {
	for client := range wr.clients {
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

	wr.clients[client] = struct{}{}
}
