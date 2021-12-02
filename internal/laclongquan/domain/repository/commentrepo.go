package repository

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

var (
	ErrCommentNotFound = errors.New("comment not found")
)

type CommentUpdateFunc func(ctx context.Context, comment *entity.Comment) (*entity.Comment, error)
type CommentRepository interface {
	Create(ctx context.Context, comment *entity.Comment) error
	Update(ctx context.Context, commentID, postID string, fn CommentUpdateFunc) error
}
