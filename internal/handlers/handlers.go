package handlers

import (
	"landing_admin_backend/internal/handlers/auths"
	"landing_admin_backend/internal/handlers/users"
	"landing_admin_backend/internal/services"
)

type Handlers struct {
	Users users.Handlers
	Auth  auths.Handlers
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		Users: users.NewHandlers(services),
		Auth:  auths.NewHandlers(services),
	}
}
