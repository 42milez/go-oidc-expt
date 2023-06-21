package main

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/idp/config"
	handler2 "github.com/42milez/go-oidc-server/app/idp/handler"
	service2 "github.com/42milez/go-oidc-server/app/idp/service"
	"github.com/42milez/go-oidc-server/app/idp/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	_ = validator.New()

	client, db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	repo := store.Repository{Clocker: xutil.RealClocker{}}

	mux := chi.NewRouter()

	//  health
	// --------------------------------------------------

	checkHealthHdlr := &handler2.CheckHealth{
		Service: &service2.CheckHealth{DB: db},
	}
	mux.HandleFunc("/health", checkHealthHdlr.ServeHTTP)

	//  admin
	// --------------------------------------------------

	readAdminHdlr := &handler2.ReadAdmin{
		Service: &service2.ReadAdmin{
			DB:   client,
			Repo: &repo,
		},
	}
	mux.Route("/admin", func(r chi.Router) {
		r.Get("/", readAdminHdlr.ServeHTTP)
	})

	//  admins
	// --------------------------------------------------

	readAdminsHdlr := &handler2.ReadAdmins{
		Service: &service2.ReadAdmins{
			DB:   client,
			Repo: &repo,
		},
	}
	mux.Route("/admins", func(r chi.Router) {
		r.Get("/", readAdminsHdlr.ServeHTTP)
	})

	return mux, cleanup, nil
}
