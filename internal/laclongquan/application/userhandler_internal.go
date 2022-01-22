package application

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

func (u UserHandler) InternalGetUser(ctx context.Context, userID string) (*entity.User, error) {
	return u.repo.GetByID(ctx, userID)
}

func (u UserHandler) InternalIsBlock(ctx context.Context, userAID, userBID string) (bool, error) {
	userA, err := u.repo.GetByID(ctx, userAID)
	if err != nil {
		return false, err
	}

	if userA.ID().String() == userBID {
		return false, nil
	}

	userB, err := u.repo.GetByID(ctx, userBID)
	if err != nil {
		return false, err
	}

	var relation *entity.Relation
	relation, err = u.relationRepo.GetRelationBetween(ctx, userA.ID().String(), userB.ID().String())
	if err != nil {
		if errors.Is(err, repository.ErrRelationNotFound) {
			return false, nil
		}

		return false, err
	}

	return relation.IsBlock(), nil
}
