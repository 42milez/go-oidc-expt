package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/rs/zerolog/log"
)

type CheckHealthHdlr struct {
	service HealthChecker
}

func (p *CheckHealthHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	errResp := func() {
		RespondJSON(w, http.StatusInternalServerError, &Health{
			Status: http.StatusServiceUnavailable,
		})
	}
	ctx := r.Context()

	if err := p.service.CheckCacheStatus(ctx); err != nil {
		log.Error().Err(err).Msgf("%s: %+v", xerr.FailedToPingCache, err)
		errResp()
		return
	}

	if err := p.service.CheckDBStatus(ctx); err != nil {
		log.Error().Err(err).Msgf("%s: %+v", xerr.FailedToPingDatabase, err)
		errResp()
		return
	}

	respBody := Health{
		Status: http.StatusOK,
	}

	RespondJSON(w, http.StatusOK, &respBody)
}
