package app

import (
	"context"

	"github.com/thanhpp/zola/internal/auco/infra/adapter/llqclient"
)

type GetListConversationRes struct {
	Data      []*GetListConversationElem
	NumNewMsg int
}

type GetListConversationElem struct {
	Room        *WsRoom
	UserInfo    *llqclient.GetUserInfoResp
	LastMessage *WsMessage
}

func (c ConversationHandler) GetListConversation(ctx context.Context, requestorID string, offset, limit int) (*GetListConversationRes, error) {

	rooms, err := c.roomRepo.GetListRoom(ctx, requestorID, offset, limit)
	if err != nil {
		return nil, err
	}

	var (
		numNewMsg int
		dataElem  = make([]*GetListConversationElem, 0, len(rooms))
	)

	for i := range rooms {
		lastMessage, err := c.msgRepo.GetLastMessageByRoomID(ctx, rooms[i].ID)
		if err != nil {
			return nil, err
		}

		// get userinfo
		var userID string
		if requestorID != rooms[i].UserA {
			userID = rooms[i].UserA
		} else {
			userID = rooms[i].UserB
		}
		userInfo, err := c.llqClient.GetUserInfo(userID)
		if err != nil {
			return nil, err
		}

		dataElem = append(dataElem, &GetListConversationElem{
			Room:        rooms[i],
			UserInfo:    userInfo,
			LastMessage: lastMessage,
		})

		if lastMessage != nil && lastMessage.SenderID != requestorID && !lastMessage.IsSeen() {
			numNewMsg++
		}
	}

	return &GetListConversationRes{
		Data:      dataElem,
		NumNewMsg: numNewMsg,
	}, nil
}
