package handlers

import (
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/handlers/advantage_photo"
	"landing_admin_backend/internal/handlers/advantages"
	"landing_admin_backend/internal/handlers/auths"
	"landing_admin_backend/internal/handlers/content"
	"landing_admin_backend/internal/handlers/docs"
	"landing_admin_backend/internal/handlers/editable"
	"landing_admin_backend/internal/handlers/hod_photos"
	"landing_admin_backend/internal/handlers/leads"
	"landing_admin_backend/internal/handlers/months"
	"landing_admin_backend/internal/handlers/plans"
	"landing_admin_backend/internal/handlers/posters"
	"landing_admin_backend/internal/handlers/users"
	video "landing_admin_backend/internal/handlers/videos"
	"landing_admin_backend/internal/handlers/years"
	"landing_admin_backend/internal/services"
)

type Handlers struct {
	Users           users.Handlers
	Auth            auths.Handlers
	Poster          posters.Handlers
	Years           years.Handlers
	Months          months.Handlers
	HodPhotos       hod_photos.Handlers
	Advantages      advantages.Handlers
	AdvantagesPhoto advantage_photo.Handlers
	Plans           plans.Handlers
	Docs            docs.Handlers
	Video           video.Handlers
	Leads           leads.Handlers
	Editable        editable.Handlers
	Content         content.Handlers
}

func NewHandlers(services *services.Services, cfg *config.Config) *Handlers {
	return &Handlers{
		Users:           users.NewHandlers(services),
		Auth:            auths.NewHandlers(services, cfg),
		Poster:          posters.NewHandlers(services),
		Years:           years.NewHandlers(services),
		Months:          months.NewHandlers(services),
		HodPhotos:       hod_photos.NewHandlers(services),
		Advantages:      advantages.NewHandlers(services),
		AdvantagesPhoto: advantage_photo.NewHandlers(services),
		Plans:           plans.NewHandlers(services),
		Docs:            docs.NewHandlers(services),
		Video:           video.NewHandlers(services),
		Leads:           leads.NewHandlers(services),
		Editable:        editable.NewHandlers(services),
		Content:         content.NewHandlers(services),
	}
}
