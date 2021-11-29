package entity_test

import (
	"errors"
	"testing"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/accountcipher"
)

func TestFactory(t *testing.T) {
	var (
		accCipher, _ = accountcipher.New("12345678901234567890234567890123")
		fac          = entity.NewUserFactory(accCipher)
	)

	t.Run("create invalid user - phone", func(t *testing.T) {
		var (
			phone  = "02"
			pass   = ""
			name   = ""
			avatar = ""
		)

		_, err := fac.NewUser(phone, pass, name, avatar)
		if errors.Is(err, entity.ErrInvalidPhone) {
			return
		}

		t.Error(err)

		return
	})

	t.Run("create invalid user - pass", func(t *testing.T) {
		var (
			phone  = "0123456789"
			pass   = "Thanh.28PP"
			name   = ""
			avatar = ""
		)

		_, err := fac.NewUser(phone, pass, name, avatar)
		if errors.Is(err, entity.ErrInvalidPassword) {
			return
		}

		t.Error(err)

		return
	})

	t.Run("create valid user", func(t *testing.T) {
		var (
			phone  = "0123456789"
			pass   = "Thanh28PP"
			name   = ""
			avatar = ""
		)

		_, err := fac.NewUser(phone, pass, name, avatar)
		if err == nil {
			return
		}

		t.Error(err)

		return
	})
}
