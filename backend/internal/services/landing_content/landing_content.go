package landing_content

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/services/advantage_photo"
	"landing_admin_backend/internal/services/advantages"
	"landing_admin_backend/internal/services/docs"
	"landing_admin_backend/internal/services/editable"
	"landing_admin_backend/internal/services/hod_photos"
	"landing_admin_backend/internal/services/months"
	"landing_admin_backend/internal/services/plans"
	"landing_admin_backend/internal/services/posters"
	"landing_admin_backend/internal/services/project_info"
	"landing_admin_backend/internal/services/settings"
	"landing_admin_backend/internal/services/titles"
	"landing_admin_backend/internal/services/video"
	"landing_admin_backend/internal/services/years"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LandingContent interface {
	Get(ctx context.Context) (content domain.LandingContent, err error)
}

type landingContent struct {
	advantages      advantages.Advantages
	advantagePhotos advantage_photo.AdvantagePhoto
	docs            docs.Docs
	editables       editable.Editable
	years           years.Years
	months          months.Months
	hodPhotos       hod_photos.Photos
	plans           plans.Plans
	posters         posters.Posters
	projectInfo     project_info.ProjectInfo
	video           video.Video
	title           titles.Title
	settings        settings.Settings
}

func NewLandingContent(advantages advantages.Advantages, docs docs.Docs, editables editable.Editable, years years.Years, months months.Months, hodPhotos hod_photos.Photos, plans plans.Plans, posters posters.Posters, video video.Video, projectInfo project_info.ProjectInfo, advantagePhotos advantage_photo.AdvantagePhoto, title titles.Title, settings settings.Settings) LandingContent {
	return &landingContent{
		advantages:      advantages,
		docs:            docs,
		editables:       editables,
		years:           years,
		months:          months,
		hodPhotos:       hodPhotos,
		plans:           plans,
		posters:         posters,
		projectInfo:     projectInfo,
		video:           video,
		advantagePhotos: advantagePhotos,
		title:           title,
		settings:        settings,
	}
}

func (lc *landingContent) Get(ctx context.Context) (content domain.LandingContent, err error) {
	advantages, err := lc.advantages.Get(ctx)
	if err != nil {
		return
	}
	advantagePhoto, err := lc.advantagePhotos.Get(ctx, primitive.NilObjectID)
	if err != nil {
		return
	}
	docs, err := lc.docs.Get(ctx)
	if err != nil {
		return
	}
	editables, err := lc.editables.Get(ctx)
	if err != nil {
		return
	}
	years, err := lc.years.Get(ctx)
	if err != nil {
		return
	}
	months, err := lc.months.Get(ctx, primitive.NilObjectID)
	if err != nil {
		return
	}
	photos, err := lc.hodPhotos.Get(ctx, primitive.NilObjectID)
	if err != nil {
		return
	}
	plans, err := lc.plans.GetActivePlans(ctx)
	if err != nil {
		return
	}

	posters, err := lc.posters.Get(ctx, map[string]interface{}{"active": true})
	if err != nil {
		return
	}
	video, err := lc.video.Get(ctx)
	if err != nil {
		return
	}
	projectInfo, err := lc.projectInfo.Get(ctx, map[string]interface{}{})
	if err != nil {
		return
	}

	settings, err := lc.settings.Get(ctx)
	if err != nil {
		return
	}

	titles, err := lc.title.Get(ctx)
	if err != nil {
		return
	}

	return domain.LandingContent{
		Advantages:      advantages,
		AdvantagePhotos: advantagePhoto,
		Docs:            docs,
		Editables:       editables,
		Years:           years,
		Months:          months,
		Photos:          photos,
		Plans:           plans,
		Posters:         posters,
		Video:           video,
		ProjectInfo:     projectInfo,
		Setting:         settings,
		Title:           titles,
	}, err
}
