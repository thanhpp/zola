package application

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/pkg/logger"
)

type LikeHandler struct {
	fac      entity.LikeFactory
	likeRepo repository.LikeRepository
	postRepo repository.PostRepository
}

func NewLikeHandler(fac entity.LikeFactory, likeRepo repository.LikeRepository, postRepo repository.PostRepository) LikeHandler {
	return LikeHandler{
		fac:      fac,
		likeRepo: likeRepo,
		postRepo: postRepo,
	}
}

func (l LikeHandler) LikePost(ctx context.Context, postID, userID string) (int, error) {
	// post id check
	logger.Debugf("post repo nil check %v", l.postRepo)
	post, err := l.postRepo.GetByID(ctx, postID)
	if err != nil {
		return 0, err
	}

	like, err := l.fac.NewLike(post, userID)
	if err != nil {
		return 0, err
	}

	if err := l.likeRepo.CreateOrDelete(ctx, like); err != nil {
		return 0, err
	}

	return l.likeRepo.Count(ctx, postID)
}
