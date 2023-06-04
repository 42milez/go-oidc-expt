package main

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/config"
	handler2 "github.com/42milez/go-oidc-server/app/idp/handler"
	service2 "github.com/42milez/go-oidc-server/app/idp/service"
	"github.com/42milez/go-oidc-server/app/idp/store"
	"github.com/42milez/go-oidc-server/pkg/clock"

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

	checkHealthHdlr := &handler2.CheckHealth{
		Service: &service2.CheckHealth{DB: db},
	}
	mux.HandleFunc("/health", checkHealthHdlr.ServeHTTP)

	//  user
	// --------------------------------------------------

	readUserHdlr := &handler2.ReadUser{
		Service: &service2.ReadUser{
			DB:   client,
			Repo: &repo,
		},
	}
	mux.Route("/user", func(r chi.Router) {
		r.Get("/", readUserHdlr.ServeHTTP)
	})

	//  users
	// --------------------------------------------------

	readUsersHdlr := &handler2.ReadUsers{
		Service: &service2.ReadUsers{
			DB:   client,
			Repo: &repo,
		},
	}
	mux.Route("/users", func(r chi.Router) {
		r.Get("/", readUsersHdlr.ServeHTTP)
	})

	return mux, cleanup, nil
}
