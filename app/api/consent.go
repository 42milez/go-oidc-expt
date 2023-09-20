package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/service"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

type ConsentHdlr struct {
	service   ConsentAcceptor
	validator *validator.Validate
}

func (ch *ConsentHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	decoder := schema.NewDecoder()
	q := &AuthorizeParams{}

	if err := decoder.Decode(q, r.URL.Query()); err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	sess, ok := service.GetSession(ctx)

	if !ok {
		RespondJSON(w, http.StatusUnauthorized, &ErrorResponse{
			Status:  http.StatusUnauthorized,
			Summary: xerr.UnauthorizedRequest,
		})
		return
	}

	if err := ch.validator.Struct(q); err != nil {
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.InvalidRequest,
		})
		return
	}

	if err := ch.service.AcceptConsent(ctx, sess.UserID, q.ClientId); err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	Redirect(w, r, config.AuthorizationEndpoint, http.StatusFound)
}
