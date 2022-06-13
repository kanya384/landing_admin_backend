package content

import (
	"context"
	"landing_admin_backend/internal/generated/operations/content"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type Handlers interface {
	Get(params content.GetContentParams) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params content.GetContentParams) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	contentL, err := h.services.LandingContent.Get(ctx)
	if err != nil {
		return content.NewGetContentBadRequest()
	}

	resContent := models.Content{}

	advantageListResponse := []*models.Advantage{}
	for _, pst := range contentL.Advantages {
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
	resContent.Advantages = advantageListResponse

	docsListResponse := []*models.Doc{}
	for _, doc := range contentL.Docs {
		docIns := models.Doc{
			ID:         doc.ID.Hex(),
			Title:      doc.Title,
			File:       doc.File,
			UpdatedAt:  strfmt.DateTime(doc.UpdateAt),
			CreatedAt:  strfmt.DateTime(doc.CreatedAt),
			ModifiedBy: doc.ModifiedBy.Hex(),
		}
		docsListResponse = append(docsListResponse, &docIns)
	}
	resContent.Docs = docsListResponse

	editablesListResponse := []*models.Editable{}
	for _, editable := range contentL.Editables {
		editableIns := models.Editable{
			ID:     editable.ID.Hex(),
			Type:   int64(editable.Type),
			Value:  editable.Value,
			Height: int64(editable.Height),
			Width:  int64(editable.Width),
		}
		editablesListResponse = append(editablesListResponse, &editableIns)
	}
	resContent.Editables = editablesListResponse

	resPlans := []*models.Plan{}
	for _, plan := range contentL.Plans {
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
	resContent.Plans = resPlans

	videoListResponse := []*models.Video{}
	for _, video := range contentL.Video {
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

	resContent.Video = videoListResponse

	return content.NewGetContentOK().WithPayload(&resContent)
}
