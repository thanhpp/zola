package app

import (
	"context"
	"strconv"
	"strings"

	"github.com/thanhpp/zola/pkg/logger"
)

func (c ConversationHandler) DeleteMessage(ctx context.Context, requestorID, msgID string) error {
	msgIDSlice := strings.Split(msgID, "|")
	if len(msgIDSlice) != 2 {
		logger.Errorf("ConversationHandler - invalid msgID len: %d", len(msgIDSlice))
		return ErrInvalidMsgID
	}

	createdAtStr := msgIDSlice[1]
	createdAtInt64, err := strconv.ParseInt(createdAtStr, 10, 64)
	if err != nil {
		logger.Errorf("ConversationHandler - invalid msgID createdAt: %v", err)
		return ErrInvalidMsgID
	}

	msg, err := c.msgRepo.GetMessage(ctx, msgIDSlice[0], createdAtInt64)
	if err != nil {
		return err
	}

	if msg.SenderID != requestorID {
		return ErrNotSender
	}

	return c.msgRepo.DeleteMessage(ctx, msgIDSlice[0], createdAtInt64)
}
