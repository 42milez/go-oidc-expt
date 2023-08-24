package handler

import (
	"encoding/json"
	"github.com/42milez/go-oidc-server/app/idp/cookie"
	"github.com/42milez/go-oidc-server/app/idp/entity"
	"github.com/42milez/go-oidc-server/app/idp/jwt"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/app/idp/service"
	"github.com/42milez/go-oidc-server/app/idp/session"
	"github.com/42milez/go-oidc-server/pkg/xutil"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/rs/zerolog/log"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/go-playground/validator/v10"
)

func NewAuthenticate(entClient *ent.Client, cookieUtil *cookie.Util, sessionUtil *session.Util, jwtUtil *jwt.Util) (*Authenticate, error) {
	return &Authenticate{
		Service: &service.Authenticate{
			Repo: &repository.User{
				Clock: &xutil.RealClocker{},
				DB:    entClient,
			},
			Token: jwtUtil,
		},
		Cookie: cookieUtil,
		Session: sessionUtil,
		Validator: validator.New(),
	}, nil
}

type Authenticate struct {
	Service   Authenticator
	Cookie    *cookie.Util
	Session   *session.Util
	Validator *validator.Validate
}

// ServeHTTP authenticates a user
//
//	@summary		TBD
//	@description	TBD
//	@id				Authenticate.ServeHTTP
//	@tags			authentication
//	@accept			json
//	@produce		json
//	@param			name		body		string	true	"TBD"
//	@param			password	body		string	true	"TBD"
//	@success		200			{object}	model.Authenticate
//	@failure		500			{object}	model.BadRequest
//	@failure		500			{object}	model.InternalServerError
//	@router			/v1/authenticate [post]
func (p *Authenticate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Error().Err(err).Msg(errFailedToDecodeRequestBody)
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	if err := p.Validator.Struct(body); err != nil {
		log.Error().Err(err).Msg(errValidationError)
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.AuthenticationFailed,
		})
		return
	}

	userID, err := p.Service.Authenticate(r.Context(), body.Name, body.Password)

	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	sessionID, err := p.Session.Create(&entity.UserSession{
		ID: userID,
	})

	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	if err = p.Cookie.SetSessionID(w, sessionID); err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	// TODO: Redirect to consent url
	// ...

	resp := ""

	RespondJSON(w, http.StatusOK, resp)
}
