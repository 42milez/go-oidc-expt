package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/repository"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/idp/config"
	_ "github.com/42milez/go-oidc-server/app/idp/docs"
	"github.com/42milez/go-oidc-server/app/idp/handler"
	"github.com/go-chi/chi/v5"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()

	// ==================================================
	//  Database and Ephemeral Store
	// ==================================================

	dbClient, err := repository.NewDBClient(ctx, cfg)

	if err != nil {
		return nil, nil, err
	}

	entClient := repository.NewEntClient(dbClient)
	redisClient, err := repository.NewCacheClient(ctx, cfg)

	if err != nil {
		return nil, nil, err
	}

	// ==================================================
	//  Route
	// ==================================================

	//  Swagger Endpoint
	// --------------------------------------------------

	if cfg.IsDevelopment() {
		mux.HandleFunc("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))
	}

	//  Health Check Endpoint
	// --------------------------------------------------

	CheckHealthHdlr := handler.NewCheckHealth(redisClient, dbClient)

	mux.HandleFunc("/health", CheckHealthHdlr.ServeHTTP)

	//  User Endpoint
	// --------------------------------------------------

	createUserHdlr, err := handler.NewCreateUser(entClient, redisClient)

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	mux.Route("/create", func(r chi.Router) {
		r.Post("/", createUserHdlr.ServeHTTP)
	})

	authenticateUserHdlr, err := handler.NewAuthenticate(entClient)

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	mux.Route("/auth", func(r chi.Router) {
		r.Post("/", authenticateUserHdlr.ServeHTTP)
	})

	//  OpenID/OAuth Endpoint
	// --------------------------------------------------

	authorizeGet, err := handler.NewAuthorizeGet()

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	authorizePost := handler.NewAuthorizePost()

	mux.Route("/idp/v1/authorize", func(r chi.Router) {
		r.Get("/", authorizeGet.ServeHTTP)
		r.Post("/", authorizePost.ServeHTTP)
	})

	return mux, func() {
		xutil.CloseConnection(entClient)
		xutil.CloseConnection(redisClient)
	}, nil
}
