package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/config"
)

type ConsentHdlr struct {
	//Session SessionUpdater
}

func (ch *ConsentHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = service.GetSessionID(r.Context())

	// TODO: Save consent status into cache
	// ...

	Redirect(w, r, config.AuthorizationEndpoint, http.StatusFound)
}
