package actions

import (
	"context"
	"fmt"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations/actions"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"landing_admin_backend/pkg/helpers"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handlers interface {
	Get(params actions.GetActionsParams) middleware.Responder
	GetActionById(params actions.GetActionsActionIDParams, input interface{}) middleware.Responder
	Create(params actions.PutActionsParams, input interface{}) middleware.Responder
	Update(params actions.PatchActionsParams, input interface{}) middleware.Responder
	Delete(params actions.DeleteActionsActionIDParams, input interface{}) middleware.Responder
	ActionsOrdersChange(params actions.PostActionsOrdersParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Create(params actions.PutActionsParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return actions.NewPutActionsBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}

	action, err := domain.NewAction("", params.Title, params.Description, params.Date, "", "", 0, time.Now(), time.Now(), userID)
	if err != nil {
		return actions.NewPutActionsBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating action"})
	}

	actionRes, err := h.services.Actions.Create(ctx, action, params.File)
	if err != nil {
		return actions.NewPutActionsBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating action"})
	}
	resAction := models.Action{
		ID:         actionRes.ID.Hex(),
		Title:      actionRes.Title,
		CreatedAt:  strfmt.DateTime(actionRes.CreatedAt),
		UpdatedAt:  strfmt.DateTime(actionRes.UpdateAt),
		Order:      int64(actionRes.Order),
		Photo:      actionRes.Photo,
		ModifiedBy: actionRes.ModifiedBy.Hex(),
	}
	return actions.NewPutActionsOK().WithPayload(&resAction)
}

func (h *handlers) Update(params actions.PatchActionsParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return actions.NewPatchActionsBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}
	photo := ""
	if params.Photo != nil {
		photo = *params.Photo
	}
	preview := ""
	if params.Preview != nil {
		preview = *params.Preview
	}
	action, err := domain.NewAction(params.ID, params.Title, params.Description, params.Date, preview, photo, int(params.Order), time.Now(), time.Now(), userID)
	if err != nil {
		return actions.NewPatchActionsBadRequest().WithPayload(&models.ResultResponse{Msg: "error updating action"})
	}
	action, err = h.services.Actions.Update(ctx, action, params.File)
	if err != nil {
		return actions.NewPatchActionsBadRequest().WithPayload(&models.ResultResponse{Msg: "error updating action"})
	}

	actionRes := &models.Action{
		ID:         action.ID.Hex(),
		Photo:      action.Photo,
		Title:      action.Title,
		Order:      int64(action.Order),
		ModifiedBy: action.ModifiedBy.Hex(),
		CreatedAt:  strfmt.DateTime(action.CreatedAt),
		UpdatedAt:  strfmt.DateTime(action.UpdateAt),
	}

	fmt.Println(actionRes)

	return actions.NewPatchActionsOK().WithPayload(actionRes)
}

func (h *handlers) Delete(params actions.DeleteActionsActionIDParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	actionID, err := primitive.ObjectIDFromHex(params.ActionID)
	if err != nil {
		return actions.NewDeleteActionsActionIDBadRequest()
	}
	err = h.services.Actions.Delete(ctx, actionID)
	if err != nil {
		return actions.NewDeleteActionsActionIDBadRequest()
	}
	return actions.NewDeleteActionsActionIDOK().WithPayload(&models.ResultResponse{Msg: "action deleted"})
}

func (h *handlers) Get(params actions.GetActionsParams) middleware.Responder {
	filter := map[string]interface{}{
		"active": "",
		"name":   "",
	}
	actionsList, err := h.services.Actions.Get(context.Background(), filter)
	actionsListResponse := models.GetActionsResponse{}
	for _, pst := range actionsList {
		postIns := models.Action{
			ID:         pst.ID.Hex(),
			Title:      pst.Title,
			Photo:      pst.Photo,
			ModifiedBy: pst.ModifiedBy.Hex(),
			Order:      int64(pst.Order),
			UpdatedAt:  strfmt.DateTime(pst.UpdateAt),
			CreatedAt:  strfmt.DateTime(pst.CreatedAt),
		}
		actionsListResponse = append(actionsListResponse, &postIns)
	}
	if err != nil {
		return actions.NewGetActionsBadRequest()
	}
	return actions.NewGetActionsOK().WithPayload(actionsListResponse)
}

func (h *handlers) GetActionById(params actions.GetActionsActionIDParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	actionID, err := primitive.ObjectIDFromHex(params.ActionID)
	if err != nil {
		return actions.NewDeleteActionsActionIDBadRequest()
	}
	action, err := h.services.Actions.GetByID(ctx, actionID)
	if err != nil {
		return actions.NewDeleteActionsActionIDBadRequest()
	}
	res := models.Action{
		ID:         action.ID.Hex(),
		Title:      action.Title,
		Photo:      action.Photo,
		Order:      int64(action.Order),
		CreatedAt:  strfmt.DateTime(action.CreatedAt),
		UpdatedAt:  strfmt.DateTime(action.UpdateAt),
		ModifiedBy: action.ModifiedBy.Hex(),
	}
	return actions.NewGetActionsActionIDOK().WithPayload(&res)
}

func (h *handlers) ActionsOrdersChange(params actions.PostActionsOrdersParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	first, err := domain.NewUpdateOrderStruct(params.Params.First.ID, int(params.Params.First.Order))
	if err != nil {
		return actions.NewPostActionsOrdersBadRequest()
	}
	second, err := domain.NewUpdateOrderStruct(params.Params.Second.ID, int(params.Params.Second.Order))
	if err != nil {
		return actions.NewPostActionsOrdersBadRequest()
	}
	err = h.services.Actions.UpdateOrder(ctx, first, second)
	if err != nil {
		return actions.NewPostActionsOrdersInternalServerError()
	}
	return actions.NewPostActionsOrdersOK()
}
