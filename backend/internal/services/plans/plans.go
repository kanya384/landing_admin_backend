package plans

import (
	"io"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/repository"
	"landing_admin_backend/pkg/helpers"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

const (
	IMAGE_WIDTH  = 600
	IMAGE_HEIGHT = 350
)

type Service interface {
	GetPlans(ctx context.Context) (plans []*domain.Plan, err error)
	GetActivePlans(ctx context.Context) (plans []*domain.Plan, err error)
	ProcessPlans(ctx context.Context, plansList []*domain.Plan) (err error)
	UpdatePlansPhoto(ctx context.Context, file io.Reader, id primitive.ObjectID) (plan domain.Plan, err error)
	UpdatePlansActivity(ctx context.Context, id primitive.ObjectID, flag bool) (err error)
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

func (s *service) GetPlans(ctx context.Context) (plans []*domain.Plan, err error) {
	return s.repository.Plans.Get(ctx, map[string]interface{}{})
}

func (s *service) GetActivePlans(ctx context.Context) (plans []*domain.Plan, err error) {
	return s.repository.Plans.Get(ctx, map[string]interface{}{"status": true})
}

func (s *service) ProcessPlans(ctx context.Context, plansList []*domain.Plan) (err error) {
	//по активности вопрос уточнить, будет приходить с ихнего бэкэнда или нет
	for _, plan := range plansList {
		plan, err := s.repository.Plans.GetByID(ctx, plan.ID)
		if err != nil {
			err = s.repository.Plans.Create(ctx, plan)
			if err != nil {
				return err
			}
		} else {
			err = s.repository.Plans.Update(ctx, plan)
			if err != nil {
				return err
			}
		}
	}
	return
}

func (s *service) UpdatePlansPhoto(ctx context.Context, file io.Reader, id primitive.ObjectID) (plan domain.Plan, err error) {
	plan, err = s.repository.Plans.GetByID(ctx, id)
	if err != nil {
		return
	}

	filename, errIm := helpers.ProcessImage(file, s.cfg.FileStore, IMAGE_WIDTH, IMAGE_HEIGHT)
	if errIm == nil {
		//need to delete old file
		plan.Image = filename
	}
	err = s.repository.Plans.Update(ctx, plan)
	return
}

func (s *service) UpdatePlansActivity(ctx context.Context, id primitive.ObjectID, flag bool) (err error) {
	plan, err := s.repository.Plans.GetByID(ctx, id)
	if err != nil {
		return
	}

	plan.Status = flag

	err = s.repository.Plans.Update(ctx, plan)
	return
}
