package application

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

type UserHandler struct {
	fac  entity.UserFactory
	repo repository.UserRepository
}

func NewUserHandler(fac entity.UserFactory, repo repository.UserRepository) UserHandler {
	return UserHandler{
		fac:  fac,
		repo: repo,
	}
}

func (u UserHandler) CreateUser(ctx context.Context, phone, pass, name, avatar string) error {
	newUser, err := u.fac.NewUser(phone, pass, name, avatar)
	if err != nil {
		return err
	}

	if err := u.repo.Create(ctx, newUser); err != nil {
		return err
	}

	return nil
}

func (u UserHandler) GetUser(ctx context.Context, phone, pass string) (*entity.User, error) {
	user, err := u.repo.GetByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}

	if err := user.PassEqual(pass); err != nil {
		return nil, err
	}

	return user, nil
}
