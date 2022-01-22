package app

import "context"

type GetListConversationRes struct {
	Data      []*GetListConversationElem
	NumNewMsg int
}

type GetListConversationElem struct {
	Room        *WsRoom
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

		dataElem = append(dataElem, &GetListConversationElem{
			Room:        rooms[i],
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
