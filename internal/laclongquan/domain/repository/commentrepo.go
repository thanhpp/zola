package repository

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *entity.Comment) error
}
