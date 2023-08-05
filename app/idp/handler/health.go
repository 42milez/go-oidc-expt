package handler

import (
	"database/sql"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/service"
	"github.com/redis/go-redis/v9"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/rs/zerolog/log"
)

func NewCheckHealth(cacheClient *redis.Client, dbClient *sql.DB) *CheckHealth {
	return &CheckHealth{
		Service: &service.CheckHealth{
			Cache: cacheClient,
			DB:    dbClient,
		},
	}
}

type CheckHealth struct {
	Service HealthChecker
}

func (p *CheckHealth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := p.Service.PingCache(ctx); err != nil {
		log.Error().Err(err).Msgf("%s: %+v", xerr.FailedToPingCache, err)
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.ServiceCurrentlyUnavailable,
		})
		return
	}

	if err := p.Service.PingDB(ctx); err != nil {
		log.Error().Err(err).Msgf("%s: %+v", xerr.FailedToPingDatabase, err)
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.ServiceCurrentlyUnavailable,
		})
		return
	}

	body := struct {
		StatusCode string `json:"statusCode"`
	}{
		http.StatusText(http.StatusOK),
	}

	RespondJSON(w, http.StatusOK, &body)
}
