package landing_content

import (
	"context"
	"landing_admin_backend/internal/domain"
	"landing_admin_backend/internal/services/advantages"
	"landing_admin_backend/internal/services/docs"
	"landing_admin_backend/internal/services/editable"
	"landing_admin_backend/internal/services/hod_photos"
	"landing_admin_backend/internal/services/months"
	"landing_admin_backend/internal/services/plans"
	"landing_admin_backend/internal/services/posters"
	"landing_admin_backend/internal/services/video"
	"landing_admin_backend/internal/services/years"
)

type LandingContent interface {
	Get(ctx context.Context) (content domain.LandingContent, err error)
}

type landingContent struct {
	advantages advantages.Advantages
	docs       docs.Docs
	editables  editable.Editable
	years      years.Years
	months     months.Months
	hodPhotos  hod_photos.Photos
	plans      plans.Plans
	posters    posters.Posters
	//projectInfo projectInfo.ProjectInfo
	video video.Video
}

func NewLandingContent(advantages advantages.Advantages, docs docs.Docs, editables editable.Editable, years years.Years, months months.Months, hodPhotos hod_photos.Photos, plans plans.Plans, posters posters.Posters, video video.Video) LandingContent {
	return &landingContent{
		advantages: advantages,
		docs:       docs,
		editables:  editables,
		years:      years,
		months:     months,
		hodPhotos:  hodPhotos,
		plans:      plans,
		posters:    posters,
		video:      video,
	}
}

func (lc *landingContent) Get(ctx context.Context) (content domain.LandingContent, err error) {
	advantages, err := lc.advantages.Get(ctx)
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
	/*months, err := lc.months.Get(ctx, years[0].)
	if err != nil {
		return
	}*/
	plans, err := lc.plans.GetPlans(ctx)
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

	return domain.LandingContent{
		Advantages: advantages,
		Docs:       docs,
		Editables:  editables,
		Hod:        years,
		Plans:      plans,
		Posters:    posters,
		Video:      video,
	}, err
}
