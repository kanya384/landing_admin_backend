package leads

import (
	"context"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"
	"time"

	"landing_admin_backend/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Leads interface {
	Get(ctx context.Context) (leads []*domain.Lead, err error)
	GetAnalytics(ctx context.Context) (analytics domain.Analytics, err error)
	Create(ctx context.Context, lead domain.Lead) (err error)
	Delete(ctx context.Context, leadID primitive.ObjectID) (err error)
}

type service struct {
	repository *repository.Repository
	cfg        *config.Config
}

func NewService(repository *repository.Repository, cfg *config.Config) Leads {
	return &service{
		repository: repository,
		cfg:        cfg,
	}
}

func (s *service) Get(ctx context.Context) (leads []*domain.Lead, err error) {
	return s.repository.Leads.Get(ctx)
}

func (s *service) GetAnalytics(ctx context.Context) (analytics domain.Analytics, err error) {
	now := time.Now()
	startOfTheDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfTheDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 0, 0, now.Location())
	startOfTheMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	todayCount, err := s.repository.Leads.CountDocuments(ctx, startOfTheDay, endOfTheDay)
	if err != nil {
		return
	}
	monthCount, err := s.repository.Leads.CountDocuments(ctx, startOfTheMonth, endOfTheDay)
	chartInfo := []*domain.DayLeadsInfo{}
	tmpDay := startOfTheMonth
	for tmpDay.Before(endOfTheDay) {
		tmpEndDay := time.Date(now.Year(), now.Month(), tmpDay.Day(), 23, 59, 0, 0, now.Location())
		dayCount, _ := s.repository.Leads.CountDocuments(ctx, tmpDay, tmpEndDay)
		info := domain.DayLeadsInfo{
			Date:  tmpDay.Format("01.02"),
			Count: dayCount,
		}

		chartInfo = append(chartInfo, &info)
		tmpDay = tmpDay.Add(time.Hour * 24)
	}
	analytics = domain.Analytics{
		TodayCount: todayCount,
		MonthCount: monthCount,
		ChartInfo:  chartInfo,
	}
	if err != nil {
		return
	}
	return
}

func (s *service) Create(ctx context.Context, lead domain.Lead) (err error) {
	//send lead to crm
	return s.repository.Leads.Create(ctx, lead)
}

func (s *service) Delete(ctx context.Context, leadID primitive.ObjectID) (err error) {
	return s.repository.Leads.Delete(ctx, leadID)
}
