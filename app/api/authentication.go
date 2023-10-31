package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/httpstore"

	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/config"
)

const sessionIDCookieName = config.SessionIDCookieName

var authenticateUserHdlr *AuthenticateHdlr

type AuthenticateHdlr struct {
	svc  Authenticator
	ck   iface.CookieWriter
	sess iface.UserInfoSessionWriter
	v    iface.StructValidator
}

func NewAuthenticateHdlr(option *HandlerOption) (*AuthenticateHdlr, error) {
	return &AuthenticateHdlr{
		svc:  service.NewAuthenticate(repository.NewUser(option.db, option.idGenerator), option.tokenGenerator),
		ck:   option.cookie,
		sess: httpstore.NewWriteSession(repository.NewSession(option.cache), &httpstore.Context{}, option.idGenerator),
		v:    option.validator,
	}, nil
}

func (ah *AuthenticateHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var q *AuthorizeParams
	var reqBody *AuthenticateJSONRequestBody
	var err error

	if q, err = parseAuthorizeParam(r, ah.v); err != nil {
		ah.respondError(w, r, err)
		return
	}

	if reqBody, err = ah.parseRequestBody(r); err != nil {
		ah.respondError(w, r, err)
		return
	}

	var userID typedef.UserID

	if userID, err = ah.svc.VerifyPassword(ctx, reqBody.Name, reqBody.Password); err != nil {
		ah.respondError(w, r, err)
		return
	}

	var sid typedef.SessionID

	if sid, err = ah.sess.WriteUserInfo(ctx, userID); err != nil {
		ah.respondError(w, r, err)
		return
	}

	if err = ah.ck.Write(w, sessionIDCookieName, strconv.FormatUint(uint64(sid), 10), config.SessionIDCookieTTL); err != nil {
		ah.respondError(w, r, err)
		return
	}

	var isConsented bool

	if isConsented, err = ah.svc.VerifyConsent(ctx, userID, q.ClientID); err != nil {
		ah.respondError(w, r, err)
		return
	}

	if !isConsented {
		Redirect(w, r, config.ConsentPath, http.StatusFound)
		return
	}

	Redirect(w, r, config.AuthorizationPath, http.StatusFound)
}

func (ah *AuthenticateHdlr) parseRequestBody(r *http.Request) (*AuthenticateJSONRequestBody, error) {
	var ret *AuthenticateJSONRequestBody

	if err := json.NewDecoder(r.Body).Decode(&ret); err != nil {
		return nil, err
	}

	if err := ah.v.Struct(ret); err != nil {
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
