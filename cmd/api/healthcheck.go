package api

import (
	"net/http"

	"github.com/42milez/go-oidc-expt/cmd/option"
	"github.com/42milez/go-oidc-expt/cmd/repository"
	"github.com/42milez/go-oidc-expt/cmd/service"
)

var healthCheck *HealthCheck

func InitHealthCheck(opt *option.Option) {
	if healthCheck == nil {
		healthCheck = &HealthCheck{
			svc: service.NewCheckHealth(repository.NewCheckHealth(opt.DB, opt.Cache)),
		}
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
