package project_info

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/generated/operations/project_info"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"landing_admin_backend/pkg/helpers"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handlers interface {
	Get(params project_info.GetProjectInfoParams) middleware.Responder
	GetPosterById(params project_info.GetProjectInfoPojectInfoIDParams, input interface{}) middleware.Responder
	Create(params project_info.PutProjectInfoParams, input interface{}) middleware.Responder
	Update(params project_info.PatchProjectInfoParams, input interface{}) middleware.Responder
	Delete(params project_info.DeleteProjectInfoPojectInfoIDParams, input interface{}) middleware.Responder
	ProjectInfoOrdersChange(params project_info.PostProjectInfoOrdersParams, input interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) Handlers {
	return &handlers{
		services: services,
	}
}

func (h *handlers) Get(params project_info.GetProjectInfoParams) middleware.Responder {
	filter := map[string]interface{}{
		"active": "",
		"name":   "",
	}
	projectInfoList, err := h.services.ProjectInfo.Get(context.Background(), filter)
	projectInfoListResponse := models.GetProjectInfosResponse{}
	for _, pst := range projectInfoList {
		postIns := models.ProjectInfo{
			ID:          pst.ID.Hex(),
			Title:       pst.Title,
			Description: pst.Description,
			Anonce:      pst.Anonce,
			Photo:       pst.Photo,
			ModifiedBy:  pst.ModifiedBy.Hex(),
			Order:       int64(pst.Order),
			UpdatedAt:   strfmt.DateTime(pst.UpdateAt),
			CreatedAt:   strfmt.DateTime(pst.CreatedAt),
		}
		projectInfoListResponse = append(projectInfoListResponse, &postIns)
	}
	if err != nil {
		return project_info.NewGetProjectInfoBadRequest()
	}
	return project_info.NewGetProjectInfoOK().WithPayload(projectInfoListResponse)
}

func (h *handlers) GetPosterById(params project_info.GetProjectInfoPojectInfoIDParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	projectInfoID, err := primitive.ObjectIDFromHex(params.PojectInfoID)
	if err != nil {
		return project_info.NewGetProjectInfoPojectInfoIDBadRequest()
	}
	projectInfo, err := h.services.ProjectInfo.GetByID(ctx, projectInfoID)
	if err != nil {
		return project_info.NewGetProjectInfoPojectInfoIDInternalServerError()
	}
	res := models.ProjectInfo{
		ID:          projectInfo.ID.Hex(),
		Title:       projectInfo.Title,
		Description: projectInfo.Description,
		Anonce:      projectInfo.Anonce,
		Photo:       projectInfo.Photo,
		Order:       int64(projectInfo.Order),
		CreatedAt:   strfmt.DateTime(projectInfo.CreatedAt),
		UpdatedAt:   strfmt.DateTime(projectInfo.UpdateAt),
		ModifiedBy:  projectInfo.ModifiedBy.Hex(),
	}
	return project_info.NewGetProjectInfoPojectInfoIDOK().WithPayload(&res)
}

func (h *handlers) Create(params project_info.PutProjectInfoParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return project_info.NewPutProjectInfoBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}

	projectInfo, err := domain.NewProjectInfo("", params.Title, params.Anonce, params.Description, "", 0, time.Now(), time.Now(), userID)
	if err != nil {
		return project_info.NewPutProjectInfoBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating projectInfo"})
	}

	projectInfoRes, err := h.services.ProjectInfo.Create(ctx, projectInfo, params.File)
	if err != nil {
		return project_info.NewPutProjectInfoBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating projectInfo"})
	}
	resProjectInfo := models.ProjectInfo{
		ID:          projectInfoRes.ID.Hex(),
		Title:       projectInfoRes.Title,
		Description: projectInfoRes.Description,
		Anonce:      projectInfoRes.Anonce,
		CreatedAt:   strfmt.DateTime(projectInfoRes.CreatedAt),
		UpdatedAt:   strfmt.DateTime(projectInfoRes.UpdateAt),
		Order:       int64(projectInfoRes.Order),
		Photo:       projectInfoRes.Photo,
		ModifiedBy:  projectInfoRes.ModifiedBy.Hex(),
	}
	return project_info.NewPutProjectInfoOK().WithPayload(&resProjectInfo)
}

func (h *handlers) Update(params project_info.PatchProjectInfoParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	userID, err := helpers.GetUserIdFromHandler(input)
	if err != nil {
		return project_info.NewPatchProjectInfoBadRequest().WithPayload(&models.ResultResponse{Msg: "wrong user"})
	}

	projectInfo, err := domain.NewProjectInfo(params.ID, params.Title, params.Anonce, params.Description, *params.Photo, 0, time.Now(), time.Now(), userID)
	if err != nil {
		return project_info.NewPatchProjectInfoBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating projectInfo"})
	}
	projectInfoRes, err := h.services.ProjectInfo.Update(ctx, projectInfo, params.File)
	if err != nil {
		return project_info.NewPatchProjectInfoBadRequest().WithPayload(&models.ResultResponse{Msg: "error creating projectInfo"})
	}
	resProjectInfo := models.ProjectInfo{
		ID:          projectInfoRes.ID.Hex(),
		Title:       projectInfoRes.Title,
		Description: projectInfoRes.Description,
		Anonce:      projectInfoRes.Anonce,
		CreatedAt:   strfmt.DateTime(projectInfoRes.CreatedAt),
		UpdatedAt:   strfmt.DateTime(projectInfoRes.UpdateAt),
		Order:       int64(projectInfoRes.Order),
		Photo:       projectInfoRes.Photo,
		ModifiedBy:  projectInfoRes.ModifiedBy.Hex(),
	}

	return project_info.NewPatchProjectInfoOK().WithPayload(&resProjectInfo)
}

func (h *handlers) Delete(params project_info.DeleteProjectInfoPojectInfoIDParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	projectInfoID, err := primitive.ObjectIDFromHex(params.PojectInfoID)
	if err != nil {
		return project_info.NewDeleteProjectInfoPojectInfoIDBadRequest()
	}
	err = h.services.ProjectInfo.Delete(ctx, projectInfoID)
	if err != nil {
		return project_info.NewDeleteProjectInfoPojectInfoIDBadRequest()
	}
	return project_info.NewDeleteProjectInfoPojectInfoIDOK().WithPayload(&models.ResultResponse{Msg: "poster deleted"})
}

func (h *handlers) ProjectInfoOrdersChange(params project_info.PostProjectInfoOrdersParams, input interface{}) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	first, err := domain.NewUpdateOrderStruct(params.Params.First.ID, int(params.Params.First.Order))
	if err != nil {
		return project_info.NewPostProjectInfoOrdersBadRequest()
	}
	second, err := domain.NewUpdateOrderStruct(params.Params.Second.ID, int(params.Params.Second.Order))
	if err != nil {
		return project_info.NewPostProjectInfoOrdersBadRequest()
	}
	err = h.services.ProjectInfo.UpdateOrder(ctx, first, second)
	if err != nil {
		return project_info.NewPostProjectInfoOrdersInternalServerError()
	}
	return project_info.NewPostProjectInfoOrdersOK()
}
