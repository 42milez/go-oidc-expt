package api

import (
	"encoding/json"
	"errors"
	"net/http"

	auth "github.com/42milez/go-oidc-server/app/pkg/xjwt"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"

	"github.com/42milez/go-oidc-server/app/api/cookie"
	"github.com/42milez/go-oidc-server/app/api/session"

	"github.com/42milez/go-oidc-server/app/model"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/redis/go-redis/v9"

	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"
	"github.com/rs/zerolog/log"

	"github.com/go-playground/validator/v10"
)

func NewAuthenticateUser(ec *ent.Client, rc *redis.Client, cookie *cookie.Cookie, jwt *auth.JWT, sess *session.Session) (*AuthenticateUser, error) {
	return &AuthenticateUser{
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

type AuthenticateUser struct {
	Service   Authenticator
	Session   SessionCreator
	Cookie    *cookie.Cookie
	validator *validator.Validate
}

const sessionIDCookieName = config.SessionIDCookieName

func (p *AuthenticateUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var reqBody model.AuthenticateRequest

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Error().Err(err).Msg(errFailedToDecodeRequestBody)
		RespondJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	if err := p.validator.Struct(reqBody); err != nil {
		log.Error().Err(err).Msg(errValidationError)
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.InvalidRequest,
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
			RespondJson500(w, xerr.UnexpectedErrorOccurred)
			return
		}
	}

	sessionID, err := p.Session.Create(r.Context(), &entity.Session{
		UserID: userID,
	})

	if err != nil {
		RespondJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	if err = p.Cookie.Set(w, sessionIDCookieName, sessionID, config.SessionIDCookieTTL); err != nil {
		RespondJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	Redirect(w, r, config.ConsentURL, http.StatusFound)
}
