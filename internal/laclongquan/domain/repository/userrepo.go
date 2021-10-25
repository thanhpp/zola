package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

var (
	ErrDuplicateUser = errors.New("duplicate user")
)

type UserUpdateFunc func(ctx context.Context, user *entity.User) (*entity.User, error)

type UserRepository interface {
	// read
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByPhone(ctx context.Context, phone string) (*entity.User, error)

	// write
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, id string, fn UserUpdateFunc) error

	// delete
	DeleteByID(ctx context.Context, id string) error
}
