package repository

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

type LikeRepository interface {
	CreateOrDelete(ctx context.Context, like *entity.Like) error
	Count(ctx context.Context, postID string) (int, error)
}
