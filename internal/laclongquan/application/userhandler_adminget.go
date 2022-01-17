package application

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

type AdminGetUserRes struct {
	FriendCount int
	IsFriend    bool
	User        *entity.User
}

func (u UserHandler) AdminGetUser(ctx context.Context, requestorID, requestedID string) (*AdminGetUserRes, error) {
	requestor, err := u.repo.GetByID(ctx, requestorID)
	if err != nil {
		return nil, err
	}

	if !requestor.IsAdmin() {
		return nil, entity.ErrPermissionDenied
	}

	requested, err := u.repo.GetByID(ctx, requestedID)
	if err != nil {
		return nil, err
	}

	friendCount, err := u.relationRepo.CountFriends(ctx, requestedID)
	if err != nil {
		return nil, err
	}

	var relation *entity.Relation
	if !requestor.Equal(requested) {
		relation, err = u.relationRepo.GetRelationBetween(ctx, requestorID, requestedID)
		if err != nil && !errors.Is(err, repository.ErrRelationNotFound) {
			return nil, err
		}
	}

	return &AdminGetUserRes{
		FriendCount: friendCount,
		User:        requested,
		IsFriend:    relation != nil && relation.IsFriend(),
	}, nil
}
