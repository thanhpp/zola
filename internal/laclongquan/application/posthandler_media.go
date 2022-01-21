package application

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

func (p PostHandler) GetMedia(ctx context.Context, userID, postID, mediaID string) (*entity.Media, error) {
	// user, err := p.userRepo.GetByID(ctx, userID)
	// if err != nil {
	// 	return nil, err
	// }

	post, err := p.repo.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	// var relation *entity.Relation
	// if userID != post.CreatorUUID().String() {
	// 	relation, err = p.relationRepo.GetRelationBetween(ctx, userID, post.CreatorUUID().String())
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return post.FindMediaByID(mediaID)
}
