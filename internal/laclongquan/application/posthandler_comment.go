package application

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
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

func (p PostHandler) UpdateComment(ctx context.Context, updaterID, postID, commentID, content string) error {
	// logger.Debugf("handler - post id %v", postID)
	// logger.Debugf("hanlder - comment id %v", commentID)

	return p.commentRepo.Update(ctx, postID, commentID, func(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {
		// get the relation if the updater is not the creator of the post
		if comment.Creator.ID().String() != updaterID {
			relation, err := p.relationRepo.GetRelationBetween(ctx, updaterID, comment.Creator.ID().String())
			if err != nil {
				return nil, err
			}

			if !relation.IsFriend() {
				return nil, ErrNotFriend
			}
		}

		// get the updater info
		updater, err := p.userRepo.GetByID(ctx, updaterID)
		if err != nil {
			return nil, err
		}

		// update the comment
		err = comment.UpdateContent(updater, content)
		if err != nil {
			return nil, err
		}

		return comment, nil
	})
}
