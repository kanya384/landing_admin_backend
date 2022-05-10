package advantage_photo

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations/advantages"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"landing_admin_backend/pkg/helpers"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handlers interface {
	Get(params advantages.GetAdvantagePhotoIDParams, input interface{}) middleware.Responder
	Create(params advantages.PutAdvantagePhotoParams, input interface{}) middleware.Responder
	Delete(params advantages.DeleteAdvantagePhotoIDParams, input interface{}) middleware.Responder
	ChangeAdvantagePhotosOrders(params advantages.PostAdvantagePhotoOrdersParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params advantages.GetAdvantagePhotoIDParams, input interface{}) middleware.Responder {
	id, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		return advantages.NewGetAdvantagesIDBadRequest()
	}
	advantagesPhotoList, err := h.services.AdvantagePhoto.Get(context.Background(), id)
	if err != nil {
		return advantages.NewGetAdvantagesIDBadRequest()
	}
	advantageListResponse := []*models.AdvantagePhoto{}
	for _, item := range advantagesPhotoList {
		advIns := models.AdvantagePhoto{
			ID:        item.ID.Hex(),
			Image:     item.Image,
			Order:     int64(item.Order),
			UpdatedAt: strfmt.DateTime(item.UpdateAt),
			CreatedAt: strfmt.DateTime(item.CreatedAt),
		}
		advantageListResponse = append(advantageListResponse, &advIns)
	}
	if err != nil {
		return advantages.NewGetAdvantagesInternalServerError()
	}
	return advantages.NewGetAdvantagePhotoIDOK().WithPayload(advantageListResponse)
}

func (h *handlers) Create(params advantages.PutAdvantagePhotoParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return advantages.NewPutAdvantagePhotoBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}

	advantagePhoto, err := domain.NewAdvantagePhoto("", params.AdvantageID, "", 0, time.Now(), time.Now(), userID)
	if err != nil {
		return advantages.NewPutAdvantagePhotoBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating advantage"})
	}

	advantageRes, err := h.services.AdvantagePhoto.Create(ctx, advantagePhoto, params.File)
	if err != nil {
		return advantages.NewPutAdvantagePhotoBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating advantage"})
	}
	res := models.AdvantagePhoto{
		ID:        advantageRes.ID.Hex(),
		CreatedAt: strfmt.DateTime(advantageRes.CreatedAt),
		UpdatedAt: strfmt.DateTime(advantageRes.UpdateAt),
		Order:     int64(advantageRes.Order),
		Image:     advantageRes.Image,
	}
	return advantages.NewPutAdvantagePhotoOK().WithPayload(&res)
}

func (h *handlers) Delete(params advantages.DeleteAdvantagePhotoIDParams, input interface{}) middleware.Responder {
	id, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		return advantages.NewDeleteAdvantagePhotoIDBadRequest()
	}
	err = h.services.Advantages.Delete(context.Background(), id)
	if err != nil {
		return advantages.NewDeleteAdvantagePhotoIDInternalServerError()
	}
	return advantages.NewDeleteAdvantagePhotoIDOK()
}

func (h *handlers) ChangeAdvantagePhotosOrders(params advantages.PostAdvantagePhotoOrdersParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	first, err := domain.NewUpdateOrderStruct(params.Params.First.ID, int(params.Params.First.Order))
	if err != nil {
		return advantages.NewPostAdvantagePhotoOrdersBadRequest()
	}
	second, err := domain.NewUpdateOrderStruct(params.Params.Second.ID, int(params.Params.Second.Order))
	if err != nil {
		return advantages.NewPostAdvantagePhotoOrdersBadRequest()
	}
	err = h.services.AdvantagePhoto.UpdateOrder(ctx, first, second)
	if err != nil {
		return advantages.NewPostAdvantagePhotoOrdersInternalServerError()
	}
	return advantages.NewPostAdvantagePhotoOrdersOK()
}
