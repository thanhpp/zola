package entity

import (
	"time"
	"unicode"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrInputTooLong       = errors.New("input too long")
	ErrEmptyInput         = errors.New("empty input")
	ErrInvalidUsername    = errors.New("invalid username")
	ErrInvalidInputLength = errors.New("invalid input length")
)

type User struct {
	id          uuid.UUID
	Username    string
	Description string
	name        string
	Link        string
	state       UserState
	account     Account
	role        UserRole
	Address     UserAddress
	Avatar      string
	CoverImg    string
	CreatedAt   time.Time
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) GetLink() string {
	return u.Link
}

func (u *User) UpdateLink(link string) {
	if u == nil {
		return
	}

	u.Link = link
}

func (u User) Account() Account {
	return u.account
}

func (u User) GetUsername() string {
	return u.Username
}

func (u *User) UpdateUsername(username string) error {
	if u == nil {
		return nil
	}

	if err := validateUsername(username); err != nil {
		return err
	}

	u.Username = username

	return nil
}

func validateUsername(username string) error {
	if !stringLengthCheck(username, 1, 50) {
		return ErrInvalidUsername
	}

	if !unicode.IsLetter(rune(username[0])) {
		return ErrInvalidUsername
	}

	for _, c := range username {
		if unicode.IsLetter(c) || c == '_' {
			continue
		}
		return ErrInvalidUsername
	}

	return nil
}

func (u User) GetDescription() string {
	return u.Description
}

func (u *User) UpdateDescription(description string) error {
	if u == nil {
		return nil
	}

	if !stringLengthCheck(description, 0, 150) {
		return ErrInvalidInputLength
	}

	u.Description = description

	return nil
}

func (u User) PassEqual(pass string, accCipher AccountCipher) error {
	return u.account.Equal(u.account.Phone, pass, accCipher)
}

func (u User) Role() string {
	return u.role.String()
}

func (u User) IsAdmin() bool {
	return u.role == UserRoleAdmin
}

func (u *User) UpdateAddress(address *UserAddress) {
	if u == nil {
		return
	}

	u.Address = *address
}

func (u *User) UpdatePass(oldPass, newPass string, accCipher AccountCipher) error {
	if u.IsLocked() {
		return ErrLockedUser
	}

	if err := u.account.UpdatePass(oldPass, newPass, accCipher); err != nil {
		return err
	}

	return nil
}

func (u User) State() UserState {
	return u.state
}

func (u User) IsLocked() bool {
	return u.state == UserStateLocked
}

func (u User) GetAvatar() string {
	return u.Avatar
}

func (u *User) UpdateAvatar(avatar string) {
	if u == nil {
		return
	}

	u.Avatar = avatar
}

func (u User) GetCoverImage() string {
	return u.CoverImg
}

func (u *User) UpdateCoverImage(coverImage string) {
	if u == nil {
		return
	}

	u.CoverImg = coverImage
}

func (u User) GetAddress() string {
	return u.Address.Address
}
func (u User) GetCity() string {
	return u.Address.City
}
func (u User) GetCountry() string {
	return u.Address.Country
}

func stringLengthCheck(input string, min, max int) bool {
	return len(input) >= min && len(input) <= max
}

func (u User) CanGetUserInfo(requestor *User, relation *Relation) error {
	if requestor.IsAdmin() {
		return nil
	}

	if u.IsLocked() {
		return ErrLockedUser
	}

	if requestor.ID().String() == u.ID().String() {
		return nil
	}

	if relation != nil && relation.IsFriend() {
		return nil
	}

	return ErrPermissionDenied
}

func (u User) Equal(user *User) bool {
	return user != nil && u.ID().String() == user.ID().String()
}

func (u User) CreatedAtUnix() int64 {
	return u.CreatedAt.Unix()
}

func (u User) CanGetUserRequestedFriend(user *User) error {
	if user == nil {
		return ErrPermissionDenied
	}

	if user.IsAdmin() {
		return nil
	}

	if u.IsLocked() {
		return ErrLockedUser
	}

	if u.Equal(user) {
		return nil
	}

	return ErrPermissionDenied
}

func (u User) CanGetUserFriends(user *User) error {
	if user == nil {
		return ErrNilUser
	}

	if u.Equal(user) || user.IsAdmin() {
		return nil
	}

	return ErrPermissionDenied
}
