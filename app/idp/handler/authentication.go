package handler

import (
	"encoding/json"
	"github.com/42milez/go-oidc-server/app/idp/cookie"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/auth"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/app/idp/service"
	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/rs/zerolog/log"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/go-playground/validator/v10"
)

func NewAuthenticate(entClient *ent.Client, cookie *cookie.Util, jwtUtil *auth.JWTUtil) (*Authenticate, error) {
	return &Authenticate{
		Cookie: cookie,
		Service: &service.Authenticate{
			Repo: &repository.User{
				Clock: &xutil.RealClocker{},
				DB:    entClient,
			},
			Token: jwtUtil,
		},
		Validator: validator.New(),
	}, nil
}

type Authenticate struct {
	Service   Authenticator
	Cookie    *cookie.Util
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

	if err := p.Service.Authenticate(r.Context(), body.Name, body.Password); err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	// TODO: Create a session with secret key
	// ...

	// TODO: Save session id in cookie
	// ...

	// TODO: Redirect to consent url
	// ...

	resp := ""

	RespondJSON(w, http.StatusOK, resp)
}
