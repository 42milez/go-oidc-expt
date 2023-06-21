package handler

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

type CheckHealth struct {
	Service CheckHealthService
}

func (p *CheckHealth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := p.Service.PingCache(ctx); err != nil {
		log.Error().Err(err).Msg("failed to ping cache storage")
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	if err := p.Service.PingDB(ctx); err != nil {
		log.Error().Err(err).Msg("failed to ping database")
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	body := struct {
		StatusCode string `json:"statusCode"`
	}{
		http.StatusText(http.StatusOK),
	}
	RespondJSON(ctx, w, &body, http.StatusOK)
}
