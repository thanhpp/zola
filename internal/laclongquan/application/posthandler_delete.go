package application

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/pkg/logger"
)

var (
	ErrPostCannotBeDeleted = errors.New("post cannot be deleted")
)

func (p PostHandler) DeletePost(ctx context.Context, userID, postID string) error {
	post, err := p.repo.GetByID(ctx, postID)
	if err != nil {
		return err
	}

	user, err := p.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if !post.CanBeDeletedBy(user) {
		return ErrPostCannotBeDeleted
	}

	if err := p.repo.Delete(ctx, postID); err != nil {
		return err
	}

	// clean related media
	for i := range post.Media() {
		p.filehdl.Cleanup(post.Media()[i].Path())
		p.filehdl.Cleanup(post.Media()[i].ThumbPath())
	}

	go func() {
		if err := p.esClient.DeletePost(postID); err != nil {
			logger.Errorf("error deleting %s post from elasticsearch %v", postID, err)
			return
		}
		logger.Infof("deleted %s post from elasticsearch", postID)
	}()

	return nil
}
