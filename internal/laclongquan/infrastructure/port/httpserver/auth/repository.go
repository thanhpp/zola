package auth

import "context"

type Repository interface {
	CheckByID(ctx context.Context, id string) error
	Cache(ctx context.Context, claims *Claims) error
	Delete(ctx context.Context, id string) error
	DeleteByUserID(ctx context.Context, userID string) error
	DeleteExpired(ctx context.Context) error
}
