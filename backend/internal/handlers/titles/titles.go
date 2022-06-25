package titles

import (
	"context"
	"landing_admin_backend/internal/domain"
	operations "landing_admin_backend/internal/generated/operations/titles"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handlers interface {
	Get(params operations.GetTitlesParams, input interface{}) middleware.Responder
	Create(params operations.PutTitlesParams, input interface{}) middleware.Responder
	Update(params operations.PatchTitlesParams, input interface{}) middleware.Responder
	Delete(params operations.DeleteTitlesIDParams, input interface{}) middleware.Responder
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

	titles, err := h.services.Titles.Get(ctx)
	if err != nil {
		return operations.NewGetTitlesBadRequest()
	}

	titlesRes := []*models.Title{}

	for _, title := range titles {
		titleInp := models.Title{
			ID:           title.ID.Hex(),
			DesktopTitle: title.DesktopTitle,
			MobileTitle:  title.MobileTitle,
			TagName:      title.TagName,
			TagValue:     title.TagValue,
		}

		titlesRes = append(titlesRes, &titleInp)
	}

	return operations.NewGetTitlesOK().WithPayload(titlesRes)
}

func (h *handlers) Create(params operations.PutTitlesParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	title, err := domain.NewTitle(params.Params.ID, params.Params.TagName, params.Params.TagValue, params.Params.DesktopTitle, params.Params.MobileTitle)
	if err != nil {
		return operations.NewPutTitlesBadRequest()
	}

	err = h.services.Titles.Create(ctx, title)
	if err != nil {
		return operations.NewPutTitlesInternalServerError()
	}

	titleRes := models.Title{
		ID:           title.ID.Hex(),
		DesktopTitle: title.DesktopTitle,
		MobileTitle:  title.MobileTitle,
		TagName:      title.TagName,
		TagValue:     title.TagValue,
	}

	return operations.NewPutTitlesOK().WithPayload(&titleRes)
}

func (h *handlers) Update(params operations.PatchTitlesParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	title, err := domain.NewTitle(params.Params.ID, params.Params.TagName, params.Params.TagValue, params.Params.DesktopTitle, params.Params.MobileTitle)
	if err != nil {
		return operations.NewPatchTitlesInternalServerError()
	}

	err = h.services.Titles.Update(ctx, title)
	if err != nil {
		return operations.NewPatchTitlesInternalServerError()
	}

	titleRes := models.Title{
		ID:           title.ID.Hex(),
		DesktopTitle: title.DesktopTitle,
		MobileTitle:  title.MobileTitle,
		TagName:      title.TagName,
		TagValue:     title.TagValue,
	}

	return operations.NewPutTitlesOK().WithPayload(&titleRes)
}

func (h *handlers) Delete(params operations.DeleteTitlesIDParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	posterID, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		return operations.NewDeleteTitlesIDBadRequest()
	}
	err = h.services.Titles.Delete(ctx, posterID)
	if err != nil {
		return operations.NewDeleteTitlesIDInternalServerError()
	}
	return operations.NewDeleteTitlesIDOK().WithPayload(&models.ResultResponse{Msg: "poster deleted"})
}
