package application

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

func (u UserHandler) GetUserByID(ctx context.Context, requestorID, requestedID string) (*entity.User, error) {
	requestor, err := u.repo.GetByID(ctx, requestorID)
	if err != nil {
		return nil, err
	}

	if requestor.IsLocked() {
		return nil, entity.ErrLockedUser
	}

	requested, err := u.repo.GetByID(ctx, requestedID)
	if err != nil {
		return nil, err
	}

	var (
		relation *entity.Relation
	)
	if !requested.Equal(requestor) {
		relation, err = u.relationRepo.GetRelationBetween(ctx, requestor.ID().String(), requested.ID().String())
		if err != nil && !errors.Is(err, repository.ErrRelationNotFound) {
			return nil, err
		}
	}
	if err := requested.CanGetUserInfo(requestor, relation); err != nil {
		return nil, err
	}

	return requested, nil
}
