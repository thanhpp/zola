package gormdb

import (
	"context"

	"github.com/thanhpp/zola/internal/auco/app"
)

func (g gormDB) GetLastMessageByRoomID(ctx context.Context, roomID string) (*app.WsMessage, error) {
	var msgDB = new(MessageDB)

	err := g.db.Model(g.msgModel).Where("room_id = ?", roomID).Order("created_at desc").Find(msgDB).Error
	if err != nil {
		return nil, err
	}

	return g.unmarshalMessage(msgDB)
}
