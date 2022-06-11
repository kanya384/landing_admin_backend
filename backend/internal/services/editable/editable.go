package editable

import (
	"context"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"
)

type Editable interface {
	Get(ctx context.Context) (ediatables []*domain.Editable, err error)
	CreateOrUpdate(ctx context.Context, editable domain.Editable) (err error)
}

type service struct {
	repository *repository.Repository
	cfg        *config.Config
}

func NewService(repository *repository.Repository, cfg *config.Config) Editable {
	return &service{
		repository: repository,
		cfg:        cfg,
	}
}

func (s *service) Get(ctx context.Context) (ediatables []*domain.Editable, err error) {
	return s.repository.Editable.Get(ctx)
}

func (s *service) CreateOrUpdate(ctx context.Context, editable domain.Editable) (err error) {
	_, err = s.repository.Editable.GetByID(ctx, editable.ID)
	if err != nil {
		err = s.repository.Editable.Create(ctx, &editable)
		return
	}
	return s.repository.Editable.Update(ctx, &editable)
}
