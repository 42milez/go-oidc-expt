package api

import (
	"encoding/json"
	"net/http"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"

	"github.com/42milez/go-oidc-server/app/api/session"

	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/model"
	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"
	"github.com/rs/zerolog/log"

	"github.com/go-playground/validator/v10"
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

func (p *RegisterUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	if err := p.validator.Struct(req); err != nil {
		log.Error().Err(err).Msg(errValidationError)
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.InvalidRequest,
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
