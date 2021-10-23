package entity

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPhoneNotEqual = errors.New("phone not equal")
	ErrPassNotEqual  = errors.New("pass not equal")
)

type Account struct {
	Phone    string
	HashPass string
}

func (acc Account) Equal(phone, pass string) error {
	if acc.Phone != phone {
		return ErrPhoneNotEqual
	}

	if err := bcrypt.CompareHashAndPassword([]byte(acc.HashPass), []byte(pass)); err != nil {
		return ErrPassNotEqual
	}

	return nil
}
