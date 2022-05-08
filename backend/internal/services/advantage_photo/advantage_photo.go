package advantage_photo

import (
	"context"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Create(ctx context.Context, advantagePhoto domain.AdvantagePhoto) (err error)
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

func (s *service) Create(ctx context.Context, advantagePhoto domain.AdvantagePhoto) (err error) {
	return s.repository.AdvantagePhoto.Create(ctx, advantagePhoto)
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
