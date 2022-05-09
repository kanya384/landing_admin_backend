package advantages

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
	Get(params advantages.GetAdvantagesParams, input interface{}) middleware.Responder
	GetByID(params advantages.GetAdvantagesIDParams, input interface{}) middleware.Responder
	Create(params advantages.PutAdvantagesParams, input interface{}) middleware.Responder
	Update(params advantages.PatchAdvantagesParams, input interface{}) middleware.Responder
	Delete(params advantages.DeleteAdvantagesIDParams, input interface{}) middleware.Responder
	ChangeAdvantageOrders(params advantages.PostAdvantagesOrdersParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params advantages.GetAdvantagesParams, input interface{}) middleware.Responder {
	advantageList, err := h.services.Advantages.Get(context.Background())
	advantageListResponse := []*models.Advantage{}
	for _, pst := range advantageList {
		postIns := models.Advantage{
			ID:          pst.ID.Hex(),
			Title:       pst.Title,
			Description: pst.Description,
			Order:       int64(pst.Order),
			UpdatedAt:   strfmt.DateTime(pst.UpdateAt),
			CreatedAt:   strfmt.DateTime(pst.CreatedAt),
		}
		advantageListResponse = append(advantageListResponse, &postIns)
	}
	if err != nil {
		return advantages.NewGetAdvantagesBadRequest()
	}
	return advantages.NewGetAdvantagesOK().WithPayload(advantageListResponse)
}

func (h *handlers) GetByID(params advantages.GetAdvantagesIDParams, input interface{}) middleware.Responder {
	id, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		return advantages.NewGetAdvantagesIDBadRequest()
	}
	advantage, err := h.services.Advantages.GetByID(context.Background(), id)
	if err != nil {
		return advantages.NewGetAdvantagesIDBadRequest()
	}
	advantageRes := models.Advantage{
		ID:          advantage.ID.Hex(),
		Title:       advantage.Title,
		Description: advantage.Description,
		Order:       int64(advantage.Order),
		CreatedAt:   strfmt.DateTime(advantage.CreatedAt),
		UpdatedAt:   strfmt.DateTime(advantage.UpdateAt),
	}
	return advantages.NewGetAdvantagesIDOK().WithPayload(&advantageRes)
}

func (h *handlers) Create(params advantages.PutAdvantagesParams, input interface{}) middleware.Responder {
	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return advantages.NewPutAdvantagesBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}
	advantage, err := domain.NewAdvantage(params.Params.ID, params.Params.Title, params.Params.Description, int(params.Params.Order), time.Now(), time.Now(), userID)
	if err != nil {
		return advantages.NewPutAdvantagesBadRequest().WithPayload(&models.ResultResponse{Msg: "check request params"})
	}
	err = h.services.Advantages.Create(context.Background(), advantage)
	if err != nil {
		return advantages.NewPutAdvantagesInternalServerError()
	}
	resAdvantage := models.Advantage{
		ID:          advantage.ID.Hex(),
		Title:       advantage.Title,
		Description: advantage.Description,
		CreatedAt:   strfmt.DateTime(advantage.CreatedAt),
		UpdatedAt:   strfmt.DateTime(advantage.UpdateAt),
		Order:       int64(advantage.Order),
	}
	return advantages.NewPutAdvantagesOK().WithPayload(&resAdvantage)
}

func (h *handlers) Update(params advantages.PatchAdvantagesParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return advantages.NewPatchAdvantagesBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}
	advantage, err := domain.NewAdvantage(params.Params.ID, params.Params.Title, params.Params.Description, int(params.Params.Order), time.Now(), time.Now(), userID)
	if err != nil {
		return advantages.NewPatchAdvantagesBadRequest().WithPayload(&models.ResultResponse{Msg: "error updating poster"})
	}
	err = h.services.Advantages.Update(ctx, advantage)
	if err != nil {
		return advantages.NewPatchAdvantagesInternalServerError()
	}
	resAdvantage := models.Advantage{
		ID:          advantage.ID.Hex(),
		Title:       advantage.Title,
		Description: advantage.Description,
		CreatedAt:   strfmt.DateTime(advantage.CreatedAt),
		UpdatedAt:   strfmt.DateTime(advantage.UpdateAt),
		Order:       int64(advantage.Order),
	}

	return advantages.NewPatchAdvantagesOK().WithPayload(&resAdvantage)
}

func (h *handlers) Delete(params advantages.DeleteAdvantagesIDParams, input interface{}) middleware.Responder {
	id, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		return advantages.NewDeleteAdvantagesIDBadRequest()
	}
	err = h.services.Advantages.Delete(context.Background(), id)
	if err != nil {
		return advantages.NewDeleteAdvantagesIDInternalServerError()
	}
	return advantages.NewDeleteAdvantagesIDOK()
}

func (h *handlers) ChangeAdvantageOrders(params advantages.PostAdvantagesOrdersParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	first, err := domain.NewUpdateOrderStruct(params.Params.First.ID, int(params.Params.First.Order))
	if err != nil {
		return advantages.NewPostAdvantagesOrdersBadRequest()
	}
	second, err := domain.NewUpdateOrderStruct(params.Params.Second.ID, int(params.Params.Second.Order))
	if err != nil {
		return advantages.NewPostAdvantagesOrdersBadRequest()
	}
	err = h.services.Advantages.UpdateOrder(ctx, first, second)
	if err != nil {
		return advantages.NewPostAdvantagesOrdersInternalServerError()
	}
	return advantages.NewPostAdvantagesOrdersOK()
}
