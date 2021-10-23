package entity

type User struct {
	id      string
	name    string
	avatar  string
	account Account
}

func (u User) ID() string {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) Avatar() string {
	return u.avatar
}

func (u User) Account() Account {
	return u.account
}

func (u User) PassEqual(pass string) error {
	return u.account.Equal(u.account.Phone, pass)
}
