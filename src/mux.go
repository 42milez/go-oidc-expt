package main

import (
	"context"
	"github.com/42milez/go-oidc-server/src/config"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	// TODO: Initialize validator
	// ...

	// TODO: Initialize store
	// ...

	// TODO: Create Repository
	// ...

	mux := chi.NewRouter()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, err := w.Write([]byte(`{"status":"ok"`))
		if err != nil {
			log.Error().Err(err).Send()
		}
	})

	// TODO: Add more endpoints
	// ...

	// TODO: Return actual cleanup function
	return mux, func() {}, nil
}
