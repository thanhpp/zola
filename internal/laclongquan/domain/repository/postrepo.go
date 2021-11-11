package repository

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

var (
	ErrPostNotFound = errors.New("post not found")
)

type PostUpdateFn func(ctx context.Context, post *entity.Post) (*entity.Post, error)

type PostRepository interface {
	// read
	GetByID(ctx context.Context, id string) (*entity.Post, error)

	// write
	Create(ctx context.Context, post *entity.Post) error
	Update(ctx context.Context, id string, fn PostUpdateFn) error

	// delete
	Delete(ctx context.Context, id string) error
}
