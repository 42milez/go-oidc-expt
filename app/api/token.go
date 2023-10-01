package api

import (
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
	"github.com/42milez/go-oidc-server/app/service"
)

var tokenHdlr *TokenHdlr

func NewTokenHdlr(option *HandlerOption) *TokenHdlr {
	return &TokenHdlr{
		svc: service.NewToken(option.db),
	}
}

type TokenHdlr struct {
	svc TokenRequestAcceptor
}

func (th *TokenHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	credentials, err := extractCredential(r)
	if err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
	}
	clientId := credentials[0]

	params, err := th.parseForm(r)
	if err != nil {
		if errors.Is(err, xerr.MalformedFormParameter) {
			RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
			return
		}
		RespondJSON500(w, r, err)
		return
	}

	grantType := params.Get("grant_type")
	if xutil.IsEmpty(grantType) {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	if grantType == config.AuthorizationCodeGrantType {
		th.handleAuthCodeGrantType(w, r, params, clientId)
		return
	} else if grantType == config.RefreshTokenGrantType {
		// TODO: Generate Access Token if grant_type is refresh_token.
		th.handleRefreshTokenGrantType()
		return
	} else {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}
}

func (th *TokenHdlr) handleAuthCodeGrantType(w http.ResponseWriter, r *http.Request, params url.Values, clientId string) {
	ctx := r.Context()

	//code := params.Get("code")
	//if xutil.IsEmpty(code) {
	//	RespondJSON400(w, r, xerr.InvalidRequest, nil, nil)
	//	return
	//}
	//
	//if err := th.svc.ValidateAuthCode(ctx, code, clientId); err != nil {
	//	if errors.Is(err, xerr.AuthCodeNotFound) {
	//		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
	//	} else if errors.Is(err, xerr.AuthCodeExpired) {
	//		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
	//	} else if errors.Is(err, xerr.AuthCodeUsed) {
	//		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
	//	} else {
	//		RespondJSON500(w, r, err)
	//	}
	//	return
	//}

	uri := params.Get("redirect_uri")
	if xutil.IsEmpty(uri) {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, nil)
		return
	}

	if err := th.svc.ValidateRedirectUri(ctx, uri, clientId); err != nil {
		if errors.Is(err, xerr.RedirectUriNotFound) {
			RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		} else {
			RespondJSON500(w, r, err)
		}
		return
	}

	// TODO: Generate Access Token, Refresh token and ID Token if grant_type is authorization_code.
	// ...
}

func (th *TokenHdlr) handleRefreshTokenGrantType() {

}

func (th *TokenHdlr) parseForm(r *http.Request) (url.Values, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	params, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, xerr.MalformedFormParameter
	}

	return params, nil
}
