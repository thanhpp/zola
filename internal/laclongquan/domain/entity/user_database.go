package entity

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewUserFromDB(userUUID, name, avatar, phone, hashpass, state string) (*User, error) {
	userID, err := uuid.Parse(userUUID)
	if err != nil {
		return nil, errors.WithMessage(err, "parse uuid")
	}

	return &User{
		id:     userID,
		name:   name,
		avatar: avatar,
		account: Account{
			Phone:    phone,
			HashPass: hashpass,
		},
		state: UserState(state),
	}, nil
}
