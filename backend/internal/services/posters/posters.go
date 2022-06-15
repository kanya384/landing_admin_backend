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
	IMAGE_WIDTH  = 570
	IMAGE_HEIGHT = 720
)

type Posters interface {
	Get(ctx context.Context, filter map[string]interface{}) (posters []*domain.Poster, err error)
	GetByID(ctx context.Context, id primitive.ObjectID) (poster domain.Poster, err error)
	Create(ctx context.Context, poster domain.Poster, file io.ReadCloser) (posterRes domain.Poster, err error)
	Update(ctx context.Context, poster domain.Poster, file interface{}) (posterRes domain.Poster, err error)
	Delete(ctx context.Context, posterID primitive.ObjectID) (err error)
	UpdateOrder(ctx context.Context, first domain.UpdateOrder, second domain.UpdateOrder) (err error)
}

type service struct {
	repository *repository.Repository
	cfg        *config.Config
}

func NewService(repository *repository.Repository, cfg *config.Config) Posters {
	return &service{
		repository: repository,
		cfg:        cfg,
	}
}

func (s *service) Get(ctx context.Context, filter map[string]interface{}) (posters []*domain.Poster, err error) {
	return s.repository.Poster.Get(ctx, filter)
}

func (s *service) GetByID(ctx context.Context, id primitive.ObjectID) (poster domain.Poster, err error) {
	return s.repository.Poster.GetByID(ctx, id)
}

func (s *service) Create(ctx context.Context, poster domain.Poster, file io.ReadCloser) (posterRes domain.Poster, err error) {
	filename, err := helpers.ProcessImage(file, s.cfg.FileStore, IMAGE_WIDTH, IMAGE_HEIGHT)
	if err != nil {
		return
	}
	poster.Photo = filename
	posterRes = poster
	err = s.repository.Poster.Create(context.Background(), poster)
	return
}

func (s *service) Update(ctx context.Context, poster domain.Poster, file interface{}) (posterRes domain.Poster, err error) {
	filename, errIm := helpers.ProcessImage(file, s.cfg.FileStore, IMAGE_WIDTH, IMAGE_HEIGHT)
	if errIm == nil {
		//need to delete old file
		poster.Photo = filename
	}
	err = s.repository.Poster.Update(ctx, poster)
	posterRes = poster
	return
}

func (s *service) Delete(ctx context.Context, posterID primitive.ObjectID) (err error) {
	return s.repository.Poster.Delete(ctx, posterID)
}

func (s *service) UpdateOrder(ctx context.Context, first domain.UpdateOrder, second domain.UpdateOrder) (err error) {
	err = s.repository.Poster.UpdateOrder(ctx, first.ID, first.Order)
	if err != nil {
		return
	}
	err = s.repository.Poster.UpdateOrder(ctx, second.ID, second.Order)
	if err != nil {
		return
	}
	return
}
