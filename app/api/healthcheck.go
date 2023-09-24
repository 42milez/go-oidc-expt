package api

import (
	"net/http"
)

type CheckHealthHdlr struct {
	service HealthChecker
}

func (ch *CheckHealthHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	errResp := func(err error) {
		RespondJSON503(w, err)
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

	RespondJSON200(w)
}
