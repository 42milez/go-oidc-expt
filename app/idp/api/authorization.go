package api

import (
	"errors"
	"net/http"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/httpstore"
	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/idp/service"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/gorilla/schema"
)

var authorizationGet *AuthorizationGet

func InitAuthorizationGet(opt *option.Option) {
	if authorizationGet == nil {
		authorizationGet = &AuthorizationGet{
			svc:     service.NewAuthorize(opt),
			context: &httpstore.Context{},
			v:       opt.V,
		}
	}
}

type AuthorizationGet struct {
	svc     Authorizer
	context iface.ContextReader
	v       iface.StructValidator
}

func (a *AuthorizationGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params, ok := a.context.Read(ctx, typedef.RequestParamKey{}).(*AuthorizeParams)
	if !ok {
		RespondServerError(w, r, xerr.TypeAssertionFailed)
	}
	if err := a.v.Struct(params); err != nil {
		RespondAuthorizationRequestError(w, r, params.RedirectURI, params.State, xerr.InvalidRequest)
		return
	}

	// TODO: Redirect unauthenticated user to the authentication endpoint with the posted parameters
	// ...

	// TODO: Redirect authenticated user to the consent endpoint with the posted parameters
	// ...

	location, authCode, err := a.svc.Authorize(r.Context(), params.ClientID, params.RedirectURI, params.State)
	if err != nil {
		if errors.Is(err, xerr.UserIDNotFoundInContext) {
			RespondAuthorizationRequestError(w, r, params.RedirectURI, params.State, xerr.AccessDenied)
		} else if errors.Is(err, xerr.InvalidRedirectURI) {
			RespondAuthorizationRequestError(w, r, params.RedirectURI, params.State, xerr.InvalidRequest)
		} else {
			RespondServerError(w, r, err)
		}
		return
	}

	if err := a.svc.SaveAuthorizationRequestFingerprint(ctx, params.ClientID, params.RedirectURI, params.Nonce, authCode); err != nil {
		if errors.Is(err, xerr.UserIDNotFoundInContext) {
			RespondAuthorizationRequestError(w, r, params.RedirectURI, params.State, xerr.AccessDenied)
		} else {
			RespondServerError(w, r, err)
		}
		return
	}

	http.Redirect(w, r, location.String(), http.StatusFound)
}

func NewAuthorizePost() *AuthorizePost {
	return &AuthorizePost{
		svc: nil,
		v:   nil,
	}
}

type AuthorizePost struct {
	svc Authorizer
	v   iface.StructValidator
}

func (p *AuthorizePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NOT IMPLEMENTED
}

func parseAuthorizeParam(r *http.Request, v iface.StructValidator) (*AuthorizeParams, error) {
	decoder := schema.NewDecoder()
	ret := &AuthorizeParams{}

	if err := decoder.Decode(ret, r.URL.Query()); err != nil {
		return nil, err
	}

	if err := v.Struct(ret); err != nil {
		return nil, xerr.FailedToValidate.Wrap(err)
	}

	return ret, nil
}
