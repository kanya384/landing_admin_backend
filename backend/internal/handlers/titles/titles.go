package titles

import (
	"context"
	"landing_admin_backend/internal/domain"
	operations "landing_admin_backend/internal/generated/operations/titles"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"landing_admin_backend/pkg/helpers"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type Handlers interface {
	Get(params operations.GetTitlesParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params operations.GetTitlesParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	titles, err := h.services.Titles.Get(c)

	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return titles.NewPutPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}

	poster, err := domain.NewPoster("", params, "", 0, time.Now(), time.Now(), userID)
	if err != nil {
		return titles.NewPutPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating poster"})
	}

	posterRes, err := h.services.Posters.Create(ctx, poster, params.File)
	if err != nil {
		return titles.NewPutPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating poster"})
	}
	resPoster := models.Poster{
		ID:         posterRes.ID.Hex(),
		Title:      posterRes.Title,
		CreatedAt:  strfmt.DateTime(posterRes.CreatedAt),
		UpdatedAt:  strfmt.DateTime(posterRes.UpdateAt),
		Order:      int64(posterRes.Order),
		Photo:      posterRes.Photo,
		ModifiedBy: posterRes.ModifiedBy.Hex(),
	}
	return titles.NewPutPostersOK().WithPayload(&resPoster)
}
