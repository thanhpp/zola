package auth

import (
	"context"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
)

type AuthService struct {
	fac  *authFactory
	repo Repository
}

func NewAuthService(cfg *Config, repo Repository) (*AuthService, error) {
	if cfg == nil {
		return nil, ErrNilConfig
	}

	fac, err := newFactoryFromConfig(cfg)
	if err != nil {
		return nil, err
	}

	return &AuthService{
		fac:  fac,
		repo: repo,
	}, nil
}

// NewTokenFromUser creates a new token from user and delete other's user token
func (s AuthService) NewTokenFromUser(ctx context.Context, user *entity.User) (string, error) {
	claims, err := s.fac.NewClaimsFromUser(user)
	if err != nil {
		return "", nil
	}

	token, err := s.fac.SignClaims(claims)
	if err != nil {
		return "", nil
	}

	if err := s.repo.DeleteByUserID(ctx, user.ID().String()); err != nil {
		return "", err
	}

	if err := s.repo.Cache(ctx, claims); err != nil {
		return "", err
	}

	return token, nil
}

func (s AuthService) NewClaimsFromToken(ctx context.Context, token string) (*Claims, error) {
	claims, err := s.fac.NewClaimsFromToken(token)
	if err != nil {
		return nil, err
	}

	err = s.repo.CheckByID(ctx, claims.Id)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
