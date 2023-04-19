package main

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/src/handler"
	"github.com/42milez/go-oidc-server/src/service"

	"github.com/42milez/go-oidc-server/src/clock"
	"github.com/42milez/go-oidc-server/src/config"
	"github.com/42milez/go-oidc-server/src/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	_ = validator.New()

	client, db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	repo := store.Repository{Clocker: clock.RealClocker{}}

	mux := chi.NewRouter()

	//  health
	// --------------------------------------------------

	checkHealthHdlr := &handler.CheckHealth{
		Service: &service.CheckHealth{DB: db},
	}
	mux.HandleFunc("/health", checkHealthHdlr.ServeHTTP)

	//  user
	// --------------------------------------------------

	readUserHdlr := &handler.ReadUser{
		Service: &service.ReadUser{
			DB:   client,
			Repo: &repo,
		},
	}
	mux.Route("/user", func(r chi.Router) {
		r.Get("/", readUserHdlr.ServeHTTP)
	})

	//  users
	// --------------------------------------------------

	readUsersHdlr := &handler.ReadUsers{
		Service: &service.ReadUsers{
			DB:   client,
			Repo: &repo,
		},
	}
	mux.Route("/users", func(r chi.Router) {
		r.Get("/", readUsersHdlr.ServeHTTP)
	})

	return mux, cleanup, nil
}
