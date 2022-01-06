package application

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

var (
	ErrNotABlockRelation = errors.New("not a block relation")
	ErrAlreadyBlocked    = errors.New("already blocked")
)

type UserHandler struct {
	fac          entity.UserFactory
	repo         repository.UserRepository
	relationRepo repository.RelationRepository
	postRepo     repository.PostRepository
	accCipher    entity.AccountCipher
}

func NewUserHandler(
	fac entity.UserFactory,
	repo repository.UserRepository,
	relationRepo repository.RelationRepository,
	postRepo repository.PostRepository,
	accountCipher entity.AccountCipher,
) UserHandler {
	return UserHandler{
		fac:          fac,
		repo:         repo,
		relationRepo: relationRepo,
		postRepo:     postRepo,
		accCipher:    accountCipher,
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

func (u UserHandler) CreateAdminUser(ctx context.Context, phone, pass, name, avatar string) error {
	newUser, err := u.fac.NewAdmin(phone, pass, name, avatar)
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

	if err := user.PassEqual(pass, u.accCipher); err != nil {
		return nil, err
	}

	if user.IsLocked() {
		return nil, entity.ErrLockedUser
	}

	return user, nil
}
