package leads

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations/leads"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"landing_admin_backend/pkg/helpers"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handlers interface {
	Get(params leads.GetLeadParams, input interface{}) middleware.Responder
	Create(params leads.PutLeadParams) middleware.Responder
	Delete(params leads.DeleteLeadIDParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params leads.GetLeadParams, input interface{}) middleware.Responder {
	leadList, err := h.services.Leads.Get(context.Background())
	leadListResponse := []*models.Lead{}
	for _, ld := range leadList {
		leadIns := models.Lead{
			ID:          ld.ID.Hex(),
			Name:        ld.Name,
			Phone:       ld.Phone,
			Email:       ld.Email,
			Roistat:     ld.Roistat,
			UtmSource:   ld.Utm_source,
			UtmMedium:   ld.Utm_medium,
			UtmTerm:     ld.Utm_term,
			UtmCampaign: ld.Utm_campaign,
			UtmContent:  ld.Utm_content,
			SendedToCrm: ld.SendedToCrm,
			UpdatedAt:   strfmt.DateTime(ld.UpdatedAt),
			CreatedAt:   strfmt.DateTime(ld.CreatedAt),
		}
		leadListResponse = append(leadListResponse, &leadIns)
	}
	if err != nil {
		return leads.NewGetLeadBadRequest()
	}
	return leads.NewGetLeadOK().WithPayload(leadListResponse)
}

func (h *handlers) Create(params leads.PutLeadParams) middleware.Responder {
	userID, err := helpers.GetUserIdFromHandler(params)
	if err != nil {
		return leads.NewPutLeadBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}
	lead, err := domain.NewLead(params.Params.ID, params.Params.Name, params.Params.Phone, params.Params.Email, params.Params.Text, params.Params.Roistat, params.Params.UtmSource, params.Params.UtmMedium, params.Params.UtmTerm, params.Params.UtmCampaign, params.Params.UtmContent, time.Now(), time.Now(), userID)
	if err != nil {
		return leads.NewPutLeadBadRequest().WithPayload(&models.ResultResponse{Msg: "check request params"})
	}
	err = h.services.Leads.Create(context.Background(), lead)
	if err != nil {
		return leads.NewPutLeadInternalServerError()
	}
	resLead := models.Lead{
		ID:          lead.ID.Hex(),
		Name:        lead.Name,
		Phone:       lead.Phone,
		Email:       lead.Email,
		Roistat:     lead.Roistat,
		UtmCampaign: lead.Utm_campaign,
		UtmSource:   lead.Utm_source,
		UtmMedium:   lead.Utm_medium,
		UtmContent:  lead.Utm_content,
		UtmTerm:     lead.Utm_term,
		SendedToCrm: lead.SendedToCrm,
		CreatedAt:   strfmt.DateTime(lead.CreatedAt),
		UpdatedAt:   strfmt.DateTime(lead.UpdatedAt),
	}
	return leads.NewPutLeadOK().WithPayload(&resLead)
}

func (h *handlers) Delete(params leads.DeleteLeadIDParams, input interface{}) middleware.Responder {
	id, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		return leads.NewDeleteLeadIDBadRequest()
	}
	err = h.services.Leads.Delete(context.Background(), id)
	if err != nil {
		return leads.NewDeleteLeadIDInternalServerError()
	}
	return leads.NewDeleteLeadIDOK()
}
