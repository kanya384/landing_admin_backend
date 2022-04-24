package hod_photos

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations/hod"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"landing_admin_backend/pkg/helpers"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handlers interface {
	Get(params hod.GetPhotosIDParams, input interface{}) middleware.Responder
	Create(params hod.PutPhotosParams, input interface{}) middleware.Responder
	ChangePhotosOrders(params hod.PostPhotosOrdersParams, input interface{}) middleware.Responder
	Delete(params hod.DeletePhotosIDParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params hod.GetPhotosIDParams, input interface{}) middleware.Responder {
	yearID, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		return hod.NewGetPhotosIDBadRequest()
	}
	photosList, err := h.services.HodPhotos.Get(context.Background(), yearID)
	if err != nil {
		return hod.NewGetPhotosIDBadRequest()
	}
	photosListResponse := []*models.HodPhoto{}
	for _, photo := range photosList {
		monthIns := models.HodPhoto{
			ID:      photo.ID.Hex(),
			Image:   photo.Image,
			MonthID: photo.MonthID.Hex(),
			Order:   int64(photo.Order),
		}
		photosListResponse = append(photosListResponse, &monthIns)
	}
	if err != nil {
		return hod.NewGetPhotosIDBadRequest()
	}
	if len(photosListResponse) > 0 {
		return hod.NewGetPhotosIDOK().WithPayload(photosListResponse)
	}
	return hod.NewGetPhotosIDOK()
}

func (h *handlers) Create(params hod.PutPhotosParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return hod.NewPutPhotosBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}

	photo, err := domain.NewHodPhoto("", params.MonthID, "", 0, time.Now(), time.Now(), userID)
	if err != nil {
		return hod.NewPutMonthsBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating poster"})
	}

	photoRes, err := h.services.HodPhotos.Create(ctx, photo, params.File)
	if err != nil {
		return hod.NewPutPhotosBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating poster"})
	}
	resPhoto := models.HodPhoto{
		ID:      photoRes.ID.Hex(),
		Image:   photoRes.Image,
		MonthID: photoRes.MonthID.Hex(),
		Order:   int64(photoRes.Order),
	}
	return hod.NewPutPhotosOK().WithPayload(&resPhoto)
}

func (h *handlers) ChangePhotosOrders(params hod.PostPhotosOrdersParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	first, err := domain.NewUpdateOrderStruct(params.Params.First.ID, int(params.Params.First.Order))
	if err != nil {
		return hod.NewPostPhotosOrdersBadRequest()
	}
	second, err := domain.NewUpdateOrderStruct(params.Params.Second.ID, int(params.Params.Second.Order))
	if err != nil {
		return hod.NewPostPhotosOrdersBadRequest()
	}
	err = h.services.HodPhotos.UpdateOrder(ctx, first, second)
	if err != nil {
		return hod.NewPostPhotosOrdersInternalServerError()
	}
	return hod.NewPostPhotosOrdersOK()
}

func (h *handlers) Delete(params hod.DeletePhotosIDParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	photoID, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		return hod.NewDeletePhotosIDBadRequest()
	}
	err = h.services.HodPhotos.Delete(ctx, photoID)
	if err != nil {
		return hod.NewDeletePhotosIDBadRequest()
	}
	return hod.NewDeletePhotosIDOK().WithPayload(&models.ResultResponse{Msg: "month deleted"})
}
