package repository

import (
	"context"
	"errors"
	"time"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

var (
	ErrPostNotFound  = errors.New("post not found")
	ErrMediaNotFound = errors.New("media not found")
)

type PostUpdateFn func(ctx context.Context, post *entity.Post) (*entity.Post, error)

type PostRepository interface {
	// read
	GetByID(ctx context.Context, id string) (*entity.Post, error)
	GetMediaByID(ctx context.Context, id string) (*entity.Media, error)
	GetListPost(ctx context.Context, requestorID string, timeMilestone time.Time, offset, limit int) ([]*entity.Post, int, error)
	GetListPostForAdmin(ctx context.Context, requestorID string, timeMileStone time.Time, offset, limit int) ([]*entity.Post, int, error)

	// write
	Create(ctx context.Context, post *entity.Post) error
	Update(ctx context.Context, id string, fn PostUpdateFn) error

	// delete
	Delete(ctx context.Context, id string) error
}
