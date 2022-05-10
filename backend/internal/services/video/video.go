package video

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
	Get(ctx context.Context) (videos []*domain.Video, err error)
	Create(ctx context.Context, video domain.Video, file io.ReadCloser, fileName string) (domain.Video, error)
	Update(ctx context.Context, video domain.Video, file interface{}, fileName string) (domain.Video, error)
	Delete(ctx context.Context, videoID primitive.ObjectID) (err error)
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

func (s *service) Get(ctx context.Context) (videos []*domain.Video, err error) {
	return s.repository.Video.Get(ctx)
}

func (s *service) Create(ctx context.Context, video domain.Video, file io.ReadCloser, fileName string) (domain.Video, error) {
	filename, err := helpers.ProcessImage(file, s.cfg.FileStore, IMAGE_WIDTH, IMAGE_HEIGHT)
	if err != nil {
		return video, err
	}
	video.Preview = filename
	err = s.repository.Video.Create(ctx, video)
	return video, err
}

func (s *service) Update(ctx context.Context, video domain.Video, file interface{}, fileName string) (domain.Video, error) {
	if fileName != "" {
		filename, err := helpers.ProcessImage(file, s.cfg.FileStore, IMAGE_WIDTH, IMAGE_HEIGHT)
		if err != nil {
			return video, err
		}
		video.Preview = filename
	}
	err := s.repository.Video.Update(ctx, video)
	return video, err
}

func (s *service) Delete(ctx context.Context, videoID primitive.ObjectID) (err error) {
	return s.repository.Video.Delete(ctx, videoID)
}
