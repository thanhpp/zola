package gormdb

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/auco/app"
	"gorm.io/gorm"
)

func (g gormDB) GetLastMessageByRoomID(ctx context.Context, roomID string) (*app.WsMessage, error) {
	var msgDB = new(MessageDB)

	err := g.db.Model(g.msgModel).Where("room_id = ?", roomID).Order("created_at desc").Find(msgDB).Error
	if err != nil {
		return nil, err
	}

	return g.unmarshalMessage(msgDB)
}

func (g gormDB) GetMessagesByRoomID(ctx context.Context, roomID string, offset, limit int) ([]*app.WsMessage, error) {
	var listDB []*MessageDB

	err := g.db.Model(g.msgModel).
		Where("room_id = ?", roomID).
		Order("created_at desc").Offset(offset).Limit(limit).
		Find(&listDB).Error
	if err != nil {
		return nil, err
	}

	var listMsg []*app.WsMessage
	for _, msgDB := range listDB {
		msg, err := g.unmarshalMessage(msgDB)
		if err != nil {
			return nil, err
		}

		listMsg = append(listMsg, msg)
	}

	return listMsg, nil
}

func (g gormDB) GetMessage(ctx context.Context, roomID string, createdAt int64) (*app.WsMessage, error) {
	var msgDB = new(MessageDB)

	err := g.db.Model(g.msgModel).
		Where("room_id = ? AND created_at = ?", roomID, createdAt).
		Take(msgDB).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app.ErrMsgNotFound
		}
		return nil, err
	}

	return g.unmarshalMessage(msgDB)
}
