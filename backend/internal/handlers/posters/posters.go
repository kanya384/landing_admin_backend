package posters

import (
	"context"
	"fmt"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations/posters"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"landing_admin_backend/pkg/helpers"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
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

	err = h.services.Posters.Create(ctx, poster, params.File)
	if err != nil {
		fmt.Println(err)
		return posters.NewPutPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating poster"})
	}

	return posters.NewPutPostersOK().WithPayload(&models.ResultResponse{Msg: "poster created"})
}

func (h *handlers) Update(params posters.PostPostersParams, input interface{}) middleware.Responder {
	/*poster := domain.Poster{
		ID:          params.ID,
		Title:       params.Title,
		Description: params.Description,
		Active:      true,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}

	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return posters.NewPutPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}
	poster.ModifiedBy = userID

	image, err := helpers.ResizeImage(params.File, IMAGE_WIDTH, IMAGE_HEIGHT)
	if err != nil {
		return posters.NewPutPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "error converting image"})
	}

	defer params.File.Close()
	filename, err := helpers.SaveImageFile(image, "tmg.jpg", h.storePath)
	if err != nil {
		return posters.NewPutPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "error saving image"})
	}
	poster.Photo = filename
	err = h.services.Posters.Create(context.Background(), poster)
	if err != nil {
		return posters.NewPutPostersBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating poster"})
	}*/
	return posters.NewPutPostersOK().WithPayload(&models.ResultResponse{Msg: "poster created"})
}

func (h *handlers) Delete(params posters.DeletePostersPosterIDParams, input interface{}) middleware.Responder {
	return posters.NewDeletePostersPosterIDOK()
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
