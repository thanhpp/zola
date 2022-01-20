package application

import (
	"context"
	"time"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

func (p PostHandler) AdminGetList(ctx context.Context, requestorID, lastPostID string, limit, offset int) (*GetListPostRes, error) {
	requestor, err := p.userRepo.GetByID(ctx, requestorID)
	if err != nil {
		return nil, err
	}

	if !requestor.IsAdmin() {
		return nil, entity.ErrPermissionDenied
	}

	var (
		timeMileStone time.Time
	)
	if len(lastPostID) != 0 {
		lastPost, err := p.repo.GetByID(ctx, lastPostID)
		if err != nil {
			return nil, err
		}

		timeMileStone = lastPost.CreatedAtTime()
	}

	posts, newItems, err := p.repo.GetListPostForAdmin(ctx, requestorID, timeMileStone, offset, limit)
	if err != nil {
		return nil, err
	}

	var (
		resPosts     = make([]*GetListPostResElem, 0, len(posts))
		creatorCache = make(map[string]*entity.User, len(posts))
	)
	creatorCache[requestor.ID().String()] = requestor
	for i := range posts {
		var resPostElem = new(GetListPostResElem)
		resPostElem.Post = posts[i]

		// like count
		likeCount, err := p.likeRepo.Count(ctx, posts[i].ID())
		if err != nil {
			return nil, err
		}
		resPostElem.LikeCount = likeCount

		// comment count
		commentCount, err := p.commentRepo.CountByPostID(ctx, posts[i].ID())
		if err != nil {
			return nil, err
		}
		resPostElem.CommentCount = commentCount

		// is liked
		resPostElem.IsLiked = p.likeRepo.IsLiked(ctx, requestorID, posts[i].ID())

		// can edit
		resPostElem.CanEdit = posts[i].CanUserEditPost(requestor) == nil

		// add user
		user, ok := creatorCache[posts[i].Creator()]
		if ok {
			resPostElem.Creator = user
			resPosts = append(resPosts, resPostElem)
			continue
		}
		user, err = p.userRepo.GetByID(ctx, posts[i].Creator())
		if err != nil {
			return nil, err
		}
		resPostElem.Creator = user
		resPosts = append(resPosts, resPostElem)
		creatorCache[posts[i].Creator()] = user
	}

	return &GetListPostRes{
		Posts:    resPosts,
		NewItems: newItems,
	}, nil
}
