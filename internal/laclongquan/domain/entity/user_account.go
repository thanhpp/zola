package entity

import (
	"errors"
	"unicode"

	"github.com/thanhpp/zola/pkg/alg"
)

var (
	ErrPhoneNotEqual = errors.New("phone not equal")
	ErrPassNotEqual  = errors.New("pass not equal")
	ErrSameOldPass   = errors.New("same old pass")
	ErrCommonPass    = errors.New("common pass")
)

type AccountCipher interface {
	Encrypt(raw string) (string, error)
	Decrypt(encrypted string) (string, error)
}

type Account struct {
	Phone    string
	HashPass string
}

func (acc Account) Equal(phone, pass string, accCipher AccountCipher) error {
	if acc.Phone != phone {
		return ErrPhoneNotEqual
	}

	if err := acc.passEqual(pass, accCipher); err != nil {
		return ErrPassNotEqual
	}

	return nil
}

func (acc *Account) passEqual(pass string, accCipher AccountCipher) error {
	plainPass, err := accCipher.Decrypt(acc.HashPass)
	if err != nil {
		return err
	}

	if plainPass != pass {
		return ErrPassNotEqual
	}

	return nil
}

func (acc *Account) UpdatePass(oldRawPass, newRawPass string, accCipher AccountCipher) error {
	if err := acc.passEqual(oldRawPass, accCipher); err != nil {
		return err
	}

	plainPass, err := accCipher.Decrypt(acc.HashPass)
	if err != nil {
		return err
	}

	commonLength := alg.LongestCommonSubstring(plainPass, newRawPass)
	if commonLength > int(float64(len(newRawPass))*0.8) {
		return ErrCommonPass
	}

	if err := validatePass(newRawPass, acc.Phone); err != nil {
		return ErrInvalidPassword
	}

	acc.HashPass, err = accCipher.Encrypt(newRawPass)
	if err != nil {
		return err
	}

	return nil
}

func (acc *Account) AdminUpdatePass(newRawPass string, accCipher AccountCipher) error {
	if acc == nil {
		return nil
	}

	err := validatePass(newRawPass, acc.Phone)
	if err != nil {
		return ErrInvalidPassword
	}

	acc.HashPass, err = accCipher.Encrypt(newRawPass)
	if err != nil {
		return err
	}

	return nil
}

func validatePass(pass string, phone string) error {
	if len(pass) > 10 || len(pass) < 6 {
		return ErrInvalidPassword
	}

	if pass == phone {
		return ErrInvalidPassword
	}

	for _, c := range pass {
		switch {
		case unicode.IsLower(c):
			continue

		case unicode.IsUpper(c):
			continue

		case unicode.IsNumber(c):
			continue
		}

		return ErrInvalidPassword
	}

	return nil
}
