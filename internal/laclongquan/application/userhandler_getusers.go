package application

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

var (
	ErrPermissionDenied = errors.New("permission denied")
)

type GetUserRes struct {
	Total    int
	UserList []*entity.User
}

func (u UserHandler) GetUsers(ctx context.Context, userID string) (*GetUserRes, error) {
	user, err := u.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if !user.IsAdmin() {
		return nil, ErrPermissionDenied
	}

	users, total, err := u.repo.GetAllUsers(ctx, -1, -1, "", "", "", "")
	if err != nil {
		return nil, err
	}

	return &GetUserRes{
		Total:    total,
		UserList: users,
	}, nil
}
