package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/42milez/go-oidc-server/app/auth"
	"github.com/42milez/go-oidc-server/app/config"
	handler2 "github.com/42milez/go-oidc-server/app/handler"
	"github.com/42milez/go-oidc-server/app/repository"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	_ "github.com/42milez/go-oidc-server/app/docs"
	"github.com/go-chi/chi/v5"
)

const (
	apiVersionV1      = "v1"
	apiVersionCurrent = apiVersionV1
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

	jwtUtil, err := auth.NewUtil(&xutil.RealClocker{})

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	cookieUtil := handler2.NewCookie(cfg.CookieHashKey, cfg.CookieBlockKey)
	sessionUtil := handler2.NewSession(redisClient, jwtUtil)

	//  Swagger Endpoint
	// --------------------------------------------------

	if cfg.IsDevelopment() {
		mux.HandleFunc("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))
	}

	//  Health Check Endpoint
	// --------------------------------------------------

	CheckHealthHdlr := handler2.NewCheckHealth(redisClient, dbClient)

	mux.HandleFunc("/health", CheckHealthHdlr.ServeHTTP)

	//  User Endpoint
	// --------------------------------------------------

	createUserHdlr, err := handler2.NewCreateUser(entClient, sessionUtil, jwtUtil)

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	mux.Route(makePattern("create"), func(r chi.Router) {
		r.Post("/", createUserHdlr.ServeHTTP)
	})

	authenticateUserHdlr, err := handler2.NewAuthenticate(entClient, cookieUtil, sessionUtil, jwtUtil)

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	mux.Route(makePattern("authenticate"), func(r chi.Router) {
		r.Post("/", authenticateUserHdlr.ServeHTTP)
	})

	//  OpenID/OAuth Endpoint
	// --------------------------------------------------

	authorizeGet, err := handler2.NewAuthorizeGet()

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	authorizePost := handler2.NewAuthorizePost()

	mux.Route(makePattern("authorize"), func(r chi.Router) {
		r.Get("/", authorizeGet.ServeHTTP)
		r.Post("/", authorizePost.ServeHTTP)
	})

	return mux, func() {
		xutil.CloseConnection(entClient)
		xutil.CloseConnection(redisClient)
	}, nil
}

func makePattern(path string) string {
	return fmt.Sprintf("/%s/%s/%s", config.AppName, apiVersionCurrent, path)
}
