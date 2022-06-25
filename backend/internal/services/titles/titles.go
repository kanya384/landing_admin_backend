package titles

import (
	"context"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Title interface {
	Get(ctx context.Context) (titles []*domain.Title, err error)
	Create(ctx context.Context, title domain.Title) error
	Update(ctx context.Context, title domain.Title) error
	Delete(ctx context.Context, titleID primitive.ObjectID) (err error)
}

type service struct {
	repository *repository.Repository
	cfg        *config.Config
}

func NewService(repository *repository.Repository, cfg *config.Config) Title {
	return &service{
		repository: repository,
		cfg:        cfg,
	}
}

func (s *service) Get(ctx context.Context) (titles []*domain.Title, err error) {
	return s.repository.Title.Get(ctx)
}

func (s *service) Create(ctx context.Context, title domain.Title) error {
	return s.repository.Title.Create(ctx, title)
}

func (s *service) Update(ctx context.Context, title domain.Title) error {
	return s.repository.Title.Update(ctx, title)
}

func (s *service) Delete(ctx context.Context, titleID primitive.ObjectID) (err error) {
	return s.repository.Title.Delete(ctx, titleID)
}
