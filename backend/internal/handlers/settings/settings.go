package settings

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations/settings"
	operations "landing_admin_backend/internal/generated/operations/settings"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"time"

	"github.com/go-openapi/runtime/middleware"
)

type Handlers interface {
	Get(params operations.GetSettingsParams, input interface{}) middleware.Responder
	Create(params operations.PutSettingsParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params settings.GetSettingsParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	settings, err := h.services.Settings.Get(ctx)
	if err != nil {
		return operations.NewGetSettingsBadRequest()
	}

	settingsRes := []*models.Setting{}

	for _, setting := range settings {
		settingsInp := models.Setting{
			ID:          setting.ID.Hex(),
			Name:        setting.Name,
			Description: setting.Description,
			Value:       int64(setting.Value),
		}

		settingsRes = append(settingsRes, &settingsInp)
	}

	return operations.NewGetSettingsOK().WithPayload(settingsRes)
}

func (h *handlers) Create(params operations.PutSettingsParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	setting, err := domain.NewSetting(params.Params.ID, params.Params.Name, params.Params.Description, int(params.Params.Value))
	if err != nil {
		return operations.NewPutSettingsBadRequest()
	}

	res, err := h.services.Settings.CreateOrUpdate(ctx, setting)
	if err != nil {
		return operations.NewPutSettingsInternalServerError()
	}

	settingRes := models.Setting{
		ID:          res.ID.Hex(),
		Name:        res.Name,
		Description: res.Description,
		Value:       int64(res.Value),
	}

	return operations.NewPutSettingsOK().WithPayload(&settingRes)
}
