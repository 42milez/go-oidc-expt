package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/httpstore"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/repository"

	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

var consentHdlr *ConsentHdlr

func NewConsentHdlr(option *HandlerOption) (*ConsentHdlr, error) {
	return &ConsentHdlr{
		service:   service.NewConsent(repository.NewUser(option.db, option.idGenerator)),
		rCtx:      &httpstore.ReadContext{},
		validator: option.validator,
	}, nil
}

type ConsentHdlr struct {
	service   ConsentAcceptor
	rCtx      ContextReader
	validator *validator.Validate
}

func (ch *ConsentHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	decoder := schema.NewDecoder()
	q := &AuthorizeParams{}

	if err := decoder.Decode(q, r.URL.Query()); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	if err := ch.validator.Struct(q); err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	uid, ok := ch.rCtx.Read(ctx, typedef.UserIdKey{}).(typedef.UserID)
	if !ok {
		RespondJSON401(w, r, xerr.UnauthorizedRequest, nil, nil)
		return
	}

	if err := ch.service.AcceptConsent(ctx, uid, q.ClientID); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	Redirect(w, r, config.AuthorizationPath, http.StatusFound)
}
