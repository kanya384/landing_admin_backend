package leads

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/domain"
	"net/http"
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
	err = s.repository.Leads.Create(ctx, lead)
	if err != nil {
		return
	}
	err = s.sendLeadToCrm(lead)
	return
}

func (s *service) Delete(ctx context.Context, leadID primitive.ObjectID) (err error) {
	return s.repository.Leads.Delete(ctx, leadID)
}

func (s *service) sendLeadToCrm(lead domain.Lead) (err error) {
	//TODO: добавить Источник (заполняется по таблице соответствия меток utm_source и источников, приложим отдельно)
	//TODO: добавить Рекламное агентство (заполняется согласно таблице соответствия меток utm_source, приложим отдельно) UF_CRM_1648203203
	body := map[string]interface{}{
		"fields": map[string]interface{}{
			"TITLE": fmt.Sprintf(`Заявка с сайта %s`, s.cfg.AppHost),
			"NAME":  lead.Name,
			"PHONE": map[string]interface{}{
				"n0": map[string]string{
					"VALUE":      lead.Phone,
					"VALUE_TYPE": "WORK",
				},
			},
			"EMAIL": map[string]interface{}{
				"n0": map[string]string{
					"VALUE":      lead.Email,
					"VALUE_TYPE": "WORK",
				},
			},
			"COMMENTS":          lead.Text,
			"ASSIGNED_BY_ID":    4134,
			"UF_CRM_1520270491": lead.Roistat,
			"UTM_SOURCE":        lead.Utm_source,
			"UTM_MEDIUM":        lead.Utm_medium,
			"UTM_CONTENT":       lead.Utm_content,
			"UTM_CAMPAIGN":      lead.Utm_campaign,
			"UTM_TERM":          lead.Utm_term,
			"UF_CRM_1648203203": map[string]interface{}{
				"n0": map[string]string{
					"VALUE":      s.cfg.AppHost,
					"VALUE_TYPE": "WORK",
				},
			},
			"UF_CRM_1648126761": 4054, //объект
			"UF_CRM_1593781584": 2773, //направление,
		},
		"params": map[string]interface{}{
			"REGISTER_SONET_EVENT": "Y",
		},
	}

	fmt.Println(body)
	jsonData, err := json.Marshal(body)

	if err != nil {
		return err
	}
	resp, err := http.Post(s.cfg.BitrixHook+"crm.lead.add.json", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res)
	return nil
}
