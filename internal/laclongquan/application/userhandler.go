package application

import (
	"context"
	"errors"

	"github.com/google/uuid"
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
}

func NewUserHandler(
	fac entity.UserFactory,
	repo repository.UserRepository,
	relationRepo repository.RelationRepository,
) UserHandler {
	return UserHandler{
		fac:          fac,
		repo:         repo,
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

func (u UserHandler) BlockUser(ctx context.Context, blockerID uuid.UUID, blockedID string) error {
	relation, err := u.relationRepo.GetRelationBetween(ctx, blockerID.String(), blockedID)
	if err != nil {
		if !errors.Is(err, repository.ErrRelationNotFound) {
			return err
		}
	}

	// if a relation exists, update it to block
	if relation != nil {
		if relation.IsBlock() {
			return ErrAlreadyBlocked
		}
		relation.Block()
		return u.relationRepo.UpdateRelation(ctx, relation)
	}

	// create a new block relation
	blocker, err := u.repo.GetByID(ctx, blockerID.String())
	if err != nil {
		return err
	}
	blocked, err := u.repo.GetByID(ctx, blockedID)
	if err != nil {
		return err
	}
	newBlockRelation, err := u.fac.NewBlockRelation(blocker, blocked)
	if err != nil {
		return err
	}

	return u.relationRepo.CreateRelation(ctx, newBlockRelation)
}

func (u UserHandler) UnblockUser(ctx context.Context, blockerID uuid.UUID, blockedID string) error {
	relation, err := u.relationRepo.GetRelationBetween(ctx, blockerID.String(), blockedID)
	if err != nil {
		return err
	}

	if !relation.IsBlock() {
		return ErrNotABlockRelation
	}

	return u.relationRepo.DeleteRelation(ctx, relation)
}
