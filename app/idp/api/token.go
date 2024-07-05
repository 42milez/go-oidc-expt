package api

import (
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/42milez/go-oidc-server/app/idp/httpstore"
	"github.com/42milez/go-oidc-server/app/idp/service"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
)

var token *Token

func InitToken(opt *option.Option) {
	if token == nil {
		token = &Token{
			acSVC: service.NewAuthCodeGrant(opt),
			rtSVC: service.NewRefreshTokenGrant(opt),
			cache: httpstore.NewCache(opt),
			v:     opt.V,
		}
	}
}

type Token struct {
	acSVC AuthCodeGrantAcceptor
	rtSVC RefreshTokenGrantAcceptor
	cache TokenCacheReadWriter
	v     iface.StructValidator
}

func (t *Token) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	credentials, err := extractCredential(r)
	if err != nil {
		RespondTokenRequestError(w, r, xerr.InvalidGrant)
		return
	}
	clientID := typedef.ClientID(credentials[0])

	param, err := t.parseForm(r)
	if err != nil {
		RespondServerError(w, r, err)
		return
	}

	if err = t.v.Struct(param); err != nil {
		RespondTokenRequestError(w, r, xerr.InvalidRequestOIDC)
		return
	}

	if param.GrantType == config.AuthorizationCodeGrantType {
		t.handleAuthCodeGrant(w, r, param, clientID)
		return
	} else if param.GrantType == config.RefreshTokenGrantType {
		t.handleRefreshTokenGrant(w, r, param, clientID)
		return
	} else {
		RespondTokenRequestError(w, r, xerr.InvalidRequestOIDC)
		return
	}
}

func (t *Token) handleAuthCodeGrant(w http.ResponseWriter, r *http.Request, param *TokenFormdataBody, clientID typedef.ClientID) {
	ctx := r.Context()

	if xutil.IsEmpty(*param.Code) || xutil.IsEmpty(*param.RedirectURI) {
		RespondTokenRequestError(w, r, xerr.InvalidRequestOIDC)
		return
	}

	if err := t.acSVC.RevokeAuthCode(ctx, *param.Code, clientID); err != nil {
		respondRevokeAuthCodeError(w, r, err)
		return
	}

	fp, err := t.cache.ReadAuthorizationRequestFingerprint(ctx, clientID, *param.Code)
	if err != nil {
		if errors.Is(err, xerr.UnauthorizedRequest) {
			RespondTokenRequestError(w, r, xerr.InvalidRequestOIDC)
		} else {
			RespondServerError(w, r, err)
		}
		return
	}

	if *param.RedirectURI != fp.RedirectURI {
		RespondTokenRequestError(w, r, xerr.InvalidGrant)
		return
	}

	tokenSet, err := t.generateTokens(fp, clientID)
	if err != nil {
		RespondServerError(w, r, err)
		return
	}

	if err = t.cache.WriteRefreshToken(ctx, tokenSet.RefreshToken, clientID, fp.UserID); err != nil {
		RespondServerError(w, r, err)
		return
	}

	respBody := &TokenResponse{
		AccessToken:  tokenSet.AccessToken,
		RefreshToken: tokenSet.RefreshToken,
		IDToken:      &tokenSet.IDToken,
		TokenType:    config.BearerTokenType,
		ExpiresIn:    3600,
	}

	headers := map[string]string{
		"Cache-Control": "no-store",
		"Pragma":        "no-cache",
	}

	RespondJSON(w, r, http.StatusOK, headers, respBody)
}

func respondRevokeAuthCodeError(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, xerr.AuthCodeExpired) {
		RespondTokenRequestError(w, r, xerr.InvalidGrant)
		return
	}

	if errors.Is(err, xerr.AuthCodeNotFound) {
		RespondTokenRequestError(w, r, xerr.InvalidGrant)
		return
	}

	if errors.Is(err, xerr.AuthCodeUsed) {
		RespondTokenRequestError(w, r, xerr.InvalidGrant)
		return
	}

	RespondServerError(w, r, err)
}

type TokenSet struct {
	AccessToken  string
	RefreshToken string
	IDToken      string
}

func (t *Token) generateTokens(param *typedef.AuthorizationRequestFingerprint, clientID typedef.ClientID) (*TokenSet, error) {
	accessToken, err := t.acSVC.GenerateAccessToken(param.UserID, nil)
	if err != nil {
		return nil, err
	}

	refreshToken, err := t.acSVC.GenerateRefreshToken(param.UserID, nil)
	if err != nil {
		return nil, err
	}

	audiences := []string{clientID.String()}
	idToken, err := t.acSVC.GenerateIDToken(param.UserID, audiences, param.AuthTime, param.Nonce)
	if err != nil {
		return nil, err
	}

	return &TokenSet{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IDToken:      idToken,
	}, nil
}

func (t *Token) handleRefreshTokenGrant(w http.ResponseWriter, r *http.Request, param *TokenFormdataBody, clientID typedef.ClientID) {
	ctx := r.Context()

	err := t.rtSVC.VerifyRefreshToken(ctx, *param.RefreshToken, clientID)
	if err != nil {
		if errors.Is(err, xerr.InvalidToken) || errors.Is(err, xerr.ClientIDNotMatched) {
			RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
			return
		} else if errors.Is(err, xerr.RefreshTokenNotFound) {
			RespondJSON401(w, r, xerr.InvalidRequest, nil, err)
			return
		} else {
			RespondJSON500(w, r, err)
			return
		}
	}

	uid, err := t.rtSVC.ExtractUserID(*param.RefreshToken)
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	accessToken, err := t.rtSVC.GenerateAccessToken(uid, nil)
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	refreshToken, err := t.rtSVC.GenerateRefreshToken(uid, nil)
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	if err = t.cache.WriteRefreshToken(ctx, refreshToken, clientID, uid); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	resp := &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    config.BearerTokenType,
		ExpiresIn:    3600,
	}

	RespondJSON(w, r, http.StatusOK, nil, resp)
}

func (t *Token) parseForm(r *http.Request) (*TokenFormdataBody, error) {
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
		GrantType: typedef.GrantType(grantType),
	}

	if !xutil.IsEmpty(code) {
		ret.Code = &code
	}

	if !xutil.IsEmpty(redirectUri) {
		ret.RedirectURI = &redirectUri
	}

	if !xutil.IsEmpty(refreshToken) {
		ret.RefreshToken = &refreshToken
	}

	return ret, nil
}
