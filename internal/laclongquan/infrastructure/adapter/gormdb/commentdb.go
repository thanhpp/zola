package gormdb

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
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

func (c commentGorm) unmarshal(commentDB *CommentDB, post *entity.Post, user *entity.User) *entity.Comment {
	return &entity.Comment{
		ID:        uuid.MustParse(commentDB.CommentUUID),
		Content:   commentDB.Content,
		Post:      post,
		Creator:   user,
		CreatedAt: commentDB.CreatedAt,
	}
}

func (c commentGorm) Create(ctx context.Context, comment *entity.Comment) error {
	return c.db.WithContext(ctx).Model(c.cmtModel).
		Create(c.marshal(comment)).Error
}

func (c commentGorm) getByPostIDCommentID(ctx context.Context, tx *gorm.DB, postID, commentID string) (*entity.Comment, error) {
	var cmtDB = new(CommentDB)

	err := tx.WithContext(ctx).Model(c.cmtModel).
		Where("post_uuid = ? AND comment_uuid = ?", postID, commentID).
		Take(cmtDB).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrCommentNotFound
		}
		return nil, err
	}

	user, err := c.userGorm.GetByID(ctx, cmtDB.CreatorUUID)
	if err != nil {
		return nil, err
	}

	post, err := c.postGorm.GetByID(ctx, cmtDB.PostUUID)
	if err != nil {
		return nil, err
	}

	return c.unmarshal(cmtDB, post, user), nil
}

func (c commentGorm) Update(ctx context.Context, postID, commentID string, fn repository.CommentUpdateFunc) error {
	return c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		comment, err := c.getByPostIDCommentID(ctx, tx, postID, commentID)
		if err != nil {
			return nil
		}

		comment, err = fn(ctx, comment)
		if err != nil {
			return err
		}

		return tx.WithContext(ctx).Model(c.cmtModel).
			Where("comment_uuid = ?", commentID).
			Save(c.marshal(comment)).Error
	})
}
