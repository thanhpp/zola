package application

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

func (u UserHandler) ChangePassword(ctx context.Context, userID string, oldPass, newPass string) error {
	return u.repo.Update(ctx, userID, func(ctx context.Context, user *entity.User) (*entity.User, error) {
		if err := user.UpdatePass(oldPass, newPass, u.accCipher); err != nil {
			return user, err
		}

		return user, nil
	})
}
