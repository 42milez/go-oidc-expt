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
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/app/idp/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	_ "github.com/42milez/go-oidc-server/app/idp/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	_ = validator.New()

	client, db, cleanup, err := repository.NewDB(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	mux := chi.NewRouter()

	//  API document
	// --------------------------------------------------

	swaggerURL := fmt.Sprintf("http://%s:%d/%s/doc.json", cfg.SwaggerHost, cfg.SwaggerPort, cfg.SwaggerPath)
	mux.HandleFunc("/swagger/*", httpSwagger.Handler(httpSwagger.URL(swaggerURL)))

	//  Health
	// --------------------------------------------------

	checkHealthHdlr := &handler.CheckHealth{
		Service: &service.CheckHealth{DB: db},
	}
	mux.HandleFunc("/health", checkHealthHdlr.ServeHTTP)

	//  Admin
	// --------------------------------------------------

	adminRepo := &repository.Admin{Clock: xutil.RealClocker{}, DB: client}
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
