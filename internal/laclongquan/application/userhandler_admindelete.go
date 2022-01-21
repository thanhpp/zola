package application

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/pkg/logger"
)

var (
	ErrSelfDelete = errors.New("self delete")
)

func (u UserHandler) AdminDeleteUser(ctx context.Context, requestorID, requestedID string) error {
	requestor, err := u.repo.GetByID(ctx, requestorID)
	if err != nil {
		return err
	}

	if !requestor.IsAdmin() {
		return entity.ErrPermissionDenied
	}

	requested, err := u.repo.GetByID(ctx, requestedID)
	if err != nil {
		return err
	}

	if requestor.Equal(requested) {
		return ErrSelfDelete
	}

	if err := u.repo.DeleteByIDCascade(ctx, requestedID); err != nil {
		return err
	}

	go func() {
		if err := u.esClient.DeleteUser(requestedID); err != nil {
			logger.Errorf("delete user %s from ES failed: %v", requestedID, err)
			return
		}
		logger.Infof("delete user %s from ES success", requestedID)
	}()

	return nil
}
