package api

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/httpstore"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
	"github.com/42milez/go-oidc-server/app/service"
)

var tokenHdlr *TokenHdlr

func NewTokenHdlr(option *HandlerOption) *TokenHdlr {
	return &TokenHdlr{
		svc: service.NewToken(option.db, option.cache, option.clock, option.tokenGenerator),
		ctx: &httpstore.Context{},
		v:   option.validator,
	}
}

type TokenHdlr struct {
	svc  TokenRequestAcceptor
	ctx  iface.ContextReader
	sess iface.RefreshTokenOwnerSessionWriter
	v    iface.StructValidator
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

func (t *TokenHdlr) revokeAuthCode(ctx context.Context, code, clientId string) error {
	if err := t.svc.ValidateAuthCode(ctx, code, clientId); err != nil {
		return err
	}
	if err := t.svc.RevokeAuthCode(ctx, code, clientId); err != nil {
		return err
	}
	return nil
}

func (t *TokenHdlr) handleAuthCodeGrantType(w http.ResponseWriter, r *http.Request, param *TokenFormdataBody, clientId string) {
	ctx := r.Context()

	if xutil.IsEmpty(*param.Code) || xutil.IsEmpty(*param.RedirectUri) {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, nil)
		return
	}

	if err := t.revokeAuthCode(ctx, *param.Code, clientId); err != nil {
		t.respondAuthCodeError(w, r, err)
		return
	}

	if err := t.svc.ValidateRedirectUri(ctx, *param.RedirectUri, clientId, *param.Code); err != nil {
		if errors.Is(err, xerr.RedirectUriNotFound) || errors.Is(err, xerr.RedirectUriNotMatched) {
			RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		} else {
			RespondJSON500(w, r, err)
		}
		return
	}

	uid, ok := t.ctx.Read(ctx, typedef.UserIdKey{}).(typedef.UserID)
	if !ok {
		RespondJSON401(w, r, xerr.UnauthorizedRequest, nil, nil)
		return
	}

	tokens, err := t.generateToken(uid)
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	if err = t.sess.WriteRefreshTokenOwner(ctx, *tokens["RefreshToken"], clientId); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	resp := &TokenResponse{
		AccessToken:  *tokens["AccessToken"],
		RefreshToken: *tokens["RefreshToken"],
		IdToken:      tokens["IdToken"],
		TokenType:    config.BearerTokenType,
		ExpiresIn:    3600,
	}

	RespondJSON(w, r, http.StatusOK, resp)
}

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
		"AccessToken":  &accessToken,
		"RefreshToken": &refreshToken,
		"IdToken":      &idToken,
	}, nil
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

	uid, ok := t.ctx.Read(r.Context(), typedef.UserIdKey{}).(typedef.UserID)
	if !ok {
		RespondJSON401(w, r, xerr.UnauthorizedRequest, nil, nil)
		return
	}

	accessToken, err := t.svc.GenerateAccessToken(uid)
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	refreshToken, err := t.svc.GenerateRefreshToken(uid)
	if err != nil {
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
