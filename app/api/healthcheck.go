package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/api/oapigen"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/rs/zerolog/log"
)

type CheckHealthHdlr struct {
	service HealthChecker
}

func (ch *CheckHealthHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	errResp := func() {
		RespondJSON(w, http.StatusServiceUnavailable, &oapigen.Health{
			Status: http.StatusServiceUnavailable,
		})
	}
	ctx := r.Context()

	if err := ch.service.CheckCacheStatus(ctx); err != nil {
		log.Error().Err(err).Msg(xerr.FailedToPingCache.Error())
		errResp()
		return
	}

	if err := ch.service.CheckDBStatus(ctx); err != nil {
		log.Error().Err(err).Msg(xerr.FailedToPingDatabase.Error())
		errResp()
		return
	}

	respBody := oapigen.Health{
		Status: http.StatusOK,
	}

	RespondJSON(w, http.StatusOK, &respBody)
}
