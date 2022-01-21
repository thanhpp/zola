package application

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/pkg/logger"
)

func (u UserHandler) AdminUpdateState(ctx context.Context, requestorID, requestedID, state string) error {
	requestor, err := u.repo.GetByID(ctx, requestorID)
	if err != nil {
		return err
	}

	if !requestor.IsAdmin() {
		return entity.ErrPermissionDenied
	}

	err = u.repo.Update(ctx, requestedID, func(ctx context.Context, user *entity.User) (*entity.User, error) {
		if err := user.SetState(state); err != nil {
			return nil, err
		}
		return user, nil
	})
	if err != nil {
		return err
	}

	requested, err := u.repo.GetByID(ctx, requestedID)
	if err != nil {
		return err
	}

	switch state {
	case entity.UserStateActive.String():
		go func() {
			if err := u.esClient.CreateOrUpdateUser(requested); err != nil {
				logger.Errorf("ES failed to create or update user: %v", err)
				return
			}
			logger.Infof("ES successfully create or update user: %s", requested.ID().String())
		}()

	case entity.UserStateLocked.String():
		go func() {
			if err := u.esClient.DeleteUser(requested.ID().String()); err != nil {
				logger.Errorf("ES failed to delete user: %v", err)
				return
			}
			logger.Infof("ES successfully delete user: %v", requested.ID().String())
		}()
	}

	return nil
}

func (u UserHandler) AdminUpdatePass(ctx context.Context, requestorID, requestedID, newPass string) error {
	requestor, err := u.repo.GetByID(ctx, requestorID)
	if err != nil {
		return err
	}

	if !requestor.IsAdmin() {
		return entity.ErrPermissionDenied
	}

	err = u.repo.Update(ctx, requestedID, func(ctx context.Context, user *entity.User) (*entity.User, error) {
		if err := user.AdminUpdatePass(newPass, u.accCipher); err != nil {
			return nil, err
		}
		return user, nil
	})
	if err != nil {
		return err
	}

	return nil
}
