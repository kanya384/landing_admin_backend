package auth

import (
	"context"
	"landing_admin_backend/internal/repository"
	"landing_admin_backend/pkg/token_manager"
)

const (
	TokenCookieKey   = "token"
	RefreshCookieKey = "refresh_token"
)

type Service interface {
	Authenticate(ctx context.Context, login string, pass string) (tokens token_manager.AuthTokens, err error)
}

type service struct {
	repository   *repository.Repository
	tokenManager token_manager.TokenManager
}

func NewService(repository *repository.Repository, tokenManager token_manager.TokenManager) Service {
	return &service{
		repository:   repository,
		tokenManager: tokenManager,
	}
}

func (s *service) Authenticate(ctx context.Context, login string, pass string) (tokens token_manager.AuthTokens, err error) {
	user, err := s.repository.Users.GetUserByCredetinals(ctx, login, pass)
	if err != nil {
		return
	}
	tokens, err = s.tokenManager.GenerateAuthTokens(user.ID.Hex(), user.Name, user.Role)
	if err != nil {
		return
	}
	err = setTokensToCookies(tokens)
	return
}

func setTokensToCookies(tokens token_manager.AuthTokens) (err error) {
	return
}
