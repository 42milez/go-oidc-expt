package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/cmd/config"
	"github.com/42milez/go-oidc-server/cmd/httpstore"
	"github.com/42milez/go-oidc-server/cmd/iface"
	"github.com/42milez/go-oidc-server/cmd/option"
	"github.com/42milez/go-oidc-server/cmd/service"
	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/gorilla/schema"
)

var consent *Consent

func InitConsent(opt *option.Option) {
	if consent == nil {
		consent = &Consent{
			svc:     service.NewConsent(opt),
			context: &httpstore.Context{},
			v:       opt.V,
		}
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
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	sess, ok := ch.context.Read(ctx, httpstore.SessionKey{}).(*httpstore.Session)
	if !ok {
		RespondJSON401(w, r, xerr.InvalidRequest, nil, nil)
		return
	}

	if err := ch.svc.AcceptConsent(ctx, sess.UserID, q.ClientID); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	Redirect(w, r, config.AuthorizationPath(), http.StatusFound)
}
