package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/42milez/go-oidc-server/app/api"
	"github.com/42milez/go-oidc-server/app/api/cookie"
	"github.com/42milez/go-oidc-server/app/api/session"
	"github.com/42milez/go-oidc-server/app/config"
	_ "github.com/42milez/go-oidc-server/app/docs"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
	auth "github.com/42milez/go-oidc-server/app/pkg/xjwt"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/go-chi/chi/v5"
)

const (
	apiVersionV1      = "v1"
	apiVersionCurrent = apiVersionV1
)

var checkHealthHdlr *api.CheckHealth
var authenticateUserHdlr *api.Authenticate
var registerUserHdlr *api.RegisterUser

type serverInterfaceImpl struct{}

func (p *serverInterfaceImpl) CheckHealth(w http.ResponseWriter, r *http.Request) {
	checkHealthHdlr.ServeHTTP(w, r)
}

func (p *serverInterfaceImpl) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	authenticateUserHdlr.ServeHTTP(w, r)
}

func (p *serverInterfaceImpl) RegisterUser(w http.ResponseWriter, r *http.Request) {
	registerUserHdlr.ServeHTTP(w, r)
}

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

	//if cfg.IsDevelopment() {
	//	mux.HandleFunc("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))
	//}

	//  Health Check Endpoint
	// --------------------------------------------------

	checkHealthHdlr = api.NewCheckHealth(rc, dc)

	//mux.HandleFunc("/health", CheckHealthHdlr.ServeHTTP)

	//  User Endpoint
	// --------------------------------------------------

	registerUserHdlr, err = api.NewRegisterUser(ec, xid.UID, sess)

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	authenticateUserHdlr, err = api.NewAuthenticate(ec, rc, ck, jwt, sess)

	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	//mux.Route(makePattern("user"), func(r chi.Router) {
	//	r.Use(RestoreSession(ck, sess))
	//	r.Post("/register", registerHdlr.ServeHTTP)
	//	r.Post("/auth", authHdlr.ServeHTTP)
	//})

	//  OpenID Endpoint
	// --------------------------------------------------

	//authorizeGet, err := NewAuthorizeGet()
	//
	//if err != nil {
	//	return nil, nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	//}
	//
	//authorizePost := NewAuthorizePost()
	//
	//mux.Route(makePattern("authorize"), func(r chi.Router) {
	//	r.Get("/", authorizeGet.ServeHTTP)
	//	r.Post("/", authorizePost.ServeHTTP)
	//})

	mw := api.NewMiddlewareFuncMap()
	mw.SetAuthenticateUserMW([]func(http.Handler) http.Handler{
		api.RestoreSession(ck, sess),
	})
	mw.SetRegisterUserMW([]func(http.Handler) http.Handler{
		api.RestoreSession(ck, sess),
	})

	mux = api.MuxWithOptions(&serverInterfaceImpl{}, &api.ChiServerOptions{
		BaseURL:     fmt.Sprintf("/%s/%s", config.AppName, apiVersionCurrent),
		BaseRouter:  mux,
		Middlewares: mw,
	})

	return mux, func() {
		xutil.CloseConnection(ec)
		xutil.CloseConnection(rc)
	}, nil
}
