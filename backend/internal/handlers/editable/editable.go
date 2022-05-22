package editable

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations/editable"
	"landing_admin_backend/internal/services"
	"time"

	"github.com/go-openapi/runtime/middleware"
)

//

type Handlers interface {
	CreateOrUpdate(params editable.PostEditableParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) CreateOrUpdate(params editable.PostEditableParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	editables := params.Params

	for _, edbl := range editables {
		item, err := domain.NewEditable(edbl.ID, int(edbl.Type), edbl.Value, int(edbl.Width), int(edbl.Height))
		if err != nil {
			return editable.NewPostEditableBadRequest()
		}
		err = h.services.Editable.CreateOrUpdate(ctx, item)
		if err != nil {
			return editable.NewPostEditableBadRequest()
		}
	}
	return editable.NewPostEditableOK()
}
