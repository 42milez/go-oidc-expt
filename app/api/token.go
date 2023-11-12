package api

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/42milez/go-oidc-server/app/option"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/httpstore"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
	"github.com/42milez/go-oidc-server/app/service"
)

var tokenHdlr *TokenHdlr

func NewTokenHdlr(opt *option.Option) *TokenHdlr {
	return &TokenHdlr{
		cache:   httpstore.NewCache(opt),
		context: &httpstore.Context{},
		svc:     service.NewToken(opt),
		v:       opt.V,
	}
}

type TokenHdlr struct {
	svc     TokenRequestAcceptor
	cache   TokenCacheReadWriter
	context iface.ContextReader
	v       iface.StructValidator
}

func (t *TokenHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	credentials, err := extractCredential(r)
	if err != nil {
		RespondTokenRequestError(w, r, xerr.InvalidGrant)
		return
	}
	clientId := credentials[0]

	param, err := t.parseForm(r)
	if err != nil {
		RespondServerError(w, r)
		return
	}

	if err = t.v.Struct(param); err != nil {
		RespondTokenRequestError(w, r, xerr.InvalidRequest)
		return
	}

	if param.GrantType == config.AuthorizationCodeGrantType {
		t.handleAuthCodeGrant(w, r, param, clientId)
		return
	} else if param.GrantType == config.RefreshTokenGrantType {
		t.handleRefreshTokenGrant(w, r, param, clientId)
		return
	} else {
		RespondTokenRequestError(w, r, xerr.InvalidRequest)
		return
	}
}

func (t *TokenHdlr) handleAuthCodeGrant(w http.ResponseWriter, r *http.Request, param *TokenFormdataBody, clientId string) {
	ctx := r.Context()

	if xutil.IsEmpty(*param.Code) || xutil.IsEmpty(*param.RedirectUri) {
		RespondTokenRequestError(w, r, xerr.InvalidRequest)
		return
	}

	if err := t.revokeAuthCode(ctx, *param.Code, clientId); err != nil {
		t.respondAuthCodeError(w, r, err)
		return
	}

	authParam, err := t.cache.ReadOpenIdParam(ctx, clientId, *param.Code)
	if err != nil {
		t.respondAuthCodeError(w, r, err)
		return
	}

	if *param.RedirectUri != authParam.RedirectURI {
		RespondTokenRequestError(w, r, xerr.InvalidGrant)
		return
	}

	tokens, err := t.generateToken(authParam.UserId)
	if err != nil {
		RespondServerError(w, r)
		return
	}

	if err = t.cache.WriteRefreshTokenPermission(ctx, *tokens[refreshTokenKey], clientId, authParam.UserId); err != nil {
		RespondServerError(w, r)
		return
	}

	resp := &TokenResponse{
		AccessToken:  *tokens[accessTokenKey],
		RefreshToken: *tokens[refreshTokenKey],
		IdToken:      tokens[idTokenKey],
		TokenType:    config.BearerTokenType,
		ExpiresIn:    3600,
	}

	RespondJSON(w, r, http.StatusOK, resp)
}

func (t *TokenHdlr) revokeAuthCode(ctx context.Context, code, clientId string) error {
	if err := t.svc.ValidateAuthCode(ctx, code, clientId); err != nil {
		return err
	}
	if err := t.svc.RevokeAuthCode(ctx, code, clientId); err != nil {
		return err
	}
	return nil
}

const accessTokenKey = "AccessToken"
const refreshTokenKey = "RefreshToken"
const idTokenKey = "IdToken"

func (t *TokenHdlr) generateToken(uid typedef.UserID) (map[string]*string, error) {
	accessToken, err := t.svc.GenerateAccessToken(uid)
	if err != nil {
		return nil, err
	}

	refreshToken, err := t.svc.GenerateRefreshToken(uid)
	if err != nil {
		return nil, err
	}

	idToken, err := t.svc.GenerateIdToken(uid)
	if err != nil {
		return nil, err
	}

	return map[string]*string{
		accessTokenKey:  &accessToken,
		refreshTokenKey: &refreshToken,
		idTokenKey:      &idToken,
	}, nil
}

func (t *TokenHdlr) respondAuthCodeError(w http.ResponseWriter, r *http.Request, err error) {
	invalidGrant := errors.Is(err, xerr.AuthCodeNotFound) || errors.Is(err, xerr.AuthCodeExpired) ||
		errors.Is(err, xerr.AuthCodeUsed)
	if invalidGrant {
		RespondTokenRequestError(w, r, xerr.InvalidGrant)
		return
	}

	if errors.Is(err, xerr.UnauthorizedRequest) {
		RespondTokenRequestError(w, r, xerr.InvalidRequest)
	}

	RespondServerError(w, r)
}

func (t *TokenHdlr) handleRefreshTokenGrant(w http.ResponseWriter, r *http.Request, param *TokenFormdataBody, clientId string) {
	ctx := r.Context()

	perm, err := t.svc.ReadRefreshTokenPermission(ctx, *param.RefreshToken, clientId)
	if err != nil {
		if errors.Is(err, xerr.InvalidToken) || errors.Is(err, xerr.ClientIdNotMatched) {
			RespondJSON400(w, r, xerr.InvalidRequest2, nil, err)
		} else if errors.Is(err, xerr.RefreshTokenPermissionNotFound) {
			RespondJSON401(w, r, xerr.InvalidRequest2, nil, err)
		} else {
			RespondJSON500(w, r, err)
		}
	}

	accessToken, err := t.svc.GenerateAccessToken(perm.UserId)
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	refreshToken, err := t.svc.GenerateRefreshToken(perm.UserId)
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	if err = t.cache.WriteRefreshTokenPermission(ctx, refreshToken, clientId, perm.UserId); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	resp := &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    config.BearerTokenType,
		ExpiresIn:    3600,
	}

	RespondJSON(w, r, http.StatusOK, resp)
}

func (t *TokenHdlr) parseForm(r *http.Request) (*TokenFormdataBody, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	params, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
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
