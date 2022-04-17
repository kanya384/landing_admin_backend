package auth

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"
)

type Service interface {
	Authenticate(ctx context.Context, login string, pass string) (tokens domain.AuthTokens, err error)
}

type service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Authenticate(ctx context.Context, login string, pass string) (tokens domain.AuthTokens, err error) {
	user, err := s.repository.Users.GetUserByCredetinals(ctx, login, pass)
	if err != nil {
		return
	}

}
