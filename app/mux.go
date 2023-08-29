package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/42milez/go-oidc-server/app/cookie"
	"github.com/42milez/go-oidc-server/app/middleware"

	"github.com/42milez/go-oidc-server/app/session"

	"github.com/42milez/go-oidc-server/pkg/xid"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/pkg/xtime"

	"github.com/42milez/go-oidc-server/app/auth"
	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/handler"
	"github.com/42milez/go-oidc-server/app/repository"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	_ "github.com/42milez/go-oidc-server/app/docs"
	"github.com/go-chi/chi/v5"
)

const (
	apiVersionV1      = "v1"
	apiVersionCurrent = apiVersionV1
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	dc, err := repository.NewDBClient(ctx, cfg)

	if err != nil {
		return nil, nil, xerr.FailedToInitialize.Wrap(err)
	}

	ec := repository.NewEntClient(dc)
	rc, err := repository.NewCacheClient(ctx, cfg)

	if err != nil {
		return nil, nil, xerr.FailedToInitialize.Wrap(err)
	}

	ck := cookie.NewCookie(cfg.CookieHashKey, cfg.CookieBlockKey)

	if err != nil {
		return nil, nil, xerr.FailedToInitialize.Wrap(err)
	}

	jwt, err := auth.NewJWT(&xtime.RealClocker{})

	if err != nil {
		return nil, nil, xerr.FailedToInitialize.Wrap(err)
	}

	sess := session.NewSession(rc, jwt)

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

	CheckHealthHdlr := handler.NewCheckHealth(rc, dc)

	mux.HandleFunc("/health", CheckHealthHdlr.ServeHTTP)

	//  Register Endpoint
	// --------------------------------------------------

	createUserHdlr, err := handler.NewCreateUser(ec, rc, xid.UID, jwt, sess)

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	mux.Route(makePattern("register"), func(r chi.Router) {
		r.Post("/", createUserHdlr.ServeHTTP)
	})

	//  Authentication Endpoint
	// --------------------------------------------------

	authenticateUserHdlr, err := handler.NewAuthenticate(ec, rc, ck, jwt, sess)

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	mux.Route(makePattern("auth"), func(r chi.Router) {
		r.Use(middleware.RestoreSession(ck, sess))
		r.Post("/", authenticateUserHdlr.ServeHTTP)
	})

	//  OpenID Endpoint
	// --------------------------------------------------

	authorizeGet, err := handler.NewAuthorizeGet()

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	authorizePost := handler.NewAuthorizePost()

	mux.Route(makePattern("authorize"), func(r chi.Router) {
		r.Get("/", authorizeGet.ServeHTTP)
		r.Post("/", authorizePost.ServeHTTP)
	})

	return mux, func() {
		xutil.CloseConnection(ec)
		xutil.CloseConnection(rc)
	}, nil
}

func makePattern(path string) string {
	return fmt.Sprintf("/%s/%s/%s", config.AppName, apiVersionCurrent, path)
}
