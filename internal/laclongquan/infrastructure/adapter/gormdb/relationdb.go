package gormdb

import (
	"time"

	"github.com/pkg/errors"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"gorm.io/gorm"
)

type RelationDB struct {
	UserA     string `gorm:"Column:user_a; Type:text; primaryKey"`
	UserB     string `gorm:"Column:user_b; Type:text; primaryKey"`
	Status    string `gorm:"Column:status; Type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type relationGorm struct {
	db    *gorm.DB
	model *RelationDB
}

func (r relationGorm) marshal(relation *entity.Relation) (*RelationDB, error) {
	if relation == nil {
		return nil, errors.New("nil input")
	}

	return &RelationDB{
		UserA:  relation.UserAIDStr(),
		UserB:  relation.UserBIDStr(),
		Status: relation.Status.String(),
	}, nil
}

func (r relationGorm) unmarshal(relationDB *RelationDB) (*entity.Relation, error) {
	if relationDB == nil {
		return nil, errors.New("nil input")
	}

	return entity.NewRelationFromDB(
		relationDB.UserA,
		relationDB.UserB,
		relationDB.Status,
	)
}
