package handler

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/session"
)

func NewConsent() (*Consent, error) {
	return &Consent{
		Session: nil,
	}, nil
}

type Consent struct {
	Session SessionUpdater
}

func (p *Consent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = session.GetSessionID(r.Context())

	// TODO: Save consent status into cache
	// ...

	Redirect(w, r, config.AuthorizationEndpoint, http.StatusFound)
}
