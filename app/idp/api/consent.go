package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/httpstore"
	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/idp/service"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/gorilla/schema"
)

var consent *Consent

func NewConsent(opt *option.Option) *Consent {
	return &Consent{
		svc:     service.NewConsent(opt),
		context: &httpstore.Context{},
		v:       opt.V,
	}
}

type Consent struct {
	svc     ConsentAcceptor
	context iface.ContextReader
	v       iface.StructValidator
}

func (ch *Consent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	decoder := schema.NewDecoder()
	q := &AuthorizeParams{}

	if err := decoder.Decode(q, r.URL.Query()); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	if err := ch.v.Struct(q); err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest2, nil, err)
		return
	}

	uid, ok := ch.context.Read(ctx, typedef.UserIdKey{}).(typedef.UserID)
	if !ok {
		RespondJSON401(w, r, xerr.InvalidRequest2, nil, nil)
		return
	}

	if err := ch.svc.AcceptConsent(ctx, uid, q.ClientID); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	Redirect2(w, r, config.AuthorizationPath, http.StatusFound)
}