package gormdb

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
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

func (c commentGorm) GetByIDAndPostID(ctx context.Context, commentID, postID string) (*entity.Comment, error) {
	return c.getByPostIDCommentID(ctx, c.db, postID, commentID)
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

	post, err := c.postGorm.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	user, err := c.userGorm.GetByID(ctx, cmtDB.CreatorUUID)
	if err != nil {
		return nil, err
	}

	return c.unmarshal(cmtDB, post, user), nil
}

func (c commentGorm) CountByPostID(ctx context.Context, postID string) (int, error) {
	var commentCount int64
	if err := c.db.WithContext(ctx).Model(c.cmtModel).
		Where("post_uuid = ?", postID).Count(&commentCount).
		Error; err != nil {
		return 0, err
	}

	return int(commentCount), nil
}

func (c commentGorm) GetByPostIDFromNonBlockedActiveUser(ctx context.Context, requestorID, postID string, offset, limit int) ([]*entity.Comment, error) {
	var list []*CommentDB

	stmt := c.db.WithContext(ctx).Model(c.cmtModel).
		Select("comment_db.*").
		Where("post_uuid = ?", postID).
		Order("created_at desc").
		Joins(`JOIN user_db ON user_db.user_uuid = comment_db.creator_uuid AND user_db.state = 'active'
		LEFT OUTER JOIN relation_db ON (
			(
				(relation_db.user_a = ? AND relation_db.user_b = comment_db.creator_uuid) 
				OR
				(relation_db.user_a = comment_db.creator_uuid AND relation_db.user_b = ?)
				AND
				(relation_db.status = ?)
			)
			OR 
			(
				comment_db.creator_uuid = ?
			)
		)`, requestorID, requestorID, entity.RelationFriend, requestorID)

	if offset >= 0 && limit > 0 {
		stmt = stmt.Offset(offset).Limit(limit)
	}

	if err := stmt.Find(&list).Error; err != nil {
		return nil, err
	}

	post, err := c.postGorm.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	var (
		comments  = make([]*entity.Comment, 0, len(list))
		userCache = make(map[string]*entity.User)
	)
	for i := range list {
		user, ok := userCache[list[i].CreatorUUID]
		if ok {
			comments = append(comments, c.unmarshal(list[i], post, user))
			continue
		}
		user, err = c.userGorm.GetByID(ctx, list[i].CreatorUUID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c.unmarshal(list[i], post, user))
		userCache[list[i].CreatorUUID] = user
	}

	return comments, nil
}

func (c commentGorm) Create(ctx context.Context, comment *entity.Comment) error {
	return c.db.WithContext(ctx).Model(c.cmtModel).
		Create(c.marshal(comment)).Error
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

func (c commentGorm) Delete(ctx context.Context, comment *entity.Comment) error {
	return c.db.WithContext(ctx).Model(c.cmtModel).
		Where("comment_uuid = ?", comment.IDString()).
		Delete(c.cmtModel).Error
}
