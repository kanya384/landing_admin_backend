package posters

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations/posters"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"landing_admin_backend/pkg/helpers"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handlers interface {
	Get(params posters.GetPostersParams) middleware.Responder
	GetPosterById(params posters.GetPostersPosterIDParams, input interface{}) middleware.Responder
	Create(params posters.PutPostersParams, input interface{}) middleware.Responder
	Update(params posters.PatchPostersParams, input interface{}) middleware.Responder
	Delete(params posters.DeletePostersPosterIDParams, input interface{}) middleware.Responder
	PostersOrdersChange(params posters.PostPostersOrdersParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Create(params posters.PutPostersParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return posters.NewPutPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}

	poster, err := domain.NewPoster("", params.Title, params.Description, "", true, 0, time.Now(), time.Now(), userID)
	if err != nil {
		return posters.NewPutPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating poster"})
	}

	posterRes, err := h.services.Posters.Create(ctx, poster, params.File)
	if err != nil {
		return posters.NewPutPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating poster"})
	}
	resPoster := models.Poster{
		ID:          posterRes.ID.Hex(),
		Title:       posterRes.Title,
		Description: posterRes.Description,
		CreatedAt:   strfmt.DateTime(posterRes.CreatedAt),
		UpdatedAt:   strfmt.DateTime(posterRes.UpdateAt),
		Order:       int64(posterRes.Order),
		Photo:       posterRes.Photo,
		ModifiedBy:  posterRes.ModifiedBy.Hex(),
	}
	return posters.NewPutPostersOK().WithPayload(&resPoster)
}

func (h *handlers) Update(params posters.PatchPostersParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return posters.NewPatchPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}
	photo := ""
	if params.Photo != nil {
		photo = *params.Photo
	}
	poster, err := domain.NewPoster(params.ID, params.Title, params.Description, photo, params.Active, int(params.Order), time.Now(), time.Now(), userID)
	if err != nil {
		return posters.NewPatchPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "error updating poster"})
	}
	err = h.services.Posters.Update(ctx, poster, params.File)
	if err != nil {
		return posters.NewPatchPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "error updating poster"})
	}

	return posters.NewPatchPostersOK().WithPayload(&models.ResultResponse{Msg: "poster created"})
}

func (h *handlers) Delete(params posters.DeletePostersPosterIDParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	posterID, err := primitive.ObjectIDFromHex(params.PosterID)
	if err != nil {
		return posters.NewDeletePostersPosterIDBadRequest()
	}
	err = h.services.Posters.Delete(ctx, posterID)
	if err != nil {
		return posters.NewDeletePostersPosterIDBadRequest()
	}
	return posters.NewDeletePostersPosterIDOK().WithPayload(&models.ResultResponse{Msg: "poster deleted"})
}

func (h *handlers) Get(params posters.GetPostersParams) middleware.Responder {
	filter := map[string]interface{}{
		"active": "",
		"name":   "",
	}
	postersList, err := h.services.Posters.Get(context.Background(), filter)
	postersListResponse := models.GetPostersResponse{}
	for _, pst := range postersList {
		postIns := models.Poster{
			ID:          pst.ID.Hex(),
			Title:       pst.Title,
			Description: pst.Description,
			Photo:       pst.Photo,
			ModifiedBy:  pst.ModifiedBy.Hex(),
			Active:      pst.Active,
			Order:       int64(pst.Order),
			UpdatedAt:   strfmt.DateTime(pst.UpdateAt),
			CreatedAt:   strfmt.DateTime(pst.CreatedAt),
		}
		postersListResponse = append(postersListResponse, &postIns)
	}
	if err != nil {
		return posters.NewGetPostersBadRequest()
	}
	return posters.NewGetPostersOK().WithPayload(postersListResponse)
}

func (h *handlers) GetPosterById(params posters.GetPostersPosterIDParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	posterID, err := primitive.ObjectIDFromHex(params.PosterID)
	if err != nil {
		return posters.NewDeletePostersPosterIDBadRequest()
	}
	poster, err := h.services.Posters.GetByID(ctx, posterID)
	if err != nil {
		return posters.NewDeletePostersPosterIDBadRequest()
	}
	res := models.Poster{
		ID:          poster.ID.Hex(),
		Title:       poster.Title,
		Description: poster.Description,
		Photo:       poster.Photo,
		Active:      poster.Active,
		Order:       int64(poster.Order),
		CreatedAt:   strfmt.DateTime(poster.CreatedAt),
		UpdatedAt:   strfmt.DateTime(poster.UpdateAt),
		ModifiedBy:  poster.ModifiedBy.Hex(),
	}
	return posters.NewGetPostersPosterIDOK().WithPayload(&res)
}

func (h *handlers) PostersOrdersChange(params posters.PostPostersOrdersParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	first, err := domain.NewUpdateOrderStruct(params.Params.First.ID, int(params.Params.First.Order))
	if err != nil {
		return posters.NewPostPostersOrdersBadRequest()
	}
	second, err := domain.NewUpdateOrderStruct(params.Params.Second.ID, int(params.Params.Second.Order))
	if err != nil {
		return posters.NewPostPostersOrdersBadRequest()
	}
	err = h.services.Posters.PostersOrdersChange(ctx, first, second)
	if err != nil {
		return posters.NewPostPostersOrdersInternalServerError()
	}
	return posters.NewPostPostersOrdersOK()
}
