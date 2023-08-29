package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/cookie"

	"github.com/42milez/go-oidc-server/app/session"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/redis/go-redis/v9"

	"github.com/42milez/go-oidc-server/pkg/xtime"

	"github.com/42milez/go-oidc-server/app/model"

	"github.com/42milez/go-oidc-server/app/auth"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"
	"github.com/rs/zerolog/log"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/go-playground/validator/v10"
)

func NewAuthenticate(ec *ent.Client, rc *redis.Client, cookie *cookie.Cookie, jwt *auth.JWT, sess *session.Session) (*Authenticate, error) {
	return &Authenticate{
		Service: &service.Authenticate{
			Repo: &repository.User{
				Clock: &xtime.RealClocker{},
				DB:    ec,
			},
			Token: jwt,
		},
		Cookie:    cookie,
		Session:   sess,
		validator: validator.New(),
	}, nil
}

type Authenticate struct {
	Service   Authenticator
	Session   SessionCreator
	Cookie    *cookie.Cookie
	validator *validator.Validate
}

const sessionIDCookieName = config.SessionIDCookieName

// ServeHTTP authenticates a user
//
//	@summary		TBD
//	@description	TBD
//	@id				Authenticate.ServeHTTP
//	@tags			authentication
//	@accept			json
//	@produce		json
//	@param			name		body		string	true	"Username"
//	@param			password	body		string	true	"Password"
//	@success		200			{object}	model.AuthenticateResponse
//	@failure		500			{object}	model.ErrorResponse
//	@failure		500			{object}	model.ErrorResponse
//	@router			/v1/authenticate [post]
func (p *Authenticate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Error().Err(err).Msg(errFailedToDecodeRequestBody)
		ResponseJsonWithInternalServerError(w)
		return
	}

	if err := p.validator.Struct(body); err != nil {
		log.Error().Err(err).Msg(errValidationError)
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.AuthenticationFailed,
		})
		return
	}

	userID, err := p.Service.Authenticate(r.Context(), body.Name, body.Password)

	if err != nil {
		if errors.Is(err, xerr.UserNotFound) {
			RespondJSON(w, http.StatusUnauthorized, &ErrResponse{
				Error: xerr.InvalidUsernameOrPassword,
			})
			return
		} else if errors.Is(err, xerr.PasswordNotMatched) {
			RespondJSON(w, http.StatusUnauthorized, &ErrResponse{
				Error: xerr.InvalidUsernameOrPassword,
			})
			return
		} else {
			ResponseJsonWithInternalServerError(w)
			return
		}
	}

	sessionID, err := p.Session.Create(r.Context(), &entity.UserSession{
		ID: userID,
	})

	if err != nil {
		ResponseJsonWithInternalServerError(w)
		return
	}

	if err = p.Cookie.Set(w, sessionIDCookieName, sessionID, config.SessionIDCookieTTL); err != nil {
		ResponseJsonWithInternalServerError(w)
		return
	}

	// TODO: Redirect to consent url
	// ...

	RespondJSON(w, http.StatusOK, &model.AuthenticateResponse{
		Error: xerr.OK.Error(),
	})
}
