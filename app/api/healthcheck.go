package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/option"

	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"
)

var checkHealthHdlr *CheckHealthHdlr

func NewCheckHealthHdlr(opt *option.Option) *CheckHealthHdlr {
	return &CheckHealthHdlr{
		svc: service.NewCheckHealth(repository.NewCheckHealth(opt.DB, opt.Cache)),
	}
}

type CheckHealthHdlr struct {
	svc HealthChecker
}

func (c *CheckHealthHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
