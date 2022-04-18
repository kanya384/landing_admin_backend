package auths

import (
	"context"
	"landing_admin_backend/internal/generated/operations"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"time"

	"github.com/go-openapi/runtime/middleware"
)

type Handlers interface {
	Authenticate(params operations.PostLoginParams) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Authenticate(params operations.PostLoginParams) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	tokens, err := h.services.Auth.Authenticate(ctx, params.Params.Login, params.Params.Pass)
	if err != nil {
		return operations.NewPostLoginBadRequest().WithPayload(&models.ResultResponse{Msg: "invalid params"})
	}

	return operations.NewPostLoginOK().WithPayload(&models.AuthenticateResponse{Token: tokens.GetToken(), RefreshToken: tokens.GetRefreshToken()})
}
