package auth

import "context"

type Repository interface {
	CheckByID(ctx context.Context, id string) error
	Cache(ctx context.Context, claims *Claims) error
	Delete(ctx context.Context, id string) error
	DeleteExpired(ctx context.Context) error
}
