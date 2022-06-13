package services

import (
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/repository"
	"landing_admin_backend/internal/services/advantage_photo"
	"landing_admin_backend/internal/services/advantages"
	"landing_admin_backend/internal/services/auth"
	"landing_admin_backend/internal/services/docs"
	"landing_admin_backend/internal/services/editable"
	"landing_admin_backend/internal/services/hod_photos"
	"landing_admin_backend/internal/services/landing_content"
	"landing_admin_backend/internal/services/leads"
	"landing_admin_backend/internal/services/months"
	"landing_admin_backend/internal/services/plans"
	"landing_admin_backend/internal/services/posters"
	"landing_admin_backend/internal/services/users"
	"landing_admin_backend/internal/services/video"
	"landing_admin_backend/internal/services/years"
	"landing_admin_backend/pkg/memcache"
	"landing_admin_backend/pkg/token_manager"

	"github.com/sirupsen/logrus"
)

type Services struct {
	Auth           auth.Auth
	Users          users.Users
	Posters        posters.Posters
	Years          years.Years
	Months         months.Months
	HodPhotos      hod_photos.Photos
	Advantages     advantages.Advantages
	AdvantagePhoto advantage_photo.AdvantagePhoto
	Plans          plans.Plans
	Video          video.Video
	Docs           docs.Docs
	Leads          leads.Leads
	Editable       editable.Editable
	LandingContent landing_content.LandingContent
}

func Setup(cfg *config.Config, repository *repository.Repository, logger *logrus.Entry, cache memcache.Cache) *Services {
	tokenManager := token_manager.NewTokenManager(cfg.TokenSecret, cfg.TokenTTL, cfg.RefreshTokenTTL)
	advantages := advantages.NewAdvantagesWithCache(advantages.NewAdvantagesWithLogrus(advantages.NewService(repository, cfg), logger), cache)
	docs := docs.NewDocsWithCache(docs.NewDocsWithLogrus(docs.NewService(repository, cfg), logger), cache)
	editables := editable.NewEditableWithCache(editable.NewEditableWithLogrus(editable.NewService(repository, cfg), logger), cache)
	video := video.NewVideoWithCache(video.NewVideoWithLogrus(video.NewService(repository, cfg), logger), cache)
	plans := plans.NewPlansWithCache(plans.NewPlansWithLogrus(plans.NewService(repository, cfg), logger), cache)
	months := months.NewMonthsWithLogrus(months.NewService(repository, cfg), logger)
	years := years.NewYearsWithLogrus(years.NewService(repository, cfg), logger)
	hodPhotos := hod_photos.NewPhotosWithLogrus(hod_photos.NewService(repository, cfg), logger)
	posters := posters.NewPostersWithCache(posters.NewPostersWithLogrus(posters.NewService(repository, cfg), logger), cache)
	return &Services{
		Auth:           auth.NewAuthWithLogrus(auth.NewService(repository, tokenManager), logger),
		Users:          users.NewUsersWithLogrus(users.NewService(repository), logger),
		Posters:        posters,
		Years:          years,
		Months:         months,
		HodPhotos:      hodPhotos,
		Advantages:     advantages,
		AdvantagePhoto: advantage_photo.NewAdvantagePhotoWithLogrus(advantage_photo.NewService(repository, cfg), logger),
		Plans:          plans,
		Video:          video,
		Docs:           docs,
		Leads:          leads.NewLeadsWithLogrus(leads.NewService(repository, cfg), logger),
		Editable:       editables,
		LandingContent: landing_content.NewLandingContent(advantages, docs, editables, years, months, hodPhotos, plans, posters, video),
	}
}
