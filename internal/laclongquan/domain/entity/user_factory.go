package entity

import (
	"errors"
	"regexp"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	regexpPhoneNumber string = "^[\\+]?[(]?[0-9]{3}[)]?[-\\s\\.]?[0-9]{3}[-\\s\\.]?[0-9]{4,6}$"
	hashCost          int    = 10
)

var (
	ErrInvalidPhone    = errors.New("invalid phone number")
	ErrInvalidPassword = errors.New("invalid password")
)

type UserFactory interface {
	NewUser(phone, pass, name, avatar string) (*User, error)
	NewBlock(blocker, blocked *User) (*Block, error)
}

func NewUserFactory() UserFactory {
	return &userFactoryImpl{
		hashCost:    hashCost,
		phoneRegexp: regexp.MustCompile(regexpPhoneNumber),
	}
}

type userFactoryImpl struct {
	hashCost    int
	phoneRegexp *regexp.Regexp
}

func (fac userFactoryImpl) NewUser(phone, pass, name, avatar string) (*User, error) {
	userID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	account, err := fac.newAccount(phone, pass)
	if err != nil {
		return nil, err
	}

	return &User{
		id:      userID,
		name:    name,
		avatar:  avatar,
		account: *account,
	}, nil
}

// ---------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- ACCOUNT ----------------------------------------------------------

func (fac userFactoryImpl) hashString(in string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(in), fac.hashCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (fac userFactoryImpl) validatePhone(phone string) error {
	if !fac.phoneRegexp.Match([]byte(phone)) {
		return ErrInvalidPhone
	}

	return nil
}

func (fac userFactoryImpl) validatePass(pass string, phone string) error {
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

func (fac userFactoryImpl) newAccount(phone, pass string) (*Account, error) {
	if err := fac.validatePhone(phone); err != nil {
		return nil, err
	}

	if err := fac.validatePass(pass, phone); err != nil {
		return nil, err
	}

	hashPass, err := fac.hashString(pass)
	if err != nil {
		return nil, err
	}

	return &Account{
		Phone:    phone,
		HashPass: hashPass,
	}, nil
}
