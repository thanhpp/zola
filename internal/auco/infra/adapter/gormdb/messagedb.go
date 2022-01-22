package gormdb

import (
	"strconv"

	"github.com/thanhpp/zola/internal/auco/app"
)

type MessageDB struct {
	RoomID     string `gorm:"Column:room_id; Type:text; primary_key"`
	CreatedAt  int64  `gorm:"Column:created_at; Type:bigint; primary_key"`
	SenderID   string `gorm:"Column:sender_id; Type:text"`
	ReceiverID string `gorm:"Column:receiver_id; Type:text"`
	Content    string `gorm:"Column:content; Type:text"`
	DeletedAt  int64  `gorm:"Column:deleted_at; Type:bigint"`
	Seen       bool   `gorm:"Column:seen; Type:boolean; Default:false"`
}

func (g gormDB) marshalMessage(msg *app.WsMessage) (*MessageDB, error) {
	createdAt, err := msg.GetCreatedAtFromID()
	if err != nil {
		return nil, err
	}

	return &MessageDB{
		RoomID:     msg.GetRoomID(),
		CreatedAt:  createdAt,
		SenderID:   msg.SenderID,
		ReceiverID: msg.ReceiverID,
		Content:    msg.Content,
		DeletedAt:  0,
		Seen:       msg.IsSeen(),
	}, nil
}

func (g gormDB) unmarshalMessage(msgDB *MessageDB) (*app.WsMessage, error) {
	msg, err := g.fac.NewMessage(msgDB.RoomID, msgDB.SenderID, msgDB.ReceiverID, strconv.FormatInt(msgDB.CreatedAt, 10), msgDB.Content, msgDB.Seen)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (g gormDB) CreateMessage(msg *app.WsMessage) error {
	msgDB, err := g.marshalMessage(msg)
	if err != nil {
		return err
	}

	return g.db.Model(g.msgModel).Create(msgDB).Error
}
