package leads

import (
	"context"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"

	"landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Get(ctx context.Context) (leads []*domain.Lead, err error)
	Create(ctx context.Context, lead domain.Lead) (err error)
	Delete(ctx context.Context, leadID primitive.ObjectID) (err error)
}

type service struct {
	repository *repository.Repository
	cfg        *config.Config
}

func NewService(repository *repository.Repository, cfg *config.Config) Service {
	return &service{
		repository: repository,
		cfg:        cfg,
	}
}

func (s *service) Get(ctx context.Context) (leads []*domain.Lead, err error) {
	return s.repository.Leads.Get(ctx)
}

func (s *service) Create(ctx context.Context, lead domain.Lead) (err error) {
	return s.repository.Leads.Create(ctx, lead)
}

func (s *service) Delete(ctx context.Context, leadID primitive.ObjectID) (err error) {
	return s.repository.Leads.Delete(ctx, leadID)
}
