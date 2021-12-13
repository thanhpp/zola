package repository

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

type LikeRepository interface {
	// read
	Count(ctx context.Context, postID string) (int, error)
	IsLiked(ctx context.Context, userID, postID string) bool

	// write
	CreateOrDelete(ctx context.Context, like *entity.Like) error
}
