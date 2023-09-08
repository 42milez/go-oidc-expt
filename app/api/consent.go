package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/config"
)

type Consent struct {
	//Session SessionUpdater
}

func (p *Consent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = service.GetSessionID(r.Context())

	// TODO: Save consent status into cache
	// ...

	Redirect(w, r, config.AuthorizationEndpoint, http.StatusFound)
}
