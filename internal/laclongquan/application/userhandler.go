package application

import (
	"context"
	"errors"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/adapter/esclient"
	"github.com/thanhpp/zola/pkg/logger"
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
	esClient     *esclient.EsClient
}

func NewUserHandler(
	fac entity.UserFactory,
	repo repository.UserRepository,
	relationRepo repository.RelationRepository,
	postRepo repository.PostRepository,
	accountCipher entity.AccountCipher,
	esClient *esclient.EsClient,
) UserHandler {
	return UserHandler{
		fac:          fac,
		repo:         repo,
		relationRepo: relationRepo,
		postRepo:     postRepo,
		accCipher:    accountCipher,
		esClient:     esClient,
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

	go func() {
		if err := u.esClient.CreateOrUpdateUser(newUser); err != nil {
			logger.Errorf("add user %s to ES failed: %v", newUser.ID(), err)
			return
		}
		logger.Infof("add user %s to ES success", newUser.ID())
	}()

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

	go func() {
		if err := u.esClient.CreateOrUpdateUser(newUser); err != nil {
			logger.Errorf("add user %s to ES failed: %v", newUser.ID(), err)
			return
		}
		logger.Infof("add user %s to ES success", newUser.ID())
	}()

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

func (u UserHandler) SyncAllUser() error {
	// get all user from es
	usersIDs, err := u.esClient.SearchUser("", 0, 10000000)
	if err != nil {
		logger.Errorf("get all user from ES failed: %v", err)
	}

	users, _, err := u.repo.GetAllUsers(context.Background(), -1, -1, "", "", "", "")
	if err != nil {
		return err
	}

	var userIDsMap = make(map[string]struct{})
	for _, user := range users {
		userIDsMap[user.ID().String()] = struct{}{}
		if err := u.esClient.CreateOrUpdateUser(user); err != nil {
			logger.Errorf("add user %s to ES failed: %v", user.ID(), err)
			continue
		}
		logger.Infof("add user %s to ES success", user.ID())
	}

	for i := range usersIDs {
		if _, ok := userIDsMap[usersIDs[i]]; !ok {
			if err := u.esClient.DeleteUser(usersIDs[i]); err != nil {
				logger.Errorf("delete user %s from ES failed: %v", usersIDs[i], err)
				continue
			}
			// delete(userIDsMap, usersIDs[i])
			logger.Infof("delete user %s from ES success", usersIDs[i])
		}
	}

	return nil
}
