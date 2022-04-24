package months

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
	Get(params hod.GetMonthsIDParams, input interface{}) middleware.Responder
	Create(params hod.PutMonthsParams, input interface{}) middleware.Responder
	Update(params hod.PatchMonthsParams, input interface{}) middleware.Responder
	Delete(params hod.DeleteMonthsIDParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params hod.GetMonthsIDParams, input interface{}) middleware.Responder {
	yearID, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		return hod.NewGetMonthsIDBadRequest()
	}
	monthsList, err := h.services.Months.Get(context.Background(), yearID)
	if err != nil {
		return hod.NewGetMonthsIDBadRequest()
	}
	monthsListResponse := []*models.Month{}
	for _, mnth := range monthsList {
		monthIns := models.Month{
			ID:     mnth.ID.Hex(),
			Name:   mnth.Name,
			Value:  int64(mnth.Value),
			YearID: mnth.YearID.Hex(),
		}
		monthsListResponse = append(monthsListResponse, &monthIns)
	}
	if err != nil {
		return hod.NewGetMonthsIDBadRequest()
	}
	if len(monthsListResponse) > 0 {
		return hod.NewGetMonthsIDOK().WithPayload(monthsListResponse)
	}
	return hod.NewGetMonthsIDOK()
}

func (h *handlers) Create(params hod.PutMonthsParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return hod.NewPutMonthsBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}

	month, err := domain.NewMonth("", params.Params.YearID, int(params.Params.Value), params.Params.Name, time.Now(), time.Now(), userID)
	if err != nil {
		return hod.NewPutMonthsBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating poster"})
	}

	err = h.services.Months.Create(ctx, month)
	if err != nil {
		return hod.NewPutMonthsBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating poster"})
	}
	resMonth := models.Month{
		ID:     month.ID.Hex(),
		Name:   month.Name,
		Value:  int64(month.Value),
		YearID: month.YearID.Hex(),
	}
	return hod.NewPutMonthsOK().WithPayload(&resMonth)
}

func (h *handlers) Update(params hod.PatchMonthsParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return hod.NewPatchMonthsBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}

	month, err := domain.NewMonth(params.Params.ID, params.Params.YearID, int(params.Params.Value), params.Params.Name, time.Now(), time.Now(), userID)
	if err != nil {
		return hod.NewPatchMonthsBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating poster"})
	}
	err = h.services.Months.Update(ctx, month)
	if err != nil {
		return hod.NewPatchMonthsBadRequest().WithPayload(&models.ResultResponse{Msg: "error updating poster"})
	}
	resMonth := models.Month{
		ID:     month.ID.Hex(),
		Name:   month.Name,
		Value:  int64(month.Value),
		YearID: month.YearID.Hex(),
	}
	return hod.NewPatchMonthsOK().WithPayload(&resMonth)
}

func (h *handlers) Delete(params hod.DeleteMonthsIDParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	monthID, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		return hod.NewDeleteMonthsIDBadRequest()
	}
	err = h.services.Months.Delete(ctx, monthID)
	if err != nil {
		return hod.NewDeleteMonthsIDBadRequest()
	}
	return hod.NewDeleteMonthsIDOK().WithPayload(&models.ResultResponse{Msg: "month deleted"})
}
