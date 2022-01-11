package repository

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

var (
	ErrRelationNotFound = errors.New("relation not found")
	ErrSameUser         = errors.New("same user")
)

type RelationRepository interface {
	// read
	GetRelationBetween(ctx context.Context, userIDA, userIDB string) (*entity.Relation, error)
	CountFriends(ctx context.Context, userID string) (int, error)
	GetActiveRequestedFriends(ctx context.Context, userID string, offset, limit int) ([]*entity.Relation, error)
	GetActiveUserFriends(ctx context.Context, userID string, offset, limit int) ([]*entity.Relation, int, error)

	// write
	CreateRelation(ctx context.Context, relation *entity.Relation) error
	UpdateRelation(ctx context.Context, relation *entity.Relation) error
	DeleteRelation(ctx context.Context, relation *entity.Relation) error
}
