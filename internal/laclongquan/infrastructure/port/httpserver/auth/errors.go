package auth

import "errors"

var (
	ErrNilConfig            = errors.New("nil config")
	ErrInvalidMethod        = errors.New("invalid method")
	ErrInvalidToken         = errors.New("invalid token")
	ErrClaimsIsNotMapClaims = errors.New("not map claims")
	ErrTokenNotFound        = errors.New("token not found")
)
