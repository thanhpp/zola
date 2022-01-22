package gormdb

import (
	"context"
	"errors"
	"time"

	"github.com/thanhpp/zola/internal/auco/app"
	"github.com/thanhpp/zola/pkg/logger"
	"gorm.io/gorm"
)

type RoomDB struct {
	ID        string
	UserA     string
	UserB     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (g gormDB) marshalRoom(room *app.WsRoom) *RoomDB {
	return &RoomDB{
		ID:    room.ID,
		UserA: room.UserA,
		UserB: room.UserB,
	}
}

func (g gormDB) unmarshalRoom(roomDB *RoomDB) *app.WsRoom {
	room, err := g.fac.NewRoomFromDB(roomDB.ID, roomDB.UserA, roomDB.UserB)
	if err != nil {
		logger.Errorf("gormDB - unmarshalRoom: %v", err)
	}
	return room
}

func (g gormDB) GetRoomByID(ctx context.Context, roomID string) (*app.WsRoom, error) {
	var roomDB = new(RoomDB)

	err := g.db.Model(g.roomModel).WithContext(ctx).Where("id = ?", roomID).Take(roomDB).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app.ErrRoomNotFound
		}
		return nil, err
	}

	return g.unmarshalRoom(roomDB), nil
}

func (g gormDB) GetListRoom(ctx context.Context, userID string, offset, limit int) ([]*app.WsRoom, error) {
	var listRoom []*RoomDB

	if err := g.db.Model(g.roomModel).WithContext(ctx).
		Where(`(user_a = ? OR user_b = ?)`, userID, userID).
		Order("created_at desc").
		Offset(offset).Limit(limit).
		Find(&listRoom).Error; err != nil {
		return nil, err
	}

	var listRoomApp = make([]*app.WsRoom, 0, len(listRoom))
	for _, roomDB := range listRoom {
		listRoomApp = append(listRoomApp, g.unmarshalRoom(roomDB))
	}

	return listRoomApp, nil
}

func (g gormDB) FindRoomBetween(userA, userB string) (*app.WsRoom, error) {
	var roomDB = new(RoomDB)

	err := g.db.Model(g.roomModel).Where(`(user_a = ? AND user_b = ?) OR (user_b = ? AND user_a = ?)`, userA, userB, userA, userB).Take(roomDB).Error
	if err != nil {
		return nil, err
	}

	return g.unmarshalRoom(roomDB), nil
}

func (g gormDB) CreateRoom(room *app.WsRoom) error {
	return g.db.Model(g.roomModel).Create(g.marshalRoom(room)).Error
}
