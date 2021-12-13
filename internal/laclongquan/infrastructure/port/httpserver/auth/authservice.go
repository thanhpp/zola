package auth

import (
	"context"
	"time"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/pkg/booting"
	"github.com/thanhpp/zola/pkg/logger"
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

func (s AuthService) DeleteUserTokens(ctx context.Context, userID string) error {
	return s.repo.DeleteByUserID(ctx, userID)
}

func (s AuthService) DeleteExpiredDaemons() booting.Daemon {
	return func(ctx context.Context) (start func() error, cleanup func()) {
		ticker := time.NewTicker(time.Hour)
		start = func() error {
			if err := s.repo.DeleteExpired(ctx); err != nil {
				logger.Errorf("Delete expired token error: %v", err)
			}

			for {
				select {
				case <-ticker.C:
					if err := s.repo.DeleteExpired(ctx); err != nil {
						logger.Errorf("Delete expired token error: %v", err)
					}
					logger.Info("Delete expired token")

				case <-ctx.Done():
					return nil
				}
			}
		}

		cleanup = func() {
			ticker.Stop()
		}

		return start, cleanup
	}
}
