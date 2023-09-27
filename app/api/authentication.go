package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/api/oapigen"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/go-playground/validator/v10"
)

const sessionIDCookieName = config.SessionIDCookieName

var authenticateUserHdlr *AuthenticateHdlr

type AuthenticateHdlr struct {
	service   Authenticator
	cookie    CookieWriter
	session   SessionCreator
	validator *validator.Validate
}

func NewAuthenticateHdlr(option *HandlerOption) (*AuthenticateHdlr, error) {
	return &AuthenticateHdlr{
		service:   service.NewAuthenticate(repository.NewUser(option.db, option.idGenerator), option.jwtUtil),
		cookie:    option.cookie,
		session:   option.sessionCreator,
		validator: option.validator,
	}, nil
}

func (ah *AuthenticateHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var q *oapigen.AuthorizeParams
	var reqBody *oapigen.AuthenticateJSONRequestBody
	var err error

	if q, err = parseAuthorizeParam(r, ah.validator); err != nil {
		ah.respondError(w, r, err)
		return
	}

	if reqBody, err = ah.parseRequestBody(r); err != nil {
		ah.respondError(w, r, err)
		return
	}

	var userID typedef.UserID

	if userID, err = ah.service.VerifyPassword(ctx, reqBody.Name, reqBody.Password); err != nil {
		ah.respondError(w, r, err)
		return
	}

	var sessionID string

	if sessionID, err = ah.session.Create(ctx, &entity.Session{
		UserID: userID,
	}); err != nil {
		ah.respondError(w, r, err)
		return
	}

	if err = ah.cookie.Write(w, sessionIDCookieName, sessionID, config.SessionIDCookieTTL); err != nil {
		ah.respondError(w, r, err)
		return
	}

	var isConsented bool

	if isConsented, err = ah.service.VerifyConsent(ctx, userID, q.ClientId); err != nil {
		ah.respondError(w, r, err)
		return
	}

	if !isConsented {
		Redirect(w, r, config.ConsentEndpoint, http.StatusFound)
		return
	}

	Redirect(w, r, config.AuthorizationEndpoint, http.StatusFound)
}

func (ah *AuthenticateHdlr) parseRequestBody(r *http.Request) (*oapigen.AuthenticateJSONRequestBody, error) {
	var ret *oapigen.AuthenticateJSONRequestBody

	if err := json.NewDecoder(r.Body).Decode(&ret); err != nil {
		return nil, err
	}

	if err := ah.validator.Struct(ret); err != nil {
		return nil, xerr.FailedToValidate.Wrap(err)
	}

	return ret, nil
}

func (ah *AuthenticateHdlr) respondError(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, xerr.FailedToValidate) {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	if errors.Is(err, xerr.PasswordNotMatched) || errors.Is(err, xerr.UserNotFound) {
		RespondJSON401(w, r, xerr.InvalidUsernameOrPassword, nil, err)
		return
	}

	RespondJSON500(w, r, err)

	return
}
