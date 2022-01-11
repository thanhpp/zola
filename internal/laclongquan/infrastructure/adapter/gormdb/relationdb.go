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
		relationDB.CreatedAt,
	)
}

func (r relationGorm) GetRelationBetween(ctx context.Context, userIDA, userIDB string) (*entity.Relation, error) {
	if userIDA == userIDB {
		return nil, repository.ErrSameUser
	}

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

func (r relationGorm) CountFriends(ctx context.Context, userID string) (int, error) {
	var count int64

	if err := r.db.WithContext(ctx).Model(r.model).
		Where("user_a = ? OR user_b = ? AND status = ?",
			userID, userID, entity.RelationFriend).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r relationGorm) GetActiveRequestedFriends(ctx context.Context, userID string, offset, limit int) ([]*entity.Relation, error) {
	var list []*RelationDB

	if err := r.db.WithContext(ctx).Model(r.model).
		Where("user_b = ? AND status = ?", userID, entity.RelationRequesting).
		Order("created_at desc").
		Joins("JOIN user_db ON user_db.user_uuid = user_a AND user_db.state = 'active'").
		Offset(offset).Limit(limit).
		Find(&list).Error; err != nil {
		return nil, err
	}

	var relations []*entity.Relation
	for _, relationDB := range list {
		relation, err := r.unmarshal(relationDB)
		if err != nil {
			return nil, err
		}
		relations = append(relations, relation)
	}

	return relations, nil
}

func (r relationGorm) GetActiveUserFriends(ctx context.Context, userID string, offset, limit int) ([]*entity.Relation, int, error) {
	var (
		list  []*RelationDB
		total = new(int64)
	)
	stmt := r.db.WithContext(ctx).Model(r.model).
		Where("(user_a = ? OR user_b = ?) AND status = ?", userID, userID, entity.RelationFriend).
		Order("created_at desc").
		Joins(`
			LEFT JOIN user_db 
				ON (user_db.user_uuid = user_a OR user_db.user_uuid = user_b)
					AND user_db.state = ? AND user_db.user_uuid <> ?`, entity.UserStateActive, userID)

	if err := stmt.Count(total).Error; err != nil {
		return nil, -1, err
	}

	if err := stmt.Offset(offset).Limit(limit).
		Find(&list).Error; err != nil {
		return nil, -1, err
	}

	var relations []*entity.Relation
	for _, relationDB := range list {
		relation, err := r.unmarshal(relationDB)
		if err != nil {
			return nil, -1, err
		}
		relations = append(relations, relation)
	}

	return relations, int(*total), nil
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
