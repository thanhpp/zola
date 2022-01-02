package application

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

type GetUserByIDRes struct {
	FriendCount int
	IsFriend    bool
	User        *entity.User
}

func (u UserHandler) GetUserByID(ctx context.Context, requestorID, requestedID string) (*GetUserByIDRes, error) {
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

	friendCount, err := u.relationRepo.CountFriends(ctx, requestedID)
	if err != nil {
		return nil, err
	}

	return &GetUserByIDRes{
		FriendCount: friendCount,
		IsFriend:    (relation != nil && relation.IsFriend()),
		User:        requested,
	}, nil
}
