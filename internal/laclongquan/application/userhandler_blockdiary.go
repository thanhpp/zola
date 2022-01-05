package application

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
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
		// if a relation not found, create a new one
		if errors.Is(err, repository.ErrRelationNotFound) {
			relation, err = u.fac.NewDiaryBlockRelation(blocker, blocked)
			if err != nil {
				return err
			}

			if err := u.relationRepo.CreateRelation(ctx, relation); err != nil {
				return err
			}

			return nil
		}
		return err
	}

	// if a relation exists, then there are 2 different users
	if err := relation.SetDiaryBlock(blocker, blocked); err != nil {
		return err
	}

	return u.relationRepo.UpdateRelation(ctx, relation)
}
