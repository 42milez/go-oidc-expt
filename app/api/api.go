package api

import (
	"context"
	"database/sql"
	chimw "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"net/http"

	"github.com/42milez/go-oidc-server/app/api/cookie"
	"github.com/42milez/go-oidc-server/app/api/session"
	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
	"github.com/42milez/go-oidc-server/app/pkg/xjwt"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen -config codegen/config.yml -o api.gen.go spec/spec.yml

var checkHealth *CheckHealth
var authUser *AuthenticateUser
var regUser *RegisterUser

type HandlerImpl struct{}

func (p *HandlerImpl) CheckHealth(w http.ResponseWriter, r *http.Request) {
	checkHealth.ServeHTTP(w, r)
}

func (p *HandlerImpl) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	authUser.ServeHTTP(w, r)
}

func (p *HandlerImpl) RegisterUser(w http.ResponseWriter, r *http.Request) {
	regUser.ServeHTTP(w, r)
}

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	var dc *sql.DB
	var ec *ent.Client
	var rc *redis.Client
	var err error

	if dc, err = repository.NewDBClient(ctx, cfg); err != nil {
		return nil, nil, err
	}

	ec = repository.NewEntClient(dc)

	if rc, err = repository.NewCacheClient(ctx, cfg); err != nil {
		return nil, nil, err
	}

	var ck *cookie.Cookie
	var sess *session.Session
	var jwt *xjwt.JWT

	ck = cookie.NewCookie(cfg.CookieHashKey, cfg.CookieBlockKey)
	sess = session.NewSession(rc, jwt)

	if jwt, err = xjwt.NewJWT(&xtime.RealClocker{}); err != nil {
		return nil, nil, err
	}

	checkHealth = NewCheckHealth(rc, dc)

	if regUser, err = NewRegisterUser(ec, xid.UID, sess); err != nil {
		return nil, nil, err
	}

	if authUser, err = NewAuthenticateUser(ec, rc, ck, jwt, sess); err != nil {
		return nil, nil, err
	}

	mux := chi.NewRouter()
	swag, err := GetSwagger()

	if err != nil {
		return nil, nil, err
	}

	swag.Servers = nil

	mux.Use(chimw.OapiRequestValidator(swag))

	mw := NewMiddlewareFuncMap()
	mw.SetAuthenticateUserMW([]MiddlewareFunc{
		RestoreSession(ck, sess),
	})
	mw.SetRegisterUserMW([]MiddlewareFunc{
		RestoreSession(ck, sess),
	})

	mux = MuxWithOptions(&HandlerImpl{}, &ChiServerOptions{
		BaseRouter:  mux,
		Middlewares: mw,
	})

	return mux, func() {
		xutil.CloseConnection(ec)
		xutil.CloseConnection(rc)
	}, nil
}
