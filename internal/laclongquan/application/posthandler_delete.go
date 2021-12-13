package application

import (
	"context"
	"errors"
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
	}

	return nil
}
