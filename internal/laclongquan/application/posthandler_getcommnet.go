package application

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

type GetPostCommentRes struct {
	IsBlocked bool
	Comment   *entity.Comment
}

func (p PostHandler) GetPostComments(ctx context.Context, requestorID, postID string, offset, limit int) ([]*GetPostCommentRes, error) {
	requestor, err := p.userRepo.GetByID(ctx, requestorID)
	if err != nil {
		return nil, err
	}

	if requestor.IsLocked() {
		return nil, entity.ErrLockedUser
	}

	// check if the requestor is able to get the post
	_, err = p.GetPost(ctx, requestorID, postID)
	if err != nil {
		return nil, err
	}

	comments, err := p.commentRepo.GetByPostIDFromNonBlockedActiveUser(
		ctx,
		requestorID,
		postID,
		offset,
		limit)
	if err != nil {
		return nil, err
	}

	var (
		res = make([]*GetPostCommentRes, 0, len(comments))
	)

	for i := range comments {
		if requestor.Equal(comments[i].Creator) {
			res = append(res, &GetPostCommentRes{
				Comment:   comments[i],
				IsBlocked: false,
			})
			continue
		}
		tmpRelation, err := p.relationRepo.GetRelationBetween(ctx, requestorID, comments[i].CreatorUUID().String())
		if err != nil && !errors.Is(err, repository.ErrRelationNotFound) {
			return nil, err
		}
		if tmpRelation != nil && tmpRelation.IsBlock() {
			res = append(res, &GetPostCommentRes{
				IsBlocked: true,
			})
		}
	}

	return res, nil
}
