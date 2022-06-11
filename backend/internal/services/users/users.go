package users

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users interface {
	Get(ctx context.Context) (users []*domain.UserInfo, err error)
	Create(ctx context.Context, user domain.User) (err error)
	Update(ctx context.Context, user domain.User) (err error)
	Delete(ctx context.Context, userID primitive.ObjectID) (err error)
}

type service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) Users {
	return &service{
		repository: repository,
	}
}

func (s *service) Get(ctx context.Context) (users []*domain.UserInfo, err error) {
	return nil, nil
}

func (s *service) Create(ctx context.Context, user domain.User) (err error) {
	return s.repository.Users.Create(ctx, user)
}

func (s *service) Update(ctx context.Context, user domain.User) (err error) {
	return nil
}

func (s *service) Delete(ctx context.Context, userID primitive.ObjectID) (err error) {
	return nil
}
