package controller

import "github.com/thanhpp/zola/internal/auco/app"

type ConversationController struct {
	conversationHandler *app.ConversationHandler
}

func NewConversationController(ch *app.ConversationHandler) *ConversationController {
	return &ConversationController{
		conversationHandler: ch,
	}
}
