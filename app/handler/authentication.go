package handler

import (
	"encoding/json"
	"net/http"

	"github.com/42milez/go-oidc-server/app/auth"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/entity"
	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"
	"github.com/42milez/go-oidc-server/app/session"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/rs/zerolog/log"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/go-playground/validator/v10"
)

func NewAuthenticate(entClient *ent.Client, cookieUtil *Cookie, sessionUtil *session.Session, jwtUtil *auth.Util) (*Authenticate, error) {
	return &Authenticate{
		service: &service.Authenticate{
			Repo: &repository.User{
				Clock: &xutil.RealClocker{},
				DB:    entClient,
			},
			Token: jwtUtil,
		},
		cookie:    cookieUtil,
		session:   sessionUtil,
		validator: validator.New(),
	}, nil
}

type Authenticate struct {
	service   Authenticator
	session   SessionCreator
	cookie    *Cookie
	validator *validator.Validate
}

const cookieNameSessionID = "sid"

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
		ResponseJSONWithInternalServerError(w)
		return
	}

	if err := p.validator.Struct(body); err != nil {
		log.Error().Err(err).Msg(errValidationError)
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.AuthenticationFailed,
		})
		return
	}

	userID, err := p.service.Authenticate(r.Context(), body.Name, body.Password)

	if err != nil {
		ResponseJSONWithInternalServerError(w)
		return
	}

	sessionID, err := p.session.Create(&entity.UserSession{
		ID: userID,
	})

	if err != nil {
		ResponseJSONWithInternalServerError(w)
		return
	}

	if err = p.cookie.Set(w, cookieNameSessionID, sessionID); err != nil {
		ResponseJSONWithInternalServerError(w)
		return
	}

	// TODO: Redirect to consent url
	// ...

	resp := struct {
		Error string `json:"error"`
	}{
		Error: "",
	}
	RespondJSON(w, http.StatusOK, resp)
}
