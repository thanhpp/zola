package gormdb

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
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

func (r relationGorm) GetRelationBetween(ctx context.Context, userIDA, userIDB string) (*entity.Relation, error) {
	var relationDB = new(RelationDB)

	if err := r.db.WithContext(ctx).Model(r.model).
		Where("(user_a = ? AND user_b = ?) OR (user_a = ? AND user_b = ?)", userIDA, userIDB, userIDB, userIDA).
		Take(relationDB).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrRelationNotFound
		}
		return nil, err
	}

	return r.unmarshal(relationDB)
}

func (r relationGorm) CreateRelation(ctx context.Context, relation *entity.Relation) error {
	if relation == nil {
		return errors.New("nil input")
	}

	relationDB, err := r.marshal(relation)
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Model(r.model).Create(relationDB).Error
}

func (r relationGorm) UpdateRelation(ctx context.Context, relation *entity.Relation) error {
	if relation == nil {
		return errors.New("nil input")
	}

	relationDB, err := r.marshal(relation)
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Model(r.model).
		Where("user_a = ? AND user_b = ?", relationDB.UserA, relationDB.UserB).
		Updates(relationDB).Error
}

func (r relationGorm) DeleteRelation(ctx context.Context, relation *entity.Relation) error {
	relationDB, err := r.marshal(relation)
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Model(r.model).
		Where("user_a = ? AND user_b = ?", relationDB.UserA, relationDB.UserB).
		Delete(relation).Error
}
