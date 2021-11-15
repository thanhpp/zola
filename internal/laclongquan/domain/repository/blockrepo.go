package repository

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

var (
	ErrBlockNotFound      = errors.New("block not found")
	ErrUserAlreadyBlocked = errors.New("user already blocked")
)

type BlockRepository interface {
	// read
	Get(ctx context.Context, blocker, blocked string) (*entity.Block, error)
	//write
	Create(ctx context.Context, block *entity.Block) error
	Delete(ctx context.Context, blocker, blocked string) error
}
