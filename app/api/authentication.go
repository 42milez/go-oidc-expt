package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/gorilla/schema"

	"github.com/42milez/go-oidc-server/app/api/oapigen"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/go-playground/validator/v10"
)

type AuthenticateHdlr struct {
	service   Authenticator
	cookie    CookieWriter
	session   SessionCreator
	validator *validator.Validate
}

const sessionIDCookieName = config.SessionIDCookieName

func (ah *AuthenticateHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var q *oapigen.AuthorizeParams
	var reqBody *oapigen.AuthenticateJSONRequestBody
	var err error

	if q, err = ah.parseQueryParam(r); err != nil {
		ah.respondError(w, err)
		return
	}

	if reqBody, err = ah.parseRequestBody(r); err != nil {
		ah.respondError(w, err)
		return
	}

	var userID typedef.UserID

	if userID, err = ah.service.VerifyPassword(ctx, reqBody.Name, reqBody.Password); err != nil {
		ah.respondError(w, err)
		return
	}

	var sessionID string

	if sessionID, err = ah.session.Create(ctx, &entity.Session{
		UserID: userID,
	}); err != nil {
		ah.respondError(w, err)
		return
	}

	if err = ah.cookie.Write(w, sessionIDCookieName, sessionID, config.SessionIDCookieTTL); err != nil {
		ah.respondError(w, err)
		return
	}

	var isConsented bool

	if isConsented, err = ah.service.VerifyConsent(ctx, userID, q.ClientId); err != nil {
		ah.respondError(w, err)
		return
	}

	if !isConsented {
		Redirect(w, r, config.ConsentEndpoint, http.StatusFound)
		return
	}

	Redirect(w, r, config.AuthorizationEndpoint, http.StatusFound)
}

func (ah *AuthenticateHdlr) parseQueryParam(r *http.Request) (*oapigen.AuthorizeParams, error) {
	decoder := schema.NewDecoder()
	ret := &oapigen.AuthorizeParams{}

	if err := decoder.Decode(ret, r.URL.Query()); err != nil {
		return nil, err
	}

	if err := ah.validator.Struct(ret); err != nil {
		return nil, xerr.FailedToValidate.Wrap(err)
	}

	return ret, nil
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

func (ah *AuthenticateHdlr) respondError(w http.ResponseWriter, err error) {
	if errors.Is(err, xerr.FailedToValidate) {
		RespondJSON400(w, xerr.InvalidRequest, nil, err)
		return
	}

	if errors.Is(err, xerr.PasswordNotMatched) || errors.Is(err, xerr.UserNotFound) {
		RespondJSON401(w, xerr.InvalidUsernameOrPassword, nil, err)
		return
	}

	RespondJSON500(w, err)

	return
}
