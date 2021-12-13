package auth

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type Claims struct {
	jwt.StandardClaims
	User JWTUserData `json:"user"`
}

type JWTUserData struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

func (c Claims) exportMapClaims() (*jwt.MapClaims, error) {
	var (
		mapClaims = new(jwt.MapClaims)
	)

	data, err := json.Marshal(c)
	if err != nil {
		return nil, errors.WithMessage(err, "marshal json")
	}

	err = json.Unmarshal(data, mapClaims)
	if err != nil {
		return nil, errors.WithMessage(err, "unmarshal json")
	}

	return mapClaims, nil
}
