package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/auth"
	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/handler"
	"github.com/42milez/go-oidc-server/app/idp/service"
	"github.com/42milez/go-oidc-server/app/idp/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	_ = validator.New()

	client, db, cleanup, err := store.NewDB(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	mux := chi.NewRouter()

	//  health
	// --------------------------------------------------

	checkHealthHdlr := &handler.CheckHealth{
		Service: &service.CheckHealth{DB: db},
	}
	mux.HandleFunc("/health", checkHealthHdlr.ServeHTTP)

	//  admin
	// --------------------------------------------------

	adminRepo := &store.AdminRepository{Clock: xutil.RealClocker{}, DB: client}
	jwtUtil, err := auth.NewJWTUtil(xutil.RealClocker{})

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	adminCreateHdlr := &handler.Create{
		Service: &service.AdminCreator{
			Repo: adminRepo,
		},
		Validator: validator.New(),
	}
	mux.Route("/admin/create", func(r chi.Router) {
		r.Post("/", adminCreateHdlr.ServeHTTP)
	})

	adminSignInHdlr := &handler.SignIn{
		Service: &service.AdminSignIn{
			Repo:           adminRepo,
			TokenGenerator: jwtUtil,
		},
		Validator: validator.New(),
	}
	mux.Route("/admin/signin", func(r chi.Router) {
		r.Post("/", adminSignInHdlr.ServeHTTP)
	})

	//readAdminHdlr := &handler.ReadAdmin{
	//	Service: &service.ReadAdmin{
	//		DB:   client,
	//		Repo: &repo,
	//	},
	//}
	//mux.Route("/admin", func(r chi.Router) {
	//	r.Get("/", readAdminHdlr.ServeHTTP)
	//})

	//readAdminsHdlr := &handler.ReadAdmins{
	//	Service: &service.ReadAdmins{
	//		DB:   client,
	//		Repo: &repo,
	//	},
	//}
	//mux.Route("/admins", func(r chi.Router) {
	//	r.Get("/", readAdminsHdlr.ServeHTTP)
	//})

	return mux, cleanup, nil
}
