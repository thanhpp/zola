package gormdb

import (
	"time"

	"github.com/thanhpp/zola/internal/auco/app"
	"github.com/thanhpp/zola/pkg/logger"
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
