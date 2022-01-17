package application

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/pkg/logger"
)

func (u UserHandler) AdminCreateUser(ctx context.Context, requestorID, phone, pass, name, username, des, address, city, country string) (*entity.User, error) {
	requestor, err := u.repo.GetByID(ctx, requestorID)
	if err != nil {
		return nil, err
	}

	if !requestor.IsAdmin() {
		return nil, entity.ErrPermissionDenied
	}

	newUser, err := u.fac.NewUser(phone, pass, name, "")
	if err != nil {
		return nil, err
	}

	newAdr, err := u.fac.NewAddress(address, city, country)
	if err != nil {
		return nil, err
	}

	newUser.UpdateAddress(newAdr)

	if len(username) != 0 {
		if err := newUser.UpdateUsername(username); err != nil {
			return nil, err
		}
	}

	if err := newUser.UpdateDescription(des); err != nil {
		return nil, err
	}

	if err := u.repo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	go func() {
		if err := u.esClient.CreateOrUpdateUser(newUser); err != nil {
			logger.Errorf("add user %s to ES failed: %v", newUser.ID(), err)
			return
		}
		logger.Infof("add user %s to ES success", newUser.ID())
	}()

	return newUser, nil
}
