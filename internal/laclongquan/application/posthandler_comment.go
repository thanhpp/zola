package application

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

func (p PostHandler) CreateComment(ctx context.Context, postID, creatorID, content string) error {
	post, err := p.repo.GetByID(ctx, postID)
	if err != nil {
		return err
	}

	postCreator, err := p.userRepo.GetByID(ctx, post.Creator())
	if err != nil {
		return err
	}

	creator, err := p.userRepo.GetByID(ctx, creatorID)
	if err != nil {
		return err
	}

	var relation *entity.Relation
	if !postCreator.Equal(creator) {
		relation, err = p.relationRepo.GetRelationBetween(ctx, postCreator.ID().String(), creator.ID().String())
		if err != nil {
			return err
		}
	}

	if err := post.CanCreateComment(postCreator, creator, relation); err != nil {
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
				return nil, entity.ErrNotFriend
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

func (p PostHandler) DeleteComment(ctx context.Context, deleterID, postID, commentID string) error {
	deleter, err := p.userRepo.GetByID(ctx, deleterID)
	if err != nil {
		return err
	}

	comment, err := p.commentRepo.GetByIDAndPostID(ctx, commentID, postID)
	if err != nil {
		return err
	}

	if err := comment.IsDeletable(deleter); err != nil {
		return err
	}

	return p.commentRepo.Delete(ctx, comment)
}
