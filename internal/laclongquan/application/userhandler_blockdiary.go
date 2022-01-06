package application

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

func (u UserHandler) BlockDiary(ctx context.Context, blockerID, blockedID string) error {
	blocker, err := u.repo.GetByID(ctx, blockerID)
	if err != nil {
		return err
	}

	blocked, err := u.repo.GetByID(ctx, blockedID)
	if err != nil {
		return err
	}

	var relation *entity.Relation
	if relation, err = u.relationRepo.GetRelationBetween(ctx, blockerID, blockedID); err != nil {
		return err
	}
	//logger.Debugf("relation between %s and %s is %s", relation.UserAIDStr(), relation.UserBIDStr(), relation.Status)

	if err := relation.BlockDiary(blocker, blocked); err != nil {
		return err
	}

	return u.relationRepo.UpdateRelation(ctx, relation)
}

func (u UserHandler) UnblockDiary(ctx context.Context, blockerID, blockedID string) error {
	blocker, err := u.repo.GetByID(ctx, blockerID)
	if err != nil {
		return err
	}

	blocked, err := u.repo.GetByID(ctx, blockedID)
	if err != nil {
		return err
	}

	var relation *entity.Relation
	if relation, err = u.relationRepo.GetRelationBetween(ctx, blockerID, blockedID); err != nil {
		return err
	}

	if err := relation.UnblockDiary(blocker, blocked); err != nil {
		return err
	}

	return u.relationRepo.UpdateRelation(ctx, relation)
}
