package entity

import (
	"errors"
	"regexp"

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
	NewAdmin(phone, pass, name, avatar string) (*User, error)
	NewFriendRequest(requestor, requestee *User) (*Relation, error)
	NewBlockRelation(blocker, blocked *User) (*Relation, error)
}

func NewUserFactory(accountCipher AccountCipher) UserFactory {
	return &userFactoryImpl{
		hashCost:      hashCost,
		phoneRegexp:   regexp.MustCompile(regexpPhoneNumber),
		accountCipher: accountCipher,
	}
}

type userFactoryImpl struct {
	hashCost      int
	phoneRegexp   *regexp.Regexp
	accountCipher AccountCipher
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
		state:   UserStateActive,
		role:    UserRoleUser,
	}, nil
}

func (fac userFactoryImpl) NewAdmin(phone, pass, name, avatar string) (*User, error) {
	newUser, err := fac.NewUser(phone, pass, name, avatar)
	if err != nil {
		return nil, err
	}
	newUser.role = UserRoleAdmin

	return newUser, nil
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

func (fac userFactoryImpl) newAccount(phone, pass string) (*Account, error) {
	if err := fac.validatePhone(phone); err != nil {
		return nil, err
	}

	if err := validatePass(pass, phone); err != nil {
		return nil, err
	}

	hashPass, err := fac.accountCipher.Encrypt(pass)
	if err != nil {
		return nil, err
	}

	return &Account{
		Phone:    phone,
		HashPass: hashPass,
	}, nil
}
