package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/42milez/go-oidc-expt/pkg/typedef"

	"github.com/42milez/go-oidc-expt/cmd/config"
	"github.com/42milez/go-oidc-expt/cmd/httpstore"
	"github.com/42milez/go-oidc-expt/cmd/iface"
	"github.com/42milez/go-oidc-expt/cmd/option"
	"github.com/42milez/go-oidc-expt/cmd/service"
	"github.com/42milez/go-oidc-expt/pkg/xerr"
)

const sessionIDCookieName = config.SessionIDCookieName

var authentication *Authentication

func InitAuthentication(opt *option.Option) {
	if authentication == nil {
		authentication = &Authentication{
			svc:    service.NewAuthenticate(opt),
			cache:  httpstore.NewCache(opt),
			cookie: opt.Cookie,
			v:      opt.V,
		}
	}
}

type Authentication struct {
	svc    Authenticator
	cache  iface.SessionCreator
	cookie iface.CookieWriter
	v      iface.StructValidator
}

func (ah *Authentication) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var q *AuthorizeParams
	var reqBody *AuthenticateJSONRequestBody
	var err error

	if q, err = parseAuthorizeParam(r, ah.v); err != nil {
		respondError(w, r, err)
		return
	}

	if reqBody, err = ah.parseRequestBody(r); err != nil {
		respondError(w, r, err)
		return
	}

	var userID typedef.UserID

	if userID, err = ah.svc.VerifyPassword(ctx, reqBody.Name, reqBody.Password); err != nil {
		respondError(w, r, err)
		return
	}

	var sid typedef.SessionID

	if sid, err = ah.cache.CreateSession(ctx, userID); err != nil {
		respondError(w, r, err)
		return
	}

	if err = ah.cookie.Write(w, sessionIDCookieName, strconv.FormatUint(uint64(sid), 10), config.SessionIDCookieTTL); err != nil {
		respondError(w, r, err)
		return
	}

	var isConsented bool

	if isConsented, err = ah.svc.VerifyConsent(ctx, userID, q.ClientID); err != nil {
		respondError(w, r, err)
		return
	}

	if !isConsented {
		Redirect(w, r, "/consent", http.StatusFound)
		return
	}

	Redirect(w, r, config.AuthorizationPath(), http.StatusFound)
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

func respondError(w http.ResponseWriter, r *http.Request, err error) {
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
