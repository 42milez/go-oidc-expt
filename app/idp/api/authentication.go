package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/httpstore"
	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/idp/service"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

const sessionIDCookieName = config.SessionIDCookieName

var authentication *Authentication

func NewAuthentication(opt *option.Option) *Authentication {
	return &Authentication{
		svc:    service.NewAuthenticate(opt),
		cache:  httpstore.NewCache(opt),
		cookie: opt.Cookie,
		v:      opt.V,
	}
}

type Authentication struct {
	svc    Authenticator
	cache  iface.UserInfoWriter
	cookie iface.CookieWriter
	v      iface.StructValidator
}

func (ah *Authentication) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	if sid, err = ah.cache.WriteUserInfo(ctx, userID); err != nil {
		ah.respondError(w, r, err)
		return
	}

	if err = ah.cookie.Write(w, sessionIDCookieName, strconv.FormatUint(uint64(sid), 10), config.SessionIDCookieTTL); err != nil {
		ah.respondError(w, r, err)
		return
	}

	var isConsented bool

	if isConsented, err = ah.svc.VerifyConsent(ctx, userID, q.ClientID); err != nil {
		ah.respondError(w, r, err)
		return
	}

	if !isConsented {
		Redirect2(w, r, config.ConsentPath, http.StatusFound)
		return
	}

	Redirect2(w, r, config.AuthorizationPath, http.StatusFound)
}

func (ah *Authentication) parseRequestBody(r *http.Request) (*AuthenticateJSONRequestBody, error) {
	var ret *AuthenticateJSONRequestBody

	if err := json.NewDecoder(r.Body).Decode(&ret); err != nil {
		return nil, err
	}

	if err := ah.v.Struct(ret); err != nil {
		return nil, xerr.FailedToValidate.Wrap(err)
	}

	return ret, nil
}

func (ah *Authentication) respondError(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, xerr.FailedToValidate) {
		RespondJSON400(w, r, xerr.InvalidRequest2, nil, err)
		return
	}

	if errors.Is(err, xerr.PasswordNotMatched) || errors.Is(err, xerr.UserNotFound) {
		RespondJSON401(w, r, xerr.InvalidUsernameOrPassword, nil, err)
		return
	}

	RespondJSON500(w, r, err)

	return
}
