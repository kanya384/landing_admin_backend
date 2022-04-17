package users

import (
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	Get() (users []*domain.UserInfo, err error)
	Create(user domain.User) (err error)
	Update(user domain.User) (err error)
	Delete(userID primitive.ObjectID) (err error)
}

type service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Get() (users []*domain.UserInfo, err error) {
	return nil, nil
}

func (s *service) Create(user domain.User) (err error) {
	return nil
}

func (s *service) Update(user domain.User) (err error) {
	return nil
}

func (s *service) Delete(userID primitive.ObjectID) (err error) {
	return nil
}
