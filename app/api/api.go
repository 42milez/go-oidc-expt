package api

import (
	"context"
	"database/sql"
	"fmt"
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

const (
	apiVersionV1      = "v1"
	apiVersionCurrent = apiVersionV1
)

var checkHealthHdlr *CheckHealth
var authenticateUserHdlr *AuthenticateUser
var registerUserHdlr *RegisterUser

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

	checkHealthHdlr = NewCheckHealth(rc, dc)

	if registerUserHdlr, err = NewRegisterUser(ec, xid.UID, sess); err != nil {
		return nil, nil, err
	}

	if authenticateUserHdlr, err = NewAuthenticateUser(ec, rc, ck, jwt, sess); err != nil {
		return nil, nil, err
	}

	mw := NewMiddlewareFuncMap()
	mw.SetAuthenticateUserMW([]MiddlewareFunc{
		RestoreSession(ck, sess),
	})
	mw.SetRegisterUserMW([]MiddlewareFunc{
		RestoreSession(ck, sess),
	})

	mux := MuxWithOptions(&ServerInterfaceImpl{}, &ChiServerOptions{
		BaseURL:     fmt.Sprintf("/%s/%s", config.AppName, apiVersionCurrent),
		BaseRouter:  chi.NewRouter(),
		Middlewares: mw,
	})

	return mux, func() {
		xutil.CloseConnection(ec)
		xutil.CloseConnection(rc)
	}, nil
}

type ServerInterfaceImpl struct{}

func (p *ServerInterfaceImpl) CheckHealth(w http.ResponseWriter, r *http.Request) {
	checkHealthHdlr.ServeHTTP(w, r)
}

func (p *ServerInterfaceImpl) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	authenticateUserHdlr.ServeHTTP(w, r)
}

func (p *ServerInterfaceImpl) RegisterUser(w http.ResponseWriter, r *http.Request) {
	registerUserHdlr.ServeHTTP(w, r)
}
