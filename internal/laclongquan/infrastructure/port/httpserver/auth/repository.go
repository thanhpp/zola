package auth

import "context"

type Repository interface {
	GetByID(ctx context.Context, id string) (*JWT, error)

	Cache(ctx context.Context, jwt *JWT) error

	Delete(ctx context.Context, id string) error
}
