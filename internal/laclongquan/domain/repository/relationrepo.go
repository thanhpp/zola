package repository

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

var (
	ErrRelationNotFound = errors.New("relation not found")
)

type RelationRepository interface {
	// read
	GetRelationBetween(ctx context.Context, userIDA, userIDB string) (*entity.Relation, error)

	// write
	CreateRelation(ctx context.Context, relation *entity.Relation) error
}
