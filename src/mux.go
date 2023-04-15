package main

import (
	"context"
	"github.com/42milez/go-oidc-server/src/handler"
	"github.com/42milez/go-oidc-server/src/service"
	"net/http"

	"github.com/42milez/go-oidc-server/src/clock"
	"github.com/42milez/go-oidc-server/src/config"
	"github.com/42milez/go-oidc-server/src/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	_ = validator.New()

	_, db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	// TODO: Create Repository
	_ = store.Repository{Clocker: clock.RealClocker{}}

	mux := chi.NewRouter()

	checkHealthHandler := &handler.CheckHealth{
		Service: &service.CheckHealth{DB: db},
	}
	mux.HandleFunc("/health", checkHealthHandler.ServeHTTP)

	// TODO: Add more endpoints
	// ...

	return mux, cleanup, nil
}
