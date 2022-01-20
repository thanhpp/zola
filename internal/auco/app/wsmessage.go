package app

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

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
	roomID     string
	createdAt  int64
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

func (wm *WsMessage) GetRoomID() string {
	if len(wm.roomID) != 0 {
		return wm.roomID
	}

	return strings.Split(wm.MsgID, "|")[0]
}

func (wm *WsMessage) GetCreatedAtFromID() (int64, error) {
	if wm.createdAt != 0 {
		return wm.createdAt, nil
	}

	idSlice := strings.Split(wm.MsgID, "|")
	if len(idSlice) != 2 {
		return 0, errors.New("invalid message id")
	}

	createdAtInt64, err := strconv.ParseInt(idSlice[1], 10, 64)
	if err != nil {
		return 0, err
	}

	return createdAtInt64, nil
}
