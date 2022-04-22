package services

import (
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/repository"
	"landing_admin_backend/internal/services/auth"
	"landing_admin_backend/internal/services/posters"
	"landing_admin_backend/internal/services/users"
	"landing_admin_backend/pkg/token_manager"
)

type Services struct {
	Auth    auth.Service
	Users   users.Service
	Posters posters.Service
}

func Setup(cfg *config.Config, repository *repository.Repository) *Services {
	tokenManager := token_manager.NewTokenManager(cfg.TokenSecret, cfg.TokenTTL, cfg.RefreshTokenTTL)
	return &Services{
		Auth:    auth.NewService(repository, tokenManager),
		Users:   users.NewService(repository),
		Posters: posters.NewService(repository),
	}
}
