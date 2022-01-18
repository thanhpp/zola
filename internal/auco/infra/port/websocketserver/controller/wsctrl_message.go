package controller

import (
	"encoding/json"

	"github.com/thanhpp/zola/pkg/logger"
)

const (
	MessageActionSend  = "send-message"
	MessageActionJoin  = "join-room"
	MessageActionLeave = "leave-room"
)

type WsMessage struct {
	Action  string  `json:"action"`
	Message string  `json:"message"`
	Target  string  `json:"target"`
	Sender  *Client `json:"sender"`
}

func (msg *WsMessage) Encode() []byte {
	msgB, err := json.Marshal(msg)
	if err != nil {
		logger.Errorf("WsMessage: encode error: %v", err)
		return nil
	}

	return msgB
}
