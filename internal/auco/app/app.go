package app

type App struct {
	ConversationHandler ConversationHandler
}

func NewApp(roomRepo RoomRepository, msgRepo MessageRepository) *App {
	return &App{
		ConversationHandler: ConversationHandler{
			roomRepo: roomRepo,
			msgRepo:  msgRepo,
		},
	}
}
