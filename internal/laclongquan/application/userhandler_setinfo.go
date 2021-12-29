package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/domain/repository"
	"github.com/thanhpp/zola/pkg/logger"
)

var (
	ErrCanNotUseMedia = errors.New("can not use media")
)

type SetUserInfoResult struct {
	AvatarLink     string
	CoverImageLink string
	Link           string
}

func (u UserHandler) SetUserInfo(
	ctx context.Context,
	userID uuid.UUID,
	username, description, address, city, country, link string,
	avatar, coverImage *entity.Media,
) error {
	return u.repo.Update(ctx, userID.String(), func(ctx context.Context, user *entity.User) (*entity.User, error) {
		if user == nil {
			return nil, repository.ErrUserNotFound
		}

		if user.IsLocked() {
			return nil, entity.ErrLockedUser
		}

		if err := user.UpdateUsername(username); err != nil {
			return nil, err
		}

		if err := user.UpdateDescription(description); err != nil {
			return nil, err
		}

		adr, err := u.fac.NewAddress(address, city, country)
		if err != nil {
			return nil, err
		}
		user.UpdateAddress(adr)

		user.UpdateLink(link)

		if avatar != nil {
			if !avatar.IsOwner(user) {
				return nil, ErrCanNotUseMedia
			}
			user.UpdateAvatar(avatar.ID())
			logger.Debugf("set user info - update avatar: %s", avatar.ID())
		}

		if coverImage != nil {
			if !coverImage.IsOwner(user) {
				return nil, ErrCanNotUseMedia
			}
			user.UpdateCoverImage(coverImage.ID())
		}

		return user, nil
	})
}
