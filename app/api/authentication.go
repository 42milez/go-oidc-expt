package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/go-playground/validator/v10"
)

const sessionIDCookieName = config.SessionIDCookieName

var authenticateUserHdlr *AuthenticateHdlr

type AuthenticateHdlr struct {
	service   Authenticator
	cookie    CookieWriter
	sess      UserIdSessionWriter
	validator *validator.Validate
}

func NewAuthenticateHdlr(option *HandlerOption) (*AuthenticateHdlr, error) {
	return &AuthenticateHdlr{
		service:   service.NewAuthenticate(repository.NewUser(option.db, option.idGenerator), option.tokenGenerator),
		cookie:    option.cookie,
		sess:      option.SessionWriter,
		validator: option.validator,
	}, nil
}

func (a *AuthenticateHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var q *AuthorizeParams
	var reqBody *AuthenticateJSONRequestBody
	var err error

	if q, err = parseAuthorizeParam(r, a.validator); err != nil {
		a.respondError(w, r, err)
		return
	}

	if reqBody, err = a.parseRequestBody(r); err != nil {
		a.respondError(w, r, err)
		return
	}

	var userID typedef.UserID

	if userID, err = a.service.VerifyPassword(ctx, reqBody.Name, reqBody.Password); err != nil {
		a.respondError(w, r, err)
		return
	}

	var sid typedef.SessionID

	if sid, err = a.sess.SaveUserId(ctx, userID); err != nil {
		a.respondError(w, r, err)
		return
	}

	if err = a.cookie.Write(w, sessionIDCookieName, strconv.FormatUint(uint64(sid), 10), config.SessionIDCookieTTL); err != nil {
		a.respondError(w, r, err)
		return
	}

	var isConsented bool

	if isConsented, err = a.service.VerifyConsent(ctx, userID, q.ClientID); err != nil {
		a.respondError(w, r, err)
		return
	}

	if !isConsented {
		Redirect(w, r, config.ConsentPath, http.StatusFound)
		return
	}

	Redirect(w, r, config.AuthorizationPath, http.StatusFound)
}

func (a *AuthenticateHdlr) parseRequestBody(r *http.Request) (*AuthenticateJSONRequestBody, error) {
	var ret *AuthenticateJSONRequestBody

	if err := json.NewDecoder(r.Body).Decode(&ret); err != nil {
		return nil, err
	}

	if err := a.validator.Struct(ret); err != nil {
		return nil, xerr.FailedToValidate.Wrap(err)
	}

	return ret, nil
}

func (a *AuthenticateHdlr) respondError(w http.ResponseWriter, r *http.Request, err error) {
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
