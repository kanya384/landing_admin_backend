package handlers

import (
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/handlers/auths"
	"landing_admin_backend/internal/handlers/posters"
	"landing_admin_backend/internal/handlers/users"
	"landing_admin_backend/internal/services"
)

type Handlers struct {
	Users  users.Handlers
	Auth   auths.Handlers
	Poster posters.Handlers
}

func NewHandlers(services *services.Services, cfg *config.Config) *Handlers {
	return &Handlers{
		Users:  users.NewHandlers(services),
		Auth:   auths.NewHandlers(services, cfg),
		Poster: posters.NewHandlers(services),
	}
}
