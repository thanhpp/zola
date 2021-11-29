package gormdb

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"gorm.io/gorm"
)

type CommentDB struct {
	CommentUUID string `gorm:"Column:comment_uuid; type:text; primaryKey"`
	PostUUID    string `gorm:"Column:post_uuid; type:text"`
	CreatorUUID string `gorum:"Column:creator_uuid; type:text"`
	Content     string `gorm:"Column:content; type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type commentGorm struct {
	db       *gorm.DB
	cmtModel *CommentDB
	postGorm postGorm
	userGorm userGorm
}

func (c commentGorm) marshal(comment *entity.Comment) *CommentDB {
	return &CommentDB{
		CommentUUID: comment.IDString(),
		PostUUID:    comment.GetPost().ID(),
		CreatorUUID: comment.GetCreator().ID().String(),
		Content:     comment.GetContent(),
		CreatedAt:   comment.GetCreatedAt(),
	}
}

func (c commentGorm) unmarshal(commentDB *CommentDB, postDB *PostDB, userDB *UserDB) (*entity.Comment, error) {
	post, err := c.postGorm.unmarshalPost(postDB)
	if err != nil {
		return nil, err
	}

	user, err := c.userGorm.unmarshalUser(userDB)
	if err != nil {
		return nil, err
	}

	return &entity.Comment{
		ID:        uuid.MustParse(commentDB.CommentUUID),
		Content:   commentDB.Content,
		Post:      post,
		Creator:   user,
		CreatedAt: commentDB.CreatedAt,
	}, nil
}

func (c commentGorm) Create(ctx context.Context, comment *entity.Comment) error {
	return c.db.WithContext(ctx).Model(c.cmtModel).
		Create(c.marshal(comment)).Error
}
