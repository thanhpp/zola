package application

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

type PermReq struct {
	CanComment bool
}

func (p PostHandler) EditPerm(ctx context.Context, userID, postID string, req PermReq) error {
	return p.repo.Update(ctx, postID, func(ctx context.Context, post *entity.Post) (*entity.Post, error) {
		user, err := p.userRepo.GetByID(ctx, userID)
		if err != nil {
			return nil, err
		}

		if err := post.UpdateCanComment(user, req.CanComment); err != nil {
			return nil, err
		}

		return post, nil
	})
}
