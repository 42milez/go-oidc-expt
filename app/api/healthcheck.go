package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"
)

var checkHealthHdlr *CheckHealthHdlr

func NewCheckHealthHdlr(option *HandlerOption) *CheckHealthHdlr {
	return &CheckHealthHdlr{
		svc: service.NewCheckHealth(repository.NewCheckHealth(option.db, option.cache)),
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
