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

func (h UserHandler) SignUp(ctx context.Context, phone, pass, name, avatar string) error {
	newUser, err := h.fac.NewUser(phone, pass, name, avatar)
	if err != nil {
		return err
	}

	if err := h.repo.Create(ctx, newUser); err != nil {
		return err
	}

	return nil
}
