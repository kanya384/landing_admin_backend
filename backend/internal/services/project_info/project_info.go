package project_info

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
	IMAGE_WIDTH  = 1280
	IMAGE_HEIGHT = 822
)

type ProjectInfo interface {
	Get(ctx context.Context, filter map[string]interface{}) (projectInfos []*domain.ProjectInfo, err error)
	GetByID(ctx context.Context, id primitive.ObjectID) (projectInfo domain.ProjectInfo, err error)
	Create(ctx context.Context, projectInfo domain.ProjectInfo, file io.ReadCloser) (projectInfoRes domain.ProjectInfo, err error)
	Update(ctx context.Context, projectInfo domain.ProjectInfo, file interface{}) (projectInfoRes domain.ProjectInfo, err error)
	Delete(ctx context.Context, projectInfoID primitive.ObjectID) (err error)
	UpdateOrder(ctx context.Context, first domain.UpdateOrder, second domain.UpdateOrder) (err error)
}

type service struct {
	repository *repository.Repository
	cfg        *config.Config
}

func NewService(repository *repository.Repository, cfg *config.Config) ProjectInfo {
	return &service{
		repository: repository,
		cfg:        cfg,
	}
}

func (s *service) Get(ctx context.Context, filter map[string]interface{}) (projectInfos []*domain.ProjectInfo, err error) {
	return s.repository.ProjectInfo.Get(ctx, filter)
}

func (s *service) GetByID(ctx context.Context, id primitive.ObjectID) (projectInfo domain.ProjectInfo, err error) {
	return s.repository.ProjectInfo.GetByID(ctx, id)
}

func (s *service) Create(ctx context.Context, projectInfo domain.ProjectInfo, file io.ReadCloser) (projectInfoRes domain.ProjectInfo, err error) {
	filename, err := helpers.ProcessImage(file, s.cfg.FileStore, IMAGE_WIDTH, IMAGE_HEIGHT)
	if err != nil {
		return
	}
	projectInfo.Photo = filename
	projectInfoRes = projectInfo
	err = s.repository.ProjectInfo.Create(context.Background(), projectInfo)
	return
}

func (s *service) Update(ctx context.Context, projectInfo domain.ProjectInfo, file interface{}) (projectInfoRes domain.ProjectInfo, err error) {
	filename, errIm := helpers.ProcessImage(file, s.cfg.FileStore, IMAGE_WIDTH, IMAGE_HEIGHT)
	if errIm == nil {
		//need to delete old file
		projectInfo.Photo = filename
	}
	err = s.repository.ProjectInfo.Update(ctx, projectInfo)
	projectInfoRes = projectInfo
	return
}

func (s *service) Delete(ctx context.Context, projectInfoID primitive.ObjectID) (err error) {
	return s.repository.ProjectInfo.Delete(ctx, projectInfoID)
}

func (s *service) UpdateOrder(ctx context.Context, first domain.UpdateOrder, second domain.UpdateOrder) (err error) {
	err = s.repository.ProjectInfo.UpdateOrder(ctx, first.ID, first.Order)
	if err != nil {
		return
	}
	err = s.repository.ProjectInfo.UpdateOrder(ctx, second.ID, second.Order)
	if err != nil {
		return
	}
	return
}
