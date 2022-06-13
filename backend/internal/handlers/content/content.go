package content

import (
	"context"
	"landing_admin_backend/internal/generated/operations/content"
	"landing_admin_backend/internal/services"
	"time"

	"github.com/go-openapi/runtime/middleware"
)

type Handlers interface {
	Get(params content.GetContentParams) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params content.GetContentParams) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	contentL, err := h.services.LandingContent.Get(ctx)
	if err != nil {
		return content.NewGetContentBadRequest()
	}

	return content.NewGetContentOK().WithPayload(&contentL)
}
