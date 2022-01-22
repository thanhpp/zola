package app

import (
	"github.com/thanhpp/zola/internal/auco/infra/adapter/llqclient"
)

type ConversationHandler struct {
	roomRepo  RoomRepository
	msgRepo   MessageRepository
	llqClient *llqclient.LacLongQuanClient
}
