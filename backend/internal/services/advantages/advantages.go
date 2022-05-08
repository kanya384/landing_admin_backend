package advantages

import (
	"context"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"

	"landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Get(ctx context.Context) (advantages []*domain.Advantage, err error)
	GetByID(ctx context.Context, id primitive.ObjectID) (advantage domain.Advantage, err error)
	Create(ctx context.Context, advantage domain.Advantage) (err error)
	Update(ctx context.Context, advantage domain.Advantage) (err error)
	Delete(ctx context.Context, advantageID primitive.ObjectID) (err error)
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

func (s *service) Get(ctx context.Context) (advantages []*domain.Advantage, err error) {
	return s.repository.Advantages.Get(ctx)
}

func (s *service) GetByID(ctx context.Context, id primitive.ObjectID) (advantage domain.Advantage, err error) {
	return s.repository.Advantages.GetByID(ctx, id)
}

func (s *service) Create(ctx context.Context, advantage domain.Advantage) (err error) {
	return s.repository.Advantages.Create(ctx, advantage)
}

func (s *service) Update(ctx context.Context, advantage domain.Advantage) (err error) {
	return s.repository.Advantages.Update(ctx, advantage)
}

func (s *service) Delete(ctx context.Context, advantageID primitive.ObjectID) (err error) {
	return s.repository.Advantages.Delete(ctx, advantageID)
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
