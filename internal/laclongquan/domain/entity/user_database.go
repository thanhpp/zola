package entity

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewUserFromDB(
	userUUID, name, phone, hashpass, state, role,
	link, avatar, coverImage,
	username, description,
	address, city, country string) (*User, error) {
	userID, err := uuid.Parse(userUUID)
	if err != nil {
		return nil, errors.WithMessage(err, "parse uuid")
	}

	return &User{
		id:   userID,
		name: name,
		account: Account{
			Phone:    phone,
			HashPass: hashpass,
		},
		state:       UserState(state),
		role:        UserRole(role),
		Link:        link,
		Avatar:      avatar,
		CoverImg:    coverImage,
		Username:    username,
		Description: description,
		Address: UserAddress{
			Address: address,
			City:    city,
			Country: country,
		},
	}, nil
}
