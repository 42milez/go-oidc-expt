package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/httpstore"
	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/gorilla/schema"
)

var authorizeGetHdlr *AuthorizeGetHdlr

func NewAuthorizeGetHdlr(option *HandlerOption) (*AuthorizeGetHdlr, error) {
	return &AuthorizeGetHdlr{
		svc:  service.NewAuthorize(repository.NewRelyingParty(option.db)),
		ctx:  &httpstore.Context{},
		sess: httpstore.NewWriteSession(repository.NewSession(option.cache), &httpstore.Context{}, option.idGenerator),
		v:    option.validator,
	}, nil
}

type AuthorizeGetHdlr struct {
	svc  Authorizer
	ctx  iface.ContextReader
	sess iface.AuthParamSessionWriter
	v    iface.StructValidator
}

func (a *AuthorizeGetHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := schema.NewDecoder()
	q := &AuthorizeParams{}

	if err := decoder.Decode(q, r.URL.Query()); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	if err := a.v.Struct(q); err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	// TODO: Redirect unauthenticated user to the authentication endpoint with the posted parameters
	// ...

	// TODO: Redirect authenticated user to the consent endpoint with the posted parameters
	// ...

	location, authCode, err := a.svc.Authorize(r.Context(), q.ClientID, q.RedirectUri, q.State)
	if err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	ctx := r.Context()

	uid, ok := a.ctx.Read(ctx, typedef.UserIdKey{}).(typedef.UserID)
	if !ok {
		RespondJSON401(w, r, xerr.UnauthorizedRequest, nil, err)
		return
	}

	authParam := &typedef.AuthParam{
		RedirectUri: q.RedirectUri,
		UserId:      uid,
	}

	if err = a.sess.WriteAuthParam(ctx, authParam, q.ClientID, authCode); err != nil {
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
