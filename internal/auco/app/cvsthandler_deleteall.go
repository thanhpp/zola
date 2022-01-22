package app

import "context"

func (c ConversationHandler) DeleteByConversationID(ctx context.Context, requestorID, conversationID string) error {
	room, err := c.roomRepo.GetRoomByID(ctx, conversationID)
	if err != nil {
		return err
	}

	if requestorID != room.UserA && requestorID != room.UserB {
		return ErrInvalidUser
	}

	if err := c.msgRepo.DeleteByRoomID(ctx, room.ID); err != nil {
		return err
	}

	return nil
}
