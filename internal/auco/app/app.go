package app

import "github.com/thanhpp/zola/internal/auco/infra/adapter/llqclient"

type App struct {
	ConversationHandler ConversationHandler
}

func NewApp(roomRepo RoomRepository, msgRepo MessageRepository, llqClient *llqclient.LacLongQuanClient) *App {
	return &App{
		ConversationHandler: ConversationHandler{
			roomRepo:  roomRepo,
			msgRepo:   msgRepo,
			llqClient: llqClient,
		},
	}
}
