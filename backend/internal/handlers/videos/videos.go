package video

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations/video"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"landing_admin_backend/pkg/helpers"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handlers interface {
	Get(params video.GetVideoParams, input interface{}) middleware.Responder
	Create(params video.PutVideoParams, input interface{}) middleware.Responder
	Update(params video.PatchVideoParams, input interface{}) middleware.Responder
	Delete(params video.DeleteVideoIDParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params video.GetVideoParams, input interface{}) middleware.Responder {
	videoList, err := h.services.Video.Get(context.Background())
	if err != nil {
		return video.NewGetVideoBadRequest()
	}
	videoListResponse := []*models.Video{}
	for _, video := range videoList {
		videoIns := models.Video{
			ID:         video.ID.Hex(),
			Preview:    video.Preview,
			URL:        video.URL,
			UpdatedAt:  strfmt.DateTime(video.UpdateAt),
			CreatedAt:  strfmt.DateTime(video.CreatedAt),
			ModifiedBy: video.ModifiedBy.Hex(),
		}
		videoListResponse = append(videoListResponse, &videoIns)
	}
	if err != nil {
		return video.NewGetVideoBadRequest()
	}
	return video.NewGetVideoOK().WithPayload(videoListResponse)
}

func (h *handlers) Create(params video.PutVideoParams, input interface{}) middleware.Responder {
	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return video.NewPutVideoBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}
	vid, err := domain.NewVideo("", params.URL, "", time.Now(), time.Now(), userID)
	if err != nil {
		return video.NewPutVideoBadRequest().WithPayload(&models.ResultResponse{Msg: "check request params"})
	}
	videoInserted, err := h.services.Video.Create(context.Background(), vid, params.File, "file.pdf")
	if err != nil {
		return video.NewPutVideoInternalServerError()
	}
	videoRes := models.Video{
		ID:         videoInserted.ID.Hex(),
		Preview:    videoInserted.Preview,
		URL:        videoInserted.URL,
		CreatedAt:  strfmt.DateTime(videoInserted.CreatedAt),
		UpdatedAt:  strfmt.DateTime(videoInserted.UpdateAt),
		ModifiedBy: videoInserted.ModifiedBy.Hex(),
	}
	return video.NewPutVideoOK().WithPayload(&videoRes)
}

func (h *handlers) Update(params video.PatchVideoParams, input interface{}) middleware.Responder {
	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return video.NewPutVideoBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}
	vid, err := domain.NewVideo("", params.URL, "", time.Now(), time.Now(), userID)
	if err != nil {
		return video.NewPutVideoBadRequest().WithPayload(&models.ResultResponse{Msg: "check request params"})
	}
	videoInserted, err := h.services.Video.Create(context.Background(), vid, params.File, "file.pdf")
	if err != nil {
		return video.NewPutVideoInternalServerError()
	}
	videoRes := models.Video{
		ID:         videoInserted.ID.Hex(),
		Preview:    videoInserted.Preview,
		URL:        videoInserted.URL,
		CreatedAt:  strfmt.DateTime(videoInserted.CreatedAt),
		UpdatedAt:  strfmt.DateTime(videoInserted.UpdateAt),
		ModifiedBy: videoInserted.ModifiedBy.Hex(),
	}
	return video.NewPatchVideoOK().WithPayload(&videoRes)
}

func (h *handlers) Delete(params video.DeleteVideoIDParams, input interface{}) middleware.Responder {
	id, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		return video.NewDeleteVideoIDBadRequest()
	}
	err = h.services.Docs.Delete(context.Background(), id)
	if err != nil {
		return video.NewDeleteVideoIDInternalServerError()
	}
	return video.NewDeleteVideoIDOK()
}
