package app

import (
	"context"

	"github.com/thanhpp/zola/internal/auco/infra/adapter/llqclient"
)

type GetConversationRes struct {
	Data []GetConversationElem
}

type GetConversationElem struct {
	Message *WsMessage
	Sender  *llqclient.GetUserInfoResp
}

func (c ConversationHandler) GetByPartnerID(ctx context.Context, requestorID, partnerID string, offset, limit int) (*GetConversationRes, error) {
	room, err := c.roomRepo.FindRoomBetween(requestorID, partnerID)
	if err != nil {
		return nil, err
	}

	isBlock, err := c.llqClient.IsBlock(room.UserA, room.UserB)
	if err != nil {
		return nil, err
	}
	if isBlock.IsBlock {
		return nil, ErrBlocked
	}

	msgs, err := c.msgRepo.GetMessagesByRoomID(ctx, room.ID, offset, limit)
	if err != nil {
		return nil, err
	}

	var (
		userCache = make(map[string]*llqclient.GetUserInfoResp)
		dataElem  = make([]GetConversationElem, 0, len(msgs))
	)

	for i := range msgs {
		userInfo, ok := userCache[msgs[i].SenderID]
		if !ok {
			userInfo, err = c.llqClient.GetUserInfo(msgs[i].SenderID)
			if err != nil {
				return nil, err
			}

			userCache[msgs[i].SenderID] = userInfo
		}
		dataElem = append(dataElem, GetConversationElem{
			Message: msgs[i],
			Sender:  userInfo,
		})
	}

	return &GetConversationRes{
		Data: dataElem,
	}, nil
}

func (c ConversationHandler) GetByRoomID(ctx context.Context, requestorID, roomID string, offset, limit int) (*GetConversationRes, error) {
	room, err := c.roomRepo.GetRoomByID(ctx, roomID)
	if err != nil {
		return nil, err
	}

	msgs, err := c.msgRepo.GetMessagesByRoomID(ctx, room.ID, offset, limit)
	if err != nil {
		return nil, err
	}

	isBlock, err := c.llqClient.IsBlock(room.UserA, room.UserB)
	if err != nil {
		return nil, err
	}
	if isBlock.IsBlock {
		return nil, ErrBlocked
	}

	var (
		userCache = make(map[string]*llqclient.GetUserInfoResp)
		dataElem  = make([]GetConversationElem, 0, len(msgs))
	)

	for i := range msgs {
		userInfo, ok := userCache[msgs[i].SenderID]
		if !ok {
			userInfo, err = c.llqClient.GetUserInfo(msgs[i].SenderID)
			if err != nil {
				return nil, err
			}

			userCache[msgs[i].SenderID] = userInfo
		}
		dataElem = append(dataElem, GetConversationElem{
			Message: msgs[i],
			Sender:  userInfo,
		})
	}

	return &GetConversationRes{
		Data: dataElem,
	}, nil
}
