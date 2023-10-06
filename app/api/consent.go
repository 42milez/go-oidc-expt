package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/repository"

	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/api/oapigen"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

var consentHdlr *ConsentHdlr

func NewConsentHdlr(option *HandlerOption) (*ConsentHdlr, error) {
	return &ConsentHdlr{
		service:   service.NewConsent(repository.NewUser(option.db, option.idGenerator)),
		validator: option.validator,
	}, nil
}

type ConsentHdlr struct {
	service   ConsentAcceptor
	validator *validator.Validate
}

func (c *ConsentHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	decoder := schema.NewDecoder()
	q := &oapigen.AuthorizeParams{}

	if err := decoder.Decode(q, r.URL.Query()); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	if err := c.validator.Struct(q); err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	sess, ok := service.GetSession(ctx)

	if !ok {
		RespondJSON401(w, r, xerr.UnauthorizedRequest, nil, nil)
		return
	}

	if err := c.service.AcceptConsent(ctx, sess.UserID, q.ClientId); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	Redirect(w, r, config.AuthorizationPath, http.StatusFound)
}
