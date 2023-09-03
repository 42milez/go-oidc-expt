package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/42milez/go-oidc-server/app/api/cookie"
	"github.com/42milez/go-oidc-server/app/api/session"

	"github.com/42milez/go-oidc-server/app/model"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/redis/go-redis/v9"

	"github.com/42milez/go-oidc-server/pkg/xtime"

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

// ServeHTTP authenticates user
//
//	@summary		authenticates user
//	@description	This endpoint authenticates user.
//	@id				Authenticate.ServeHTTP
//	@tags			User
//	@accept			json
//	@produce		json
//	@param			user	body		model.AuthenticateRequest	true	"user credential"
//	@success		200		{object}	model.AuthenticateResponse
//	@failure		400		{object}	model.ErrorResponse
//	@failure		401		{object}	model.ErrorResponse
//	@failure		500		{object}	model.ErrorResponse
//	@router			/v1/user/authenticate [post]
func (p *Authenticate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var reqBody model.AuthenticateRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Error().Err(err).Msg(errFailedToDecodeRequestBody)
		ResponseJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	if err := p.validator.Struct(reqBody); err != nil {
		log.Error().Err(err).Msg(errValidationError)
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.AuthenticationFailed,
		})
		return
	}

	userID, err := p.Service.Authenticate(r.Context(), reqBody.Name, reqBody.Password)

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
			ResponseJson500(w, xerr.UnexpectedErrorOccurred)
			return
		}
	}

	sessionID, err := p.Session.Create(r.Context(), &entity.Session{
		UserID: userID,
	})

	if err != nil {
		ResponseJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	if err = p.Cookie.Set(w, sessionIDCookieName, sessionID, config.SessionIDCookieTTL); err != nil {
		ResponseJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	Redirect(w, r, config.ConsentURL, http.StatusFound)
}
