package application

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

type GetPostResult struct {
	Post         *entity.Post
	Author       *entity.User
	LikeCount    int
	CommentCount int
	IsLiked      bool
	CanEdit      bool
	CanComment   bool
}

func (p PostHandler) GetPost(ctx context.Context, userID, postID string) (*GetPostResult, error) {
	user, err := p.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	post, err := p.repo.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	// get the relation between the user and the post creator
	var relation *entity.Relation
	if userID != post.Creator() {
		relation, err = p.relationRepo.GetRelationBetween(ctx, userID, post.Creator())
		if err != nil {
			return nil, err
		}

		if relation.IsBlock() {
			return nil, ErrAlreadyBlocked
		}
	}

	// authorization
	if err := post.CanUserGetPost(user, relation); err != nil {
		return nil, err
	}

	// get the author of the post
	author, err := p.userRepo.GetByID(ctx, post.Creator())
	if err != nil {
		return nil, err
	}

	// like count
	likeCount, err := p.likeRepo.Count(ctx, postID)
	if err != nil {
		return nil, err
	}

	// comment count
	commentCount, err := p.commentRepo.CountByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}

	// form response
	var (
		canEdit bool
		// FIXME: canComment bool
	)
	if err := post.CanUserEditPost(user); err == nil {
		canEdit = true
	}

	return &GetPostResult{
		Post:         post,
		Author:       author,
		CommentCount: commentCount,
		LikeCount:    likeCount,
		IsLiked:      p.likeRepo.IsLiked(ctx, userID, postID),
		CanEdit:      canEdit,
	}, nil
}
