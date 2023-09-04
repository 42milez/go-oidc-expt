package main

import (
	"context"
	"fmt"
	"net/http"

	auth "github.com/42milez/go-oidc-server/app/pkg/xjwt"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/api"
	"github.com/42milez/go-oidc-server/app/api/cookie"
	"github.com/42milez/go-oidc-server/app/api/session"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/repository"
	httpSwagger "github.com/swaggo/http-swagger/v2"

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

	CheckHealthHdlr := api.NewCheckHealth(rc, dc)

	mux.HandleFunc("/health", CheckHealthHdlr.ServeHTTP)

	//  User Endpoint
	// --------------------------------------------------

	registerHdlr, err := api.NewRegisterUser(ec, xid.UID, sess)

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	authHdlr, err := api.NewAuthenticate(ec, rc, ck, jwt, sess)

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	mux.Route(makePattern("user"), func(r chi.Router) {
		r.Use(api.RestoreSession(ck, sess))
		r.Post("/register", registerHdlr.ServeHTTP)
		r.Post("/auth", authHdlr.ServeHTTP)
	})

	//  OpenID Endpoint
	// --------------------------------------------------

	authorizeGet, err := api.NewAuthorizeGet()

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	authorizePost := api.NewAuthorizePost()

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
