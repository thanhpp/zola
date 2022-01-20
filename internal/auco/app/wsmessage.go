package app

import (
	"encoding/json"
	"errors"

	"github.com/thanhpp/zola/pkg/logger"
)

var (
	ErrNilMessage = errors.New("nil message")
)

type WsMessage struct {
	MsgID      string `json:"message_id"`
	Event      string `json:"event"`
	SenderID   string `json:"sender"`
	ReceiverID string `json:"receiver"`
	Created    string `json:"created"`
	Content    string `json:"content"`
}

func (wm *WsMessage) Encode() []byte {
	msgB, err := json.Marshal(wm)
	if err != nil {
		logger.Errorf("WsMessage: encode error: %v", err)
		return nil
	}

	return msgB
}

func (wm *WsMessage) Decode(msgB []byte) error {
	if wm == nil || msgB == nil {
		return ErrNilMessage
	}

	return json.Unmarshal(msgB, wm)
}
