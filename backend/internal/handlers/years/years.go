package years

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
	"go.mongodb.org/mongo-driver/mongo"
)

type Handlers interface {
	Get(params hod.GetYearsParams, input interface{}) middleware.Responder
	Create(params hod.PutYearsParams, input interface{}) middleware.Responder
	Update(params hod.PatchYearsParams, input interface{}) middleware.Responder
	Delete(params hod.DeleteYearsIDParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params hod.GetYearsParams, input interface{}) middleware.Responder {
	yearsList, errD := h.services.Years.Get(context.Background())
	if errD != nil && errD != mongo.ErrNoDocuments {
		return hod.NewGetMonthsIDBadRequest()
	}
	yearsListResponse := models.GetYearsResponse{}
	for _, yr := range yearsList {
		yearIns := models.Year{
			ID:    yr.ID.Hex(),
			Value: int64(yr.Value),
		}
		yearsListResponse = append(yearsListResponse, &yearIns)
	}
	return hod.NewGetYearsOK().WithPayload(yearsListResponse)
}

func (h *handlers) Create(params hod.PutYearsParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return hod.NewPutYearsBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}

	year, err := domain.NewYear("", int(params.Params.Value), time.Now(), time.Now(), userID)
	if err != nil {
		return hod.NewPutYearsBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating year"})
	}

	err = h.services.Years.Create(ctx, year)
	if err != nil {
		return hod.NewPutYearsBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating year"})
	}
	resYear := models.Year{
		ID:    year.ID.Hex(),
		Value: int64(year.Value),
	}
	return hod.NewPutYearsOK().WithPayload(&resYear)
}

func (h *handlers) Update(params hod.PatchYearsParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return hod.NewPatchYearsBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}

	year, err := domain.NewYear(params.Params.ID, int(params.Params.Value), time.Now(), time.Now(), userID)
	if err != nil {
		return hod.NewPatchYearsBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating year"})
	}
	err = h.services.Years.Update(ctx, year)
	if err != nil {
		return hod.NewPatchYearsBadRequest().WithPayload(&models.ResultResponse{Msg: "error updating year"})
	}
	resYear := models.Year{
		ID:    year.ID.Hex(),
		Value: int64(year.Value),
	}
	return hod.NewPatchYearsOK().WithPayload(&resYear)
}

func (h *handlers) Delete(params hod.DeleteYearsIDParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	yearID, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		return hod.NewDeleteYearsIDBadRequest()
	}
	err = h.services.Years.Delete(ctx, yearID)
	if err != nil {
		return hod.NewDeleteYearsIDBadRequest()
	}
	return hod.NewDeleteYearsIDOK().WithPayload(&models.ResultResponse{Msg: "year deleted"})
}
