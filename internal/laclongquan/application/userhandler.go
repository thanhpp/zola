package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
)

type UserHandler struct {
	fac          entity.UserFactory
	repo         repository.UserRepository
	blockRepo    repository.BlockRepository
	relationRepo repository.RelationRepository
}

func NewUserHandler(
	fac entity.UserFactory,
	repo repository.UserRepository,
	blockRepo repository.BlockRepository,
	relationRepo repository.RelationRepository,
) UserHandler {
	return UserHandler{
		fac:          fac,
		repo:         repo,
		blockRepo:    blockRepo,
		relationRepo: relationRepo,
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

func (u UserHandler) BlockUser(ctx context.Context, blocker uuid.UUID, blocked string) error {
	blockedUser, err := u.repo.GetByID(ctx, blocked)
	if err != nil {
		return err
	}

	blockerUser, err := u.repo.GetByID(ctx, blocker.String())
	if err != nil {
		return err
	}

	block, err := u.fac.NewBlock(blockerUser, blockedUser)
	if err != nil {
		return err
	}

	return u.blockRepo.Create(ctx, block)
}

func (u UserHandler) UnblockUser(ctx context.Context, blocker uuid.UUID, blocked string) error {
	blockedUser, err := u.repo.GetByID(ctx, blocked)
	if err != nil {
		return err
	}

	return u.blockRepo.Delete(ctx, blocker.String(), blockedUser.ID().String())
}
