package handler

import (
	"encoding/json"
	"net/http"

	"github.com/42milez/go-oidc-server/app/auth"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/model"
	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"
	"github.com/42milez/go-oidc-server/pkg/xutil"
	"github.com/rs/zerolog/log"

	"github.com/go-playground/validator/v10"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

func NewCreateUser(entClient *ent.Client, sessionUtil *Session, jwtUtil *auth.Util) (*CreateUser, error) {
	return &CreateUser{
		Service: &service.CreateUser{
			Repo: &repository.User{
				Clock: &xutil.RealClocker{},
				DB:    entClient,
			},
		},
		Session:   sessionUtil,
		Validator: validator.New(),
	}, nil
}

type CreateUser struct {
	Service   UserCreator
	Session   SessionRestorer
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
