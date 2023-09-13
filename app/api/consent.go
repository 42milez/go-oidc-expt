package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/entity"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/config"
)

type ConsentHdlr struct {
	session SessionUpdater
}

func (ch *ConsentHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var sessId typedef.SessionID
	var sess *entity.Session
	var ok bool

	if sessId, ok = service.GetSessionID(ctx); !ok {
		RespondJSON(w, http.StatusUnauthorized, &ErrorResponse{
			Detail: xerr.UnauthorizedRequest.Error(),
			Status: http.StatusUnauthorized,
		})
		return
	}

	if sess, ok = service.GetSession(ctx); !ok {
		RespondJSON(w, http.StatusUnauthorized, &ErrorResponse{
			Detail: xerr.UnauthorizedRequest.Error(),
			Status: http.StatusUnauthorized,
		})
		return
	}

	sess.Consent = true

	if err := ch.session.Update(ctx, sessId, sess); err != nil {
		RespondJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	Redirect(w, r, config.AuthorizationEndpoint, http.StatusFound)
}
