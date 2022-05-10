package docs

import (
	"context"
	"fmt"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations/docs"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"landing_admin_backend/pkg/helpers"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handlers interface {
	Get(params docs.GetDocParams, input interface{}) middleware.Responder
	Create(params docs.PutDocParams, input interface{}) middleware.Responder
	Update(params docs.PatchDocParams, input interface{}) middleware.Responder
	Delete(params docs.DeleteDocIDParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params docs.GetDocParams, input interface{}) middleware.Responder {
	fmt.Println("get")
	docsList, err := h.services.Docs.Get(context.Background())
	if err != nil {
		return docs.NewGetDocsBadRequest()
	}
	docsListResponse := []*models.Doc{}
	for _, doc := range docsList {
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
	if err != nil {
		return docs.NewGetDocsBadRequest()
	}
	return docs.NewGetDocsOK().WithPayload(docsListResponse)
}

func (h *handlers) Create(params docs.PutDocParams, input interface{}) middleware.Responder {
	fmt.Println("insert")
	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return docs.NewPutDocsBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}
	doc, err := domain.NewDoc("", params.Title, "", time.Now(), time.Now(), userID)
	if err != nil {
		return docs.NewPutDocsBadRequest().WithPayload(&models.ResultResponse{Msg: "check request params"})
	}
	docInserted, err := h.services.Docs.Create(context.Background(), doc, params.File, "file.pdf")
	if err != nil {
		return docs.NewPutDocsInternalServerError()
	}
	docRes := models.Doc{
		ID:         docInserted.ID.Hex(),
		Title:      docInserted.Title,
		File:       docInserted.File,
		CreatedAt:  strfmt.DateTime(docInserted.CreatedAt),
		UpdatedAt:  strfmt.DateTime(docInserted.UpdateAt),
		ModifiedBy: docInserted.ModifiedBy.Hex(),
	}
	return docs.NewPutDocsOK().WithPayload(&docRes)
}

func (h *handlers) Update(params docs.PatchDocParams, input interface{}) middleware.Responder {
	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return docs.NewPutDocsBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}
	doc, err := domain.NewDoc("", params.Title, "", time.Now(), time.Now(), userID)
	if err != nil {
		return docs.NewPutDocsBadRequest().WithPayload(&models.ResultResponse{Msg: "check request params"})
	}
	docInserted, err := h.services.Docs.Create(context.Background(), doc, params.File, "file.pdf")
	if err != nil {
		return docs.NewPutDocsInternalServerError()
	}
	docRes := models.Doc{
		ID:         docInserted.ID.Hex(),
		Title:      docInserted.Title,
		File:       docInserted.File,
		CreatedAt:  strfmt.DateTime(docInserted.CreatedAt),
		UpdatedAt:  strfmt.DateTime(docInserted.UpdateAt),
		ModifiedBy: docInserted.ModifiedBy.Hex(),
	}
	return docs.NewPatchDocsOK().WithPayload(&docRes)
}

func (h *handlers) Delete(params docs.DeleteDocIDParams, input interface{}) middleware.Responder {
	id, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		return docs.NewDeleteDocsIDBadRequest()
	}
	err = h.services.Docs.Delete(context.Background(), id)
	if err != nil {
		return docs.NewDeleteDocsIDInternalServerError()
	}
	return docs.NewDeleteDocsIDOK()
}
