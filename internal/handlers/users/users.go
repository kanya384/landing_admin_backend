package users

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"time"

	"github.com/go-openapi/runtime/middleware"
)

type Handlers interface {
	Create(params operations.PutUsersParams) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Create(params operations.PutUsersParams) middleware.Responder {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	user, err := domain.NewUser("", params.Params.Name, params.Params.Login, params.Params.Pass, params.Params.Role)
	if err != nil {
		return operations.NewPutUsersBadRequest().WithPayload(&models.ResultResponse{Msg: "invalid params params"})
	}
	err = h.services.Users.Create(ctx, user)
	if err != nil {
		return operations.NewPutUsersInternalServerError()
	}
	return operations.NewPutUsersOK().WithPayload(&models.ResultResponse{Msg: "success"})
}
