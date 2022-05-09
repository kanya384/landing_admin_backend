package advantage_photo

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
	Get(ctx context.Context, advantageID primitive.ObjectID) (list []domain.AdvantagePhoto, err error)
	Create(ctx context.Context, advantagePhoto domain.AdvantagePhoto, file io.ReadCloser) (advantagePhotoRes domain.AdvantagePhoto, err error)
	Delete(ctx context.Context, advantagePhotoID primitive.ObjectID) (err error)
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

func (s *service) Get(ctx context.Context, advantageID primitive.ObjectID) (list []domain.AdvantagePhoto, err error) {
	return s.repository.AdvantagePhoto.Get(ctx, advantageID)
}

func (s *service) Create(ctx context.Context, advantagePhoto domain.AdvantagePhoto, file io.ReadCloser) (advantagePhotoRes domain.AdvantagePhoto, err error) {
	filename, err := helpers.ProcessImage(file, s.cfg.FileStore, IMAGE_WIDTH, IMAGE_HEIGHT)
	if err != nil {
		return
	}
	advantagePhoto.Image = filename
	advantagePhotoRes = advantagePhoto
	err = s.repository.AdvantagePhoto.Create(context.Background(), advantagePhoto)
	return
}

func (s *service) Delete(ctx context.Context, advantagePhotoID primitive.ObjectID) (err error) {
	return s.repository.AdvantagePhoto.Delete(ctx, advantagePhotoID)
}

func (s *service) UpdateOrder(ctx context.Context, first domain.UpdateOrder, second domain.UpdateOrder) (err error) {
	err = s.repository.Advantages.UpdateOrder(ctx, first.ID, first.Order)
	if err != nil {
		return
	}
	err = s.repository.Advantages.UpdateOrder(ctx, second.ID, second.Order)
	if err != nil {
		return
	}
	return
}
