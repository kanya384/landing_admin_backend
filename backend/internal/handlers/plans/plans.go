package plans

import (
	"context"
	"fmt"
	"landing_admin_backend/internal/generated/operations/plans"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"landing_admin_backend/pkg/helpers"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type Handlers interface {
	ProcessPlans(params plans.PutPlansParams) middleware.Responder
	GetPlans(params plans.GetPlansParams, input interface{}) middleware.Responder
	UpdatePlan(params plans.PatchPlansParams, input interface{}) middleware.Responder
	UpdatePlansActivity(params plans.PostPlansParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) GetPlans(params plans.GetPlansParams, input interface{}) middleware.Responder {
	plansList, err := h.services.Plans.GetPlans(context.Background())
	if err != nil {
		fmt.Println(err)
		return plans.NewGetPlansBadRequest()
	}
	resPlans := []*models.Plan{}
	for _, plan := range plansList {
		tPlan := models.Plan{
			ID:                  plan.ID,
			Image:               plan.Image,
			Entrance:            int64(plan.Entrance),
			Commerce:            plan.Commerce,
			Floor:               int64(plan.Floor),
			Number:              int64(plan.Number),
			Rooms:               int64(plan.Rooms),
			Area:                plan.Area,
			LivingArea:          plan.LivingArea,
			KitchenArea:         plan.KitchenArea,
			Price:               int64(plan.Price),
			DiscountPrice:       int64(plan.DiscountPrice),
			SquarePrice:         int64(plan.SquarePrice),
			SquareDiscountPrice: int64(plan.SquareDiscountPrice),
			Discount:            plan.Discount,
			Status:              plan.Status,
			WindowView:          plan.WindowView,
			Liter:               plan.Liter,
			CreatedAt:           strfmt.DateTime(plan.CreatedAt),
		}
		resPlans = append(resPlans, &tPlan)
	}
	return plans.NewGetPlansOK().WithPayload(resPlans)
}

func (h *handlers) ProcessPlans(params plans.PutPlansParams) middleware.Responder {
	records, err := helpers.ReadCsvFile(params.File)
	if err != nil {
		return plans.NewPutPlansBadRequest()
	}
	err = h.services.Plans.ProcessPlans(context.Background(), records)
	if err != nil {
		return plans.NewPutPlansBadRequest()
	}
	return plans.NewPutPlansOK()
}

func (h *handlers) UpdatePlan(params plans.PatchPlansParams, input interface{}) middleware.Responder {
	plan, err := h.services.Plans.UpdatePlansPhoto(context.Background(), params.File, params.ID)
	if err != nil {
		return plans.NewPatchPlansBadRequest()
	}
	tPlan := models.Plan{
		ID:                  plan.ID,
		Image:               plan.Image,
		Entrance:            int64(plan.Entrance),
		Commerce:            plan.Commerce,
		Floor:               int64(plan.Floor),
		Number:              int64(plan.Number),
		Rooms:               int64(plan.Rooms),
		Area:                plan.Area,
		LivingArea:          plan.LivingArea,
		KitchenArea:         plan.KitchenArea,
		Price:               int64(plan.Price),
		DiscountPrice:       int64(plan.DiscountPrice),
		SquarePrice:         int64(plan.SquarePrice),
		SquareDiscountPrice: int64(plan.SquareDiscountPrice),
		Discount:            plan.Discount,
		Status:              plan.Status,
		WindowView:          plan.WindowView,
		Liter:               plan.Liter,
		CreatedAt:           strfmt.DateTime(plan.CreatedAt),
	}
	return plans.NewPatchPlansOK().WithPayload(&tPlan)
}

func (h *handlers) UpdatePlansActivity(params plans.PostPlansParams, input interface{}) middleware.Responder {
	err := h.services.Plans.UpdatePlansActivity(context.Background(), params.Params.ID, params.Params.Status)
	if err != nil {
		return plans.NewPutPlansBadRequest()
	}
	return plans.NewPostPlansOK()
}
