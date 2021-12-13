package gormdb

import (
	"context"
	"time"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"gorm.io/gorm"
)

type LikeDB struct {
	PostUUID    string    `gorm:"Column:post_uuid; Type:text; primaryKey"`
	CreatorUUID string    `gorm:"Column:creator_uuid; Type:text; primaryKey"`
	CreatedAt   time.Time `gorm:"Column:created_at; Type:timestamp"`
}

type likeGorm struct {
	db    *gorm.DB
	model *LikeDB
}

func (l likeGorm) marshal(like *entity.Like) *LikeDB {
	return &LikeDB{
		PostUUID:    like.PostID,
		CreatorUUID: like.Creator,
	}
}

func (l likeGorm) unmarshal(likeDB *LikeDB) *entity.Like {
	return &entity.Like{
		PostID:  likeDB.PostUUID,
		Creator: likeDB.CreatorUUID,
	}
}

func (l likeGorm) Count(ctx context.Context, postID string) (int, error) {
	var count = new(int64)

	if err := l.db.WithContext(ctx).Model(l.model).
		Where("post_uuid = ?", postID).
		Count(count).Error; err != nil {
		return 0, err
	}

	return int(*count), nil
}

func (l likeGorm) IsLiked(ctx context.Context, userID, postID string) bool {
	var count = new(int64)

	if err := l.db.WithContext(ctx).Model(l.model).
		Where("post_uuid = ? AND creator_uuid = ?", postID, userID).
		Count(count).Error; err != nil {
		return false
	}

	return *count > 0
}

func (l likeGorm) CreateOrDelete(ctx context.Context, like *entity.Like) error {
	likeDB := l.marshal(like)

	if err := l.db.WithContext(ctx).Model(l.model).Create(likeDB).Error; err != nil {
		if isDuplicate(err) {
			if err := l.db.WithContext(ctx).Model(l.model).
				Where("post_uuid = ? AND creator_uuid = ?", like.PostID, like.Creator).
				Delete(l.model).Error; err != nil {
				return err
			}
			return nil
		}

		return err
	}

	return nil
}
