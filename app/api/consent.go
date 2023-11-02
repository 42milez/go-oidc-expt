package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/option"

	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/httpstore"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/gorilla/schema"
)

var consentHdlr *ConsentHdlr

func NewConsentHdlr(opt *option.Option) *ConsentHdlr {
	return &ConsentHdlr{
		svc:     service.NewConsent(opt),
		context: &httpstore.Context{},
		v:       opt.V,
	}
}

type ConsentHdlr struct {
	svc     ConsentAcceptor
	context iface.ContextReader
	v       iface.StructValidator
}

func (ch *ConsentHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	decoder := schema.NewDecoder()
	q := &AuthorizeParams{}

	if err := decoder.Decode(q, r.URL.Query()); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	if err := ch.v.Struct(q); err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	uid, ok := ch.context.Read(ctx, typedef.UserIdKey{}).(typedef.UserID)
	if !ok {
		RespondJSON401(w, r, xerr.UnauthorizedRequest, nil, nil)
		return
	}

	if err := ch.svc.AcceptConsent(ctx, uid, q.ClientID); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	Redirect(w, r, config.AuthorizationPath, http.StatusFound)
}
