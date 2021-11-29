package application

import (
	"context"
	"errors"
)

var (
	ErrNotFriend = errors.New("not friend")
)

func (p PostHandler) CreateComment(ctx context.Context, postID, creatorID, content string) error {
	post, err := p.repo.GetByID(ctx, postID)
	if err != nil {
		return err
	}

	if post.Creator() != creatorID {
		relation, err := p.relationRepo.GetRelationBetween(ctx, creatorID, post.CreatorUUID().String())
		if err != nil {
			return err
		}

		if relation.IsBlock() {
			return ErrAlreadyBlocked
		}

		if !relation.IsFriend() {
			return ErrNotFriend
		}
	}

	creator, err := p.userRepo.GetByID(ctx, creatorID)
	if err != nil {
		return err
	}

	comment, err := p.fac.NewComment(content, post, creator)
	if err != nil {
		return err
	}

	return p.commentRepo.Create(ctx, comment)
}
