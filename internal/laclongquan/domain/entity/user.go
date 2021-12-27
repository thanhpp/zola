package entity

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrInputTooLong = errors.New("input too long")
)

type User struct {
	id          uuid.UUID
	Username    string
	Description string
	name        string
	avatar      string
	Link        string
	state       UserState
	account     Account
	role        UserRole
	Address     UserAddress
	Avatar      string
	CoverImg    string
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

	if len(username) > 500 {
		return ErrInputTooLong
	}

	u.Username = username

	return nil
}

func (u User) GetDescription() string {
	return u.Description
}

func (u *User) UpdateDescription(description string) error {
	if u == nil {
		return nil
	}

	if len(description) > 500 {
		return ErrInputTooLong
	}

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
