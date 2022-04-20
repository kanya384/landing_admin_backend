package posters

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Get(ctx context.Context, filter map[string]interface{}) (posters []domain.Poster, err error)
	GetByID(ctx context.Context, id primitive.ObjectID) (poster domain.Poster, err error)
	Create(ctx context.Context, poster domain.Poster) (err error)
	Update(ctx context.Context, poster domain.Poster) (err error)
	Delete(ctx context.Context, posterID primitive.ObjectID) (err error)
}

type service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Get(ctx context.Context, filter map[string]interface{}) (posters []domain.Poster, err error) {
	return
}

func (s *service) GetByID(ctx context.Context, id primitive.ObjectID) (poster domain.Poster, err error) {
	return
}

func (s *service) Create(ctx context.Context, poster domain.Poster) (err error) {
	return
}

func (s *service) Update(ctx context.Context, poster domain.Poster) (err error) {
	return
}

func (s *service) Delete(ctx context.Context, posterID primitive.ObjectID) (err error) {
	return
}
