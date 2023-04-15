package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ErrResponse struct {
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func RespondJSON(ctx context.Context, w http.ResponseWriter, body any, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Error().Err(err).Msg("failed to encode response")
		w.WriteHeader(http.StatusInternalServerError)
		resp := ErrResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Error().Err(err).Msg("failed to write response")
		}
		return
	}

	w.WriteHeader(status)

	if _, err := fmt.Fprintf(w, "%s", bodyBytes); err != nil {
		log.Error().Err(err).Msg("failed to write response")
	}
}
