package posters

import (
	"context"
	"io"
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
	Get(ctx context.Context, filter map[string]interface{}) (posters []*domain.Poster, err error)
	GetByID(ctx context.Context, id primitive.ObjectID) (poster domain.Poster, err error)
	Create(ctx context.Context, poster domain.Poster, file io.ReadCloser) (err error)
	Update(ctx context.Context, poster domain.Poster) (err error)
	Delete(ctx context.Context, posterID primitive.ObjectID) (err error)
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

func (s *service) Get(ctx context.Context, filter map[string]interface{}) (posters []*domain.Poster, err error) {
	return s.repository.Poster.Get(ctx, filter)
}

func (s *service) GetByID(ctx context.Context, id primitive.ObjectID) (poster domain.Poster, err error) {
	return
}

func (s *service) Create(ctx context.Context, poster domain.Poster, file io.ReadCloser) (err error) {
	image, err := helpers.ResizeImage(file, IMAGE_WIDTH, IMAGE_HEIGHT)
	if err != nil {
		return
	}
	defer file.Close()

	filename, err := helpers.SaveImageFile(image, "tmg.jpg", s.cfg.FileStore)
	if err != nil {
		return
	}
	poster.Photo = filename
	err = s.repository.Poster.Create(context.Background(), poster)
	if err != nil {
		return
	}
	return
}

func (s *service) Update(ctx context.Context, poster domain.Poster) (err error) {
	return
}

func (s *service) Delete(ctx context.Context, posterID primitive.ObjectID) (err error) {
	return
}
