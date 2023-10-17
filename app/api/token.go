package api

import (
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/go-playground/validator/v10"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
	"github.com/42milez/go-oidc-server/app/service"
)

var tokenHdlr *TokenHdlr

func NewTokenHdlr(option *HandlerOption) *TokenHdlr {
	return &TokenHdlr{
		svc:       service.NewToken(option.db, option.clock),
		tokenGen:  option.tokenGenerator,
		validator: option.validator,
	}
}

type TokenHdlr struct {
	svc       TokenRequestAcceptor
	tokenGen  service.TokenGenerator
	validator *validator.Validate
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

	if err = t.validator.Struct(param); err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	// TODO: Validate token request
	// ...

	if param.GrantType == config.AuthorizationCodeGrantType {
		t.handleAuthCodeGrantType(w, r, param, clientId)
		return
	} else if param.GrantType == config.RefreshTokenGrantType {
		// TODO: Generate Access Token if grant_type is refresh_token.
		t.handleRefreshTokenGrantType()
		return
	} else {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}
}

func (t *TokenHdlr) handleAuthCodeGrantType(w http.ResponseWriter, r *http.Request, param *TokenFormdataBody, clientId string) {
	ctx := r.Context()

	if xutil.IsEmpty(*param.Code) {
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

	if xutil.IsEmpty(*param.RedirectUri) {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, nil)
		return
	}

	if err := t.svc.ValidateRedirectUri(ctx, *param.RedirectUri, clientId); err != nil {
		if errors.Is(err, xerr.RedirectUriNotFound) {
			RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		} else {
			RespondJSON500(w, r, err)
		}
		return
	}

	accessToken, err := t.tokenGen.GenerateToken("")
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	refreshToken, err := t.tokenGen.GenerateToken("")
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	idToken, err := t.tokenGen.GenerateToken("")
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	resp := &TokenResponse{
		AccessToken:  string(accessToken),
		RefreshToken: string(refreshToken),
		IdToken:      string(idToken),
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

func (t *TokenHdlr) handleRefreshTokenGrantType() {

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

	return &TokenFormdataBody{
		Code:         &code,
		GrantType:    grantType,
		RedirectUri:  &redirectUri,
		RefreshToken: &refreshToken,
	}, nil
}
