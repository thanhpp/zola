package application

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

var (
	ErrRelationExisted = errors.New("relation existed")
)

func (u UserHandler) NewFriendRequest(ctx context.Context, requestorID, requesteeID string) error {
	relation, err := u.relationRepo.GetRelationBetween(ctx, requestorID, requesteeID)
	if err != nil {
		if !errors.Is(err, repository.ErrRelationNotFound) {
			return err
		}
	}

	if relation != nil {
		return ErrRelationExisted
	}

	requestor, err := u.repo.GetByID(ctx, requestorID)
	if err != nil {
		return err
	}

	requestee, err := u.repo.GetByID(ctx, requesteeID)
	if err != nil {
		return err
	}

	newRelation, err := u.fac.NewFriendRequest(requestor, requestee)
	if err != nil {
		return err
	}

	return u.relationRepo.CreateRelation(ctx, newRelation)
}
