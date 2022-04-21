package auths

import (
	"context"
	"encoding/json"
	"landing_admin_backend/internal/config"
	"landing_admin_backend/internal/generated/operations"
	"landing_admin_backend/internal/services"
	"landing_admin_backend/models"
	"landing_admin_backend/pkg/helpers"
	"net/http"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

type Handlers interface {
	Authenticate(params operations.PostLoginParams) middleware.Responder
	Ping(params operations.GetPingParams, claims interface{}) middleware.Responder
}

type handlers struct {
	services *services.Services
	cfg      *config.Config
}

func NewHandlers(services *services.Services, cfg *config.Config) Handlers {
	return &handlers{
		services: services,
		cfg:      cfg,
	}
}

func (h *handlers) Authenticate(params operations.PostLoginParams) middleware.Responder {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	tokens, err := h.services.Auth.Authenticate(ctx, params.Params.Login, params.Params.Pass)
	if err != nil {
		return operations.NewPostLoginForbidden().WithPayload(&models.AuthenticateResponse{Error: "invalid params"})
	}

	return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
		helpers.SetValueToCookies(w, "token", tokens.GetToken(), h.cfg.TokenTTL, false)
		w.WriteHeader(200)
		js, err := json.Marshal(models.AuthenticateResponse{Token: tokens.GetToken(), RefreshToken: tokens.GetRefreshToken(), Error: ""})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

func (h *handlers) Ping(params operations.GetPingParams, claims interface{}) middleware.Responder {
	return operations.NewGetPingOK()
}
