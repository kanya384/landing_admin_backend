package users

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"strings"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	user, err := domain.NewUser("", params.Params.Name, params.Params.Login, params.Params.Pass, params.Params.Role)
	if err != nil {
		return operations.NewPutUsersBadRequest().WithPayload(&models.ResultResponse{Msg: "invalid params params"})
	}
	err = h.services.Users.Create(ctx, user)
	if err != nil {
		if strings.Contains(err.Error(), "E11000 duplicate key error collection") {
			return operations.NewPutUsersBadRequest().WithPayload(&models.ResultResponse{Msg: "user with this login already exists"})
		}
		return operations.NewPutUsersInternalServerError()
	}
	return operations.NewPutUsersOK().WithPayload(&models.ResultResponse{Msg: "success"})
}
