package api

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/datastore"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"
	chimw "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type HandlerOption struct {
	cache           *datastore.Cache
	cookie          *service.Cookie
	db              *datastore.Database
	idGenerator     *xid.UniqueID
	jwtUtil         *JWT
	sessionCreator  *service.CreateSession
	sessionRestorer *service.RestoreSession
	validationUtil  *validator.Validate
}

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	var err error

	// --------------------------------------------------
	//  DATASTORE
	// --------------------------------------------------

	var db *datastore.Database
	var cache *datastore.Cache

	if db, err = datastore.NewDatabase(ctx, cfg); err != nil {
		return nil, nil, err
	}

	if cache, err = datastore.NewCache(ctx, cfg); err != nil {
		return nil, nil, err
	}

	// --------------------------------------------------
	//  HANDLER OPTION
	// --------------------------------------------------

	option := &HandlerOption{
		cache:           cache,
		cookie:          service.NewCookie(rawHashKey, rawBlockKey),
		db:              db,
		idGenerator:     xid.UID,
		sessionCreator:  service.NewCreateSession(cache),
		sessionRestorer: service.NewRestoreSession(cache),
		validationUtil:  validator.New(),
	}

	if option.jwtUtil, err = NewJWT(xtime.RealClocker{}, rawPrivateKey, rawPublicKey); err != nil {
		return nil, nil, err
	}

	// --------------------------------------------------
	//  HANDLER
	// --------------------------------------------------

	checkHealthHdlr = NewCheckHealthHdlr(option)
	registerUserHdlr = NewRegisterUserHdlr(option)
	authenticateUserHdlr = NewAuthenticateUserHdlr(option)

	// --------------------------------------------------
	//  ROUTER
	// --------------------------------------------------

	var mux *chi.Mux
	var mw *MiddlewareFuncMap
	var swag *openapi3.T

	mux = chi.NewRouter()

	if swag, err = GetSwagger(); err != nil {
		return nil, nil, err
	}

	swag.Servers = nil

	mux.Use(chimw.OapiRequestValidator(swag))

	mw = NewMiddlewareFuncMap()
	mw.SetAuthenticateUserMW([]MiddlewareFunc{
		//RestoreSession(ck, sess),
	})
	mw.SetRegisterUserMW([]MiddlewareFunc{
		//RestoreSession(ck, sess),
	})

	mux = MuxWithOptions(&HandlerImpl{}, &ChiServerOptions{
		BaseRouter:  mux,
		Middlewares: mw,
	})

	return mux, func() {
		xutil.CloseConnection(db.Client)
		xutil.CloseConnection(cache.Client)
	}, nil
}

func NewCheckHealthHdlr(option *HandlerOption) *CheckHealthHdlr {
	return &CheckHealthHdlr{
		service: service.NewCheckHealth(repository.NewCheckHealth(option.db, option.cache)),
	}
}

func NewRegisterUserHdlr(option *HandlerOption) *RegisterUserHdlr {
	return &RegisterUserHdlr{
		service:   service.NewCreateUser(repository.NewUser(option.db, option.idGenerator)),
		session:   option.sessionRestorer,
		validator: option.validationUtil,
	}
}

func NewAuthenticateUserHdlr(option *HandlerOption) *AuthenticateUserHdlr {
	return &AuthenticateUserHdlr{
		service:   service.NewAuthenticate(repository.NewUser(option.db, option.idGenerator), option.jwtUtil),
		cookie:    option.cookie,
		session:   option.sessionCreator,
		validator: option.validationUtil,
	}
}

//func NewConsent() (*Consent, error) {
//	return &Consent{
//		Session: nil,
//	}, nil
//}
