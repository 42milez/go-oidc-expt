package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/option"

	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/httpstore"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/gorilla/schema"
)

var authorizationGet *AuthorizationGet

func NewAuthorizationGet(opt *option.Option) *AuthorizationGet {
	return &AuthorizationGet{
		svc:     service.NewAuthorize(opt),
		cache:   httpstore.NewCache(opt),
		context: &httpstore.Context{},
		v:       opt.V,
	}
}

type AuthorizationGet struct {
	svc     Authorizer
	cache   iface.OpenIdParamWriter
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
		RespondJSON400(w, r, xerr.InvalidRequest2, nil, err)
		return
	}

	// TODO: Redirect unauthenticated user to the authentication endpoint with the posted parameters
	// ...

	// TODO: Redirect authenticated user to the consent endpoint with the posted parameters
	// ...

	location, authCode, err := a.svc.Authorize(r.Context(), params.ClientID, params.RedirectUri, params.State)
	if err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest2, nil, err)
		return
	}

	uid, ok := a.context.Read(ctx, typedef.UserIdKey{}).(typedef.UserID)
	if !ok {
		RespondJSON401(w, r, xerr.InvalidRequest2, nil, err)
		return
	}

	authParam := &typedef.OpenIdParam{
		RedirectURI: params.RedirectUri,
		UserId:      uid,
	}

	if err = a.cache.WriteOpenIdParam(ctx, authParam, params.ClientID, authCode); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	http.Redirect(w, r, location, http.StatusFound)
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
