package advantage_photo

import (
	"landing_admin_backend/internal/generated/operations/advantages"
	"landing_admin_backend/internal/services"

	"github.com/go-openapi/runtime/middleware"
)

type Handlers interface {
	Get(params advantages.GetAdvantagePhotoIDParams, input interface{}) middleware.Responder
	Create(params advantages.PutAdvantagePhotoParams, input interface{}) middleware.Responder
	Delete(params advantages.DeleteAdvantagePhotoIDParams, input interface{}) middleware.Responder
	ChangeAdvantagePhotosOrders(params advantages.PostAdvantagePhotoOrdersParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params advantages.GetAdvantagePhotoIDParams, input interface{}) middleware.Responder {
	return advantages.NewGetAdvantagePhotoIDOK()
}

func (h *handlers) Create(params advantages.PutAdvantagePhotoParams, input interface{}) middleware.Responder {
	return advantages.NewPutAdvantagePhotoOK()
}

func (h *handlers) Delete(params advantages.DeleteAdvantagePhotoIDParams, input interface{}) middleware.Responder {
	return advantages.NewDeleteAdvantagePhotoIDOK()
}

func (h *handlers) ChangeAdvantagePhotosOrders(params advantages.PostAdvantagePhotoOrdersParams, input interface{}) middleware.Responder {
	return advantages.NewPostAdvantagesOrdersOK()
}
