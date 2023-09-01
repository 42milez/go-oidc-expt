package handler

import (
	"database/sql"
	"net/http"

	"github.com/42milez/go-oidc-server/app/model"

	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"

	"github.com/redis/go-redis/v9"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/rs/zerolog/log"
)

func NewCheckHealth(cacheClient *redis.Client, dbClient *sql.DB) *CheckHealth {
	return &CheckHealth{
		Service: &service.CheckHealth{
			Repo: &repository.CheckHealth{
				Cache: cacheClient,
				DB:    dbClient,
			},
		},
	}
}

type CheckHealth struct {
	Service HealthChecker
}

// ServeHTTP checks service condition
//
//	@summary		checks service condition
//	@description	This endpoint checks service condition
//	@id				CheckHealth.ServeHTTP
//	@tags			HealthCheck
//	@accept			json
//	@produce		json
//	@success		200	{object}	model.CheckHealthResponse
//	@failure		503	{object}	model.CheckHealthResponse
//	@router			/health [get]
func (p *CheckHealth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	errResp := func() {
		RespondJSON(w, http.StatusInternalServerError, &model.CheckHealthResponse{
			Status: http.StatusServiceUnavailable,
		})
	}
	ctx := r.Context()

	if err := p.Service.CheckCacheStatus(ctx); err != nil {
		log.Error().Err(err).Msgf("%s: %+v", xerr.FailedToPingCache, err)
		errResp()
		return
	}

	if err := p.Service.CheckDBStatus(ctx); err != nil {
		log.Error().Err(err).Msgf("%s: %+v", xerr.FailedToPingDatabase, err)
		errResp()
		return
	}

	respBody := model.CheckHealthResponse{
		Status: http.StatusOK,
	}

	RespondJSON(w, http.StatusOK, &respBody)
}
