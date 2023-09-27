package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"
)

var checkHealthHdlr *CheckHealthHdlr

func NewCheckHealthHdlr(option *HandlerOption) *CheckHealthHdlr {
	return &CheckHealthHdlr{
		service: service.NewCheckHealth(repository.NewCheckHealth(option.db, option.cache)),
	}
}

type CheckHealthHdlr struct {
	service HealthChecker
}

func (ch *CheckHealthHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	errResp := func(err error) {
		RespondJSON503(w, r, err)
	}
	ctx := r.Context()

	if err := ch.service.CheckCacheStatus(ctx); err != nil {
		errResp(err)
		return
	}

	if err := ch.service.CheckDBStatus(ctx); err != nil {
		errResp(err)
		return
	}

	RespondJSON200(w, r)
}
