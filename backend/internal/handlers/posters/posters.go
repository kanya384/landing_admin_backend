package posters

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations/posters"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handlers interface {
	Get(params posters.GetPostersParams) middleware.Responder
	Create(params posters.PutPostersParams, input interface{}) middleware.Responder
	Update(params posters.PostPostersParams, input interface{}) middleware.Responder
	Delete(params posters.DeletePostersPosterIDParams, input interface{}) middleware.Responder
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
	poster := domain.Poster{
		ID:          primitive.NewObjectID(),
		Title:       params.Title,
		Description: params.Description,
		//Photo:       params.Photo,
		Active:     true,
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
		ModifiedBy: primitive.NewObjectID(),
	}
	//image := params.File
	err := h.services.Posters.Create(context.Background(), poster)
	if err != nil {
		return posters.NewPutPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating poster"})
	}
	return posters.NewPutPostersOK().WithPayload(&models.ResultResponse{Msg: "poster created"})
}

func (h *handlers) Update(params posters.PostPostersParams, input interface{}) middleware.Responder {
	return posters.NewPostPostersOK()
}

func (h *handlers) Delete(params posters.DeletePostersPosterIDParams, input interface{}) middleware.Responder {
	return posters.NewDeletePostersPosterIDOK()
}

func (h *handlers) Get(params posters.GetPostersParams) middleware.Responder {
	filter := map[string]interface{}{
		"active": params.Params.Active,
		"name":   params.Params.Name,
	}
	postersList, err := h.services.Posters.Get(context.Background(), filter)
	postersListResponse := models.GetPostersResponse{}
	for _, pst := range postersList {
		postIns := models.Poster{
			ID:          pst.ID.Hex(),
			Title:       pst.Title,
			Description: pst.Description,
			Photo:       pst.Photo,
		}
		postersListResponse = append(postersListResponse, &postIns)
	}
	if err != nil {
		return posters.NewGetPostersBadRequest()
	}
	return posters.NewGetPostersOK().WithPayload(postersListResponse)
}
