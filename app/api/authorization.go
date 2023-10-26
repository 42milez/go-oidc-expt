package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/httpstore"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

var authorizeGetHdlr *AuthorizeGetHdlr

func NewAuthorizeGetHdlr(option *HandlerOption) (*AuthorizeGetHdlr, error) {
	v, err := NewAuthorizeParamValidator()
	if err != nil {
		return nil, err
	}
	return &AuthorizeGetHdlr{
		service:   service.NewAuthorize(repository.NewRelyingParty(option.db)),
		validator: v,
		rCtx:      &httpstore.ReadContext{},
		sess:      option.SessionWriter,
	}, nil
}

type AuthorizeGetHdlr struct {
	service   Authorizer
	validator *validator.Validate
	rCtx      ContextReader
	sess      RedirectUriSessionWriter
}

func (a *AuthorizeGetHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := schema.NewDecoder()
	q := &AuthorizeParams{}

	if err := decoder.Decode(q, r.URL.Query()); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	if err := a.validator.Struct(q); err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	// TODO: Redirect unauthenticated user to the authentication endpoint with the posted parameters
	// ...

	// TODO: Redirect authenticated user to the consent endpoint with the posted parameters
	// ...

	location, err := a.service.Authorize(r.Context(), q.ClientID, q.RedirectUri, q.State)

	if err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	ctx := r.Context()

	sid, ok := a.rCtx.Read(ctx, typedef.SessionIdKey{}).(typedef.SessionID)
	if !ok {
		RespondJSON401(w, r, xerr.UnauthorizedRequest, nil, err)
		return
	}

	if err = a.sess.WriteRedirectUri(ctx, sid, q.RedirectUri); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	http.Redirect(w, r, location, http.StatusFound)
}

func NewAuthorizePost() *AuthorizePost {
	return &AuthorizePost{}
}

type AuthorizePost struct {
	Service   Authorizer
	Validator *validator.Validate
}

func (p *AuthorizePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NOT IMPLEMENTED YET
}

func parseAuthorizeParam(r *http.Request, v *validator.Validate) (*AuthorizeParams, error) {
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
