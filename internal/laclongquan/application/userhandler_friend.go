package application

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
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

func (u UserHandler) UpdateFriendRequest(ctx context.Context, userAID, userBID string, accept bool) error {
	userA, err := u.repo.GetByID(ctx, userAID)
	if err != nil {
		return err
	}

	userB, err := u.repo.GetByID(ctx, userBID)
	if err != nil {
		return err
	}

	if userA.IsLocked() || userB.IsLocked() {
		return entity.ErrLockedUser
	}

	relation, err := u.relationRepo.GetRelationBetween(ctx, userAID, userBID)
	if err != nil {
		return err
	}

	switch accept {
	case true:
		if err := relation.AcceptFriendRequest(); err != nil {
			return err
		}
		if err := u.relationRepo.UpdateRelation(ctx, relation); err != nil {
			return err
		}
		return nil

	case false:
		if err := relation.RejectFriendRequest(); err != nil {
			return err
		}
		if err := u.relationRepo.DeleteRelation(ctx, relation); err != nil {
			return err
		}
		return nil
	}

	return nil
}
