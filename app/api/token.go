package api

import (
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/42milez/go-oidc-server/app/httpstore"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/go-playground/validator/v10"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
	"github.com/42milez/go-oidc-server/app/service"
)

var tokenHdlr *TokenHdlr

func NewTokenHdlr(option *HandlerOption) *TokenHdlr {
	return &TokenHdlr{
		svc: service.NewToken(option.db, option.clock, option.sessionReader, option.tokenGenerator),
		cr:  &httpstore.ReadContext{},
		v:   option.validator,
	}
}

type TokenHdlr struct {
	svc TokenRequestAcceptor
	cr  ContextReader
	v   *validator.Validate
}

func (t *TokenHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	credentials, err := extractCredential(r)
	if err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
	}

	clientId := credentials[0]

	param, err := t.parseForm(r)
	if err != nil {
		if errors.Is(err, xerr.MalformedFormParameter) {
			RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
			return
		}
		RespondJSON500(w, r, err)
		return
	}

	if err = t.v.Struct(param); err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	if param.GrantType == config.AuthorizationCodeGrantType {
		t.handleAuthCodeGrantType(w, r, param, clientId)
		return
	} else if param.GrantType == config.RefreshTokenGrantType {
		t.handleRefreshTokenGrantType(w, r, param)
		return
	} else {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}
}

func (t *TokenHdlr) handleAuthCodeGrantType(w http.ResponseWriter, r *http.Request, param *TokenFormdataBody, clientId string) {
	ctx := r.Context()

	if xutil.IsEmpty(*param.Code) || xutil.IsEmpty(*param.RedirectUri) {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, nil)
		return
	}

	if err := t.svc.ValidateAuthCode(ctx, *param.Code, clientId); err != nil {
		t.respondAuthCodeError(w, r, err)
		return
	}

	if err := t.svc.RevokeAuthCode(ctx, *param.Code, clientId); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	if err := t.svc.ValidateRedirectUri(ctx, *param.RedirectUri, clientId); err != nil {
		if errors.Is(err, xerr.RedirectUriNotFound) || errors.Is(err, xerr.RedirectUriNotMatched) {
			RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		} else {
			RespondJSON500(w, r, err)
		}
		return
	}

	uid, ok := t.cr.Read(ctx, typedef.UserIDKey{}).(typedef.UserID)
	if !ok {
		RespondJSON401(w, r, xerr.UnauthorizedRequest, nil, nil)
		return
	}

	tokenSet, err := t.svc.CreateTokenSet(uid)
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	resp := &TokenResponse{
		AccessToken:  tokenSet.AccessToken,
		RefreshToken: tokenSet.RefreshToken,
		IdToken:      tokenSet.IdToken,
		TokenType:    config.BearerTokenType,
		ExpiresIn:    3600,
	}

	RespondJSON(w, r, http.StatusOK, resp)
}

func (t *TokenHdlr) respondAuthCodeError(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, xerr.AuthCodeNotFound) {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
	} else if errors.Is(err, xerr.AuthCodeExpired) {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
	} else if errors.Is(err, xerr.AuthCodeUsed) {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
	} else {
		RespondJSON500(w, r, err)
	}
}

func (t *TokenHdlr) handleRefreshTokenGrantType(w http.ResponseWriter, r *http.Request, param *TokenFormdataBody) {
	if err := t.svc.ValidateRefreshToken(param.RefreshToken); err != nil {
		if errors.Is(err, xerr.InvalidToken) {
			RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		}
		return
	}

	// TODO: Return access token
	// ...
}

func (t *TokenHdlr) parseForm(r *http.Request) (*TokenFormdataBody, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	params, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, xerr.MalformedFormParameter
	}

	code := params.Get("code")
	grantType := params.Get("grant_type")
	redirectUri := params.Get("redirect_uri")
	refreshToken := params.Get("refresh_token")

	ret := &TokenFormdataBody{
		GrantType: grantType,
	}

	if !xutil.IsEmpty(code) {
		ret.Code = &code
	}

	if !xutil.IsEmpty(redirectUri) {
		ret.RedirectUri = &redirectUri
	}

	if !xutil.IsEmpty(refreshToken) {
		ret.RefreshToken = &refreshToken
	}

	return ret, nil
}
