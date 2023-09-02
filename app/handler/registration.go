package handler

import (
	"encoding/json"
	"net/http"

	"github.com/42milez/go-oidc-server/app/handler/session"

	"github.com/42milez/go-oidc-server/pkg/xid"

	"github.com/42milez/go-oidc-server/pkg/xtime"

	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/model"
	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"
	"github.com/rs/zerolog/log"

	"github.com/go-playground/validator/v10"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

func NewRegisterUser(ec *ent.Client, idGen *xid.UniqueID, sess *session.Session) (*RegisterUser, error) {
	return &RegisterUser{
		Service: &service.CreateUser{
			Repo: &repository.User{
				Clock: &xtime.RealClocker{},
				DB:    ec,
				IDGen: idGen,
			},
		},
		Session:   sess,
		validator: validator.New(),
	}, nil
}

type RegisterUser struct {
	Service   UserCreator
	Session   SessionRestorer
	validator *validator.Validate
}

// ServeHTTP registers user
//
//	@summary		registers user
//	@description	This endpoint registers user.
//	@id				Register.ServeHTTP
//	@tags			User
//	@accept			json
//	@produce		json
//	@param			user	body		model.RegisterUserRequest	true	"user credential"
//	@success		200		{object}	model.RegisterUserResponse
//	@failure		400		{object}	model.ErrorResponse
//	@failure		401		{object}	model.ErrorResponse
//	@failure		500		{object}	model.ErrorResponse
//	@router			/v1/user/register [post]
func (p *RegisterUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	if err := p.validator.Struct(req); err != nil {
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

	resp := model.RegisterUserResponse{
		ID:   user.ID,
		Name: user.Name,
	}

	RespondJSON(w, http.StatusOK, resp)
}
