package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/app/idp/service"
)

var healthCheck *HealthCheck

func NewHealthCheck(opt *option.Option) *HealthCheck {
	return &HealthCheck{
		svc: service.NewCheckHealth(repository.NewCheckHealth(opt.DB, opt.Cache)),
	}
}

type HealthCheck struct {
	svc HealthChecker
}

func (c *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	errResp := func(err error) {
		RespondJSON503(w, r, err)
	}

	ctx := r.Context()

	if err := c.svc.CheckCacheStatus(ctx); err != nil {
		errResp(err)
		return
	}

	if err := c.svc.CheckDBStatus(ctx); err != nil {
		errResp(err)
		return
	}

	RespondJSON200(w, r)
}
