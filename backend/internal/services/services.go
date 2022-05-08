package services

import (
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/repository"
	"landing_admin_backend/internal/services/advantage_photo"
	"landing_admin_backend/internal/services/advantages"
	"landing_admin_backend/internal/services/auth"
	"landing_admin_backend/internal/services/hod_photos"
	"landing_admin_backend/internal/services/months"
	"landing_admin_backend/internal/services/posters"
	"landing_admin_backend/internal/services/users"
	"landing_admin_backend/internal/services/years"
	"landing_admin_backend/pkg/token_manager"
)

type Services struct {
	Auth           auth.Service
	Users          users.Service
	Posters        posters.Service
	Years          years.Service
	Months         months.Service
	HodPhotos      hod_photos.Service
	Advantages     advantages.Service
	AdvantagePhoto advantage_photo.Service
}

func Setup(cfg *config.Config, repository *repository.Repository) *Services {
	tokenManager := token_manager.NewTokenManager(cfg.TokenSecret, cfg.TokenTTL, cfg.RefreshTokenTTL)
	return &Services{
		Auth:           auth.NewService(repository, tokenManager),
		Users:          users.NewService(repository),
		Posters:        posters.NewService(repository, cfg),
		Years:          years.NewService(repository, cfg),
		Months:         months.NewService(repository, cfg),
		HodPhotos:      hod_photos.NewService(repository, cfg),
		Advantages:     advantages.NewService(repository, cfg),
		AdvantagePhoto: advantage_photo.NewService(repository, cfg),
	}
}
