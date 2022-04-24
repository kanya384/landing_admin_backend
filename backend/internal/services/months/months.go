package months

import (
	"context"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"
	"landing_admin_backend/pkg/helpers"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	IMAGE_WIDTH  = 600
	IMAGE_HEIGHT = 350
)

type Service interface {
	Get(ctx context.Context, yearID primitive.ObjectID) (months []domain.Month, err error)
	Create(ctx context.Context, month domain.Month) (err error)
	Update(ctx context.Context, month domain.Month) (err error)
	Delete(ctx context.Context, monthID primitive.ObjectID) (err error)
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

func (s *service) Get(ctx context.Context, yearID primitive.ObjectID) (months []domain.Month, err error) {
	return s.repository.Month.Get(ctx, yearID)
}

func (s *service) Create(ctx context.Context, month domain.Month) (err error) {
	return s.repository.Month.Create(context.Background(), month)
}

func (s *service) Update(ctx context.Context, month domain.Month) (err error) {
	return s.repository.Month.Update(ctx, month)
}

func (s *service) Delete(ctx context.Context, monthID primitive.ObjectID) (err error) {
	hodPhotos, _ := s.repository.HodPhotos.Get(ctx, monthID)
	for _, photo := range hodPhotos {
		helpers.DeleteFile(photo.Image, "file_store")
	}
	err = s.repository.Month.Delete(ctx, monthID)
	return
}
