package settings

import (
	"context"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"
)

type Settings interface {
	Get(ctx context.Context) (settings []*domain.Setting, err error)
	CreateOrUpdate(ctx context.Context, setting domain.Setting) (settingRes *domain.Setting, err error)
}

type service struct {
	repository *repository.Repository
	cfg        *config.Config
}

func NewService(repository *repository.Repository, cfg *config.Config) Settings {
	return &service{
		repository: repository,
		cfg:        cfg,
	}
}

func (s *service) Get(ctx context.Context) (settings []*domain.Setting, err error) {
	return s.repository.Settings.Get(ctx)
}

func (s *service) CreateOrUpdate(ctx context.Context, setting domain.Setting) (settingRes *domain.Setting, err error) {
	_, err = s.repository.Settings.GetByID(ctx, setting.ID)
	if err != nil {
		err = s.repository.Settings.Create(ctx, setting)
		return
	}
	settingRes = &setting
	return settingRes, s.repository.Settings.Update(ctx, setting)
}
