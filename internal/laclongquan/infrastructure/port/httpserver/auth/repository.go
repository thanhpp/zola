package auth

import "context"

type Repository interface {
	GetByID(ctx context.Context, id string) (*Claims, error)

	Cache(ctx context.Context, jwt *Claims) error

	Delete(ctx context.Context, id string) error
}
