package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/auth"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/app/idp/service"
	"github.com/42milez/go-oidc-server/pkg/xutil"
	"github.com/redis/go-redis/v9"

	"github.com/rs/zerolog/log"

	"github.com/go-playground/validator/v10"

	"github.com/42milez/go-oidc-server/app/idp/model"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

func NewCreateUser(entClient *ent.Client, redisClient *redis.Client) (*CreateUser, error) {
	jwtUtil, err := auth.NewJWTUtil(&xutil.RealClocker{})

	if err != nil {
		return nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	return &CreateUser{
		Service: &service.CreateUser{
			Repo: &repository.User{
				Clock: &xutil.RealClocker{},
				DB:    entClient,
			},
		},
		Session: &service.Session{
			Repo: &repository.Session{
				Cache: redisClient,
				JWT:   jwtUtil,
			},
		},
		Validator: validator.New(),
	}, nil
}

type CreateUser struct {
	Service   UserCreator
	Session   SessionManager
	Validator *validator.Validate
}

func (p *CreateUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req model.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	if err := p.Validator.Struct(req); err != nil {
		log.Error().Err(err).Msg(errValidationError)
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.AuthenticationFailed,
		})
		return
	}

	_, err := p.Service.CreateUser(r.Context(), req.Name, req.Password)

	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}
}

type SelectUser struct {
	Service UserSelector
}

func (p *SelectUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := p.Service.SelectUser(ctx)

	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
	}

	resp := model.CreateUserResponse{
		ID:   user.ID,
		Name: user.Name,
	}

	RespondJSON(w, http.StatusOK, resp)
}
