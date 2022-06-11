package years

import (
	"context"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"
	"landing_admin_backend/pkg/helpers"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	IMAGE_WIDTH  = 600
	IMAGE_HEIGHT = 350
)

type Years interface {
	Get(ctx context.Context) (years []domain.Year, err error)
	Create(ctx context.Context, year domain.Year) (err error)
	Update(ctx context.Context, year domain.Year) (err error)
	Delete(ctx context.Context, yearID primitive.ObjectID) (err error)
}

type service struct {
	repository *repository.Repository
	cfg        *config.Config
}

func NewService(repository *repository.Repository, cfg *config.Config) Years {
	return &service{
		repository: repository,
		cfg:        cfg,
	}
}

func (s *service) Get(ctx context.Context) (years []domain.Year, err error) {
	return s.repository.Year.Get(ctx)
}

func (s *service) Create(ctx context.Context, year domain.Year) (err error) {
	return s.repository.Year.Create(context.Background(), year)
}

func (s *service) Update(ctx context.Context, year domain.Year) (err error) {
	return s.repository.Year.Update(ctx, year)
}

func (s *service) Delete(ctx context.Context, yearID primitive.ObjectID) (err error) {
	months, err := s.repository.Month.Get(ctx, yearID)
	if err != nil && err != mongo.ErrNoDocuments {
		return
	}
	imagesToDelete := []string{}

	if err != mongo.ErrNoDocuments {
		for _, month := range months {
			hodPhotos, _ := s.repository.HodPhotos.Get(ctx, month.ID)
			for _, photo := range hodPhotos {
				imagesToDelete = append(imagesToDelete, photo.Image)
			}
		}
	}

	err = s.repository.Year.Delete(ctx, yearID)
	//удаляем фотки
	if err == nil {
		for _, photo := range imagesToDelete {
			err = helpers.DeleteFile(photo, "file_store")
		}
	}
	return
}
