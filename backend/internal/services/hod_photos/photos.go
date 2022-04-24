package hod_photos

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
	Get(ctx context.Context, monthID primitive.ObjectID) (hodPhotos []domain.HodPhoto, err error)
	Create(ctx context.Context, hodPhoto domain.HodPhoto, file io.ReadCloser) (hodPhotoRes domain.HodPhoto, err error)
	Delete(ctx context.Context, hodPhotoID primitive.ObjectID) (err error)
	UpdateOrder(ctx context.Context, first domain.UpdateOrder, second domain.UpdateOrder) (err error)
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

func (s *service) Get(ctx context.Context, monthID primitive.ObjectID) (hodPhotos []domain.HodPhoto, err error) {
	return s.repository.HodPhotos.Get(ctx, monthID)
}

func (s *service) Create(ctx context.Context, hodPhoto domain.HodPhoto, file io.ReadCloser) (hodPhotoRes domain.HodPhoto, err error) {
	filename, err := helpers.ProcessImage(file, s.cfg.FileStore, IMAGE_WIDTH, IMAGE_HEIGHT)
	if err != nil {
		return
	}
	hodPhoto.Image = filename
	hodPhotoRes = hodPhoto
	err = s.repository.HodPhotos.Create(ctx, hodPhoto)
	return
}

func (s *service) Delete(ctx context.Context, hodPhotoID primitive.ObjectID) (err error) {
	return s.repository.HodPhotos.Delete(ctx, hodPhotoID)
}

func (s *service) UpdateOrder(ctx context.Context, first domain.UpdateOrder, second domain.UpdateOrder) (err error) {
	err = s.repository.HodPhotos.UpdateOrder(ctx, first.ID, first.Order)
	if err != nil {
		return
	}
	err = s.repository.HodPhotos.UpdateOrder(ctx, second.ID, second.Order)
	if err != nil {
		return
	}
	return
}
