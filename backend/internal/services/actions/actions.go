package actions

import (
	"context"
	"fmt"
	"io"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"
	"landing_admin_backend/pkg/helpers"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	IMAGE_PREVIEW_WIDTH  = 353
	IMAGE_PREVIEW_HEIGHT = 285
	IMAGE_WIDTH          = 413
	IMAGE_HEIGHT         = 1058
)

type Actions interface {
	Get(ctx context.Context, filter map[string]interface{}) (actions []*domain.Action, err error)
	GetByID(ctx context.Context, id primitive.ObjectID) (action domain.Action, err error)
	Create(ctx context.Context, action domain.Action, preview io.ReadCloser) (actionRes domain.Action, err error)
	Update(ctx context.Context, action domain.Action, file interface{}) (actionRes domain.Action, err error)
	Delete(ctx context.Context, actionID primitive.ObjectID) (err error)
	UpdateOrder(ctx context.Context, first domain.UpdateOrder, second domain.UpdateOrder) (err error)
}

type service struct {
	repository *repository.Repository
	cfg        *config.Config
}

func NewService(repository *repository.Repository, cfg *config.Config) Actions {
	return &service{
		repository: repository,
		cfg:        cfg,
	}
}

func (s *service) Get(ctx context.Context, filter map[string]interface{}) (actions []*domain.Action, err error) {
	return s.repository.Action.Get(ctx, filter)
}

func (s *service) GetByID(ctx context.Context, id primitive.ObjectID) (action domain.Action, err error) {
	return s.repository.Action.GetByID(ctx, id)
}

func (s *service) Create(ctx context.Context, action domain.Action, file io.ReadCloser) (actionRes domain.Action, err error) {
	photo, preview, err := helpers.ProcessImages(file, s.cfg.FileStore, IMAGE_WIDTH, IMAGE_HEIGHT, IMAGE_PREVIEW_WIDTH, IMAGE_PREVIEW_HEIGHT)
	if err != nil {
		return
	}

	action.Photo = photo
	action.Preview = preview

	err = s.repository.Action.Create(context.Background(), action)
	return action, err
}

func (s *service) Update(ctx context.Context, action domain.Action, file interface{}) (actionRes domain.Action, err error) {
	photo, preview, errIm := helpers.ProcessImages(file, s.cfg.FileStore, IMAGE_WIDTH, IMAGE_HEIGHT, IMAGE_PREVIEW_WIDTH, IMAGE_PREVIEW_HEIGHT)
	if err != nil {
		return
	}

	if errIm == nil {
		action.Photo = photo
		action.Preview = preview
	}
	fmt.Println(action)
	err = s.repository.Action.Update(ctx, action)
	actionRes = action
	return
}

func (s *service) Delete(ctx context.Context, actionID primitive.ObjectID) (err error) {
	return s.repository.Action.Delete(ctx, actionID)
}

func (s *service) UpdateOrder(ctx context.Context, first domain.UpdateOrder, second domain.UpdateOrder) (err error) {
	err = s.repository.Action.UpdateOrder(ctx, first.ID, first.Order)
	if err != nil {
		return
	}
	err = s.repository.Action.UpdateOrder(ctx, second.ID, second.Order)
	if err != nil {
		return
	}
	return
}
