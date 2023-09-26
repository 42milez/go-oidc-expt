package api

import (
	"context"
	_ "embed"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/go-chi/httplog"

	"github.com/42milez/go-oidc-server/app/api/oapigen"
	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"
	chimw "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen -config oapigen/config.yml -o oapigen/api.gen.go spec/spec.yml

//go:embed secret/key/block.key
var rawBlockKey []byte

//go:embed secret/key/hash.key
var rawHashKey []byte

var appLogger zerolog.Logger

var authenticateUserHdlr *AuthenticateHdlr
var authorizeGetHdlr *AuthorizeGetHdlr
var checkHealthHdlr *CheckHealthHdlr
var consentHdlr *ConsentHdlr
var registerUserHdlr *RegisterHdlr
var tokenHdlr *TokenHdlr

type HandlerImpl struct{}

func (_ *HandlerImpl) Authenticate(w http.ResponseWriter, r *http.Request) {
	authenticateUserHdlr.ServeHTTP(w, r)
}

func (_ *HandlerImpl) Authorize(w http.ResponseWriter, r *http.Request) {
	authorizeGetHdlr.ServeHTTP(w, r)
}

func (_ *HandlerImpl) CheckHealth(w http.ResponseWriter, r *http.Request) {
	checkHealthHdlr.ServeHTTP(w, r)
}

func (_ *HandlerImpl) Consent(w http.ResponseWriter, r *http.Request) {
	consentHdlr.ServeHTTP(w, r)
}

func (_ *HandlerImpl) Register(w http.ResponseWriter, r *http.Request) {
	registerUserHdlr.ServeHTTP(w, r)
}

func (_ *HandlerImpl) Token(w http.ResponseWriter, r *http.Request) {
	tokenHdlr.ServeHTTP(w, r)
}

type HandlerOption struct {
	cache           *datastore.Cache
	cookie          *service.Cookie
	db              *datastore.Database
	idGenerator     *xid.UniqueID
	jwtUtil         *JWT
	sessionCreator  *service.CreateSession
	sessionRestorer *service.RestoreSession
	sessionUpdater  *service.UpdateSession
}

func NewMux(ctx context.Context, cfg *config.Config, logger *zerolog.Logger) (http.Handler, func(), error) {
	var err error

	appLogger = logger.With().Str(config.LoggerTagKey, config.AppLoggerTagValue).Logger()

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

	//  HANDLER OPTION
	// --------------------------------------------------

	var idGen *xid.UniqueID

	if idGen, err = xid.GetUniqueID(); err != nil {
		return nil, nil, err
	}

	option := &HandlerOption{
		cache:           cache,
		cookie:          service.NewCookie(rawHashKey, rawBlockKey, xtime.RealClocker{}),
		db:              db,
		idGenerator:     idGen,
		sessionCreator:  service.NewCreateSession(repository.NewSession(cache)),
		sessionRestorer: service.NewRestoreSession(repository.NewSession(cache)),
		sessionUpdater:  service.NewUpdateSession(repository.NewSession(cache)),
	}

	if option.jwtUtil, err = NewJWT(xtime.RealClocker{}); err != nil {
		return nil, nil, err
	}

	//  HANDLER
	// --------------------------------------------------

	checkHealthHdlr = NewCheckHealthHdlr(option)
	tokenHdlr = NewTokenHdlr(option)

	if authenticateUserHdlr, err = NewAuthenticateHdlr(option); err != nil {
		return nil, nil, err
	}

	if authorizeGetHdlr, err = NewAuthorizeGetHdlr(option); err != nil {
		return nil, nil, err
	}

	if consentHdlr, err = NewConsentHdlr(option); err != nil {
		return nil, nil, err
	}

	if registerUserHdlr, err = NewRegisterHdlr(option); err != nil {
		return nil, nil, err
	}

	//  ROUTER
	// --------------------------------------------------

	var mux *chi.Mux
	var mw *oapigen.MiddlewareFuncMap
	var swag *openapi3.T

	mux = chi.NewRouter()

	// Swagger

	if swag, err = oapigen.GetSwagger(); err != nil {
		return nil, nil, err
	}

	swag.Servers = nil

	mux.Use(chimw.OapiRequestValidator(swag))

	// Common Middleware Configuration

	mwLogger := logger.With().Str(config.LoggerTagKey, config.MWLoggerTagValue).Logger()

	mux.Use(httplog.RequestLogger(mwLogger))

	// Middleware Configuration on Each Handler

	mw = oapigen.NewMiddlewareFuncMap()

	rs := RestoreSession(option)
	mw.SetAuthenticateMW(rs).SetAuthorizeMW(rs).SetConsentMW(rs).SetRegisterMW(rs).SetTokenMW(rs)

	mux = oapigen.MuxWithOptions(&HandlerImpl{}, &oapigen.ChiServerOptions{
		BaseRouter:  mux,
		Middlewares: mw,
	})

	return mux, func() {
		xutil.CloseConnection(db.Client)
		xutil.CloseConnection(cache.Client)
	}, nil
}

func NewJWT(clock xtime.Clocker) (*JWT, error) {
	parseKey := func(key []byte) (jwk.Key, error) {
		return jwk.ParseKey(key, jwk.WithPEM(true))
	}

	ret := &JWT{
		clock: clock,
	}

	var err error

	if ret.privateKey, err = parseKey(rawPrivateKey); err != nil {
		return nil, err
	}

	if ret.publicKey, err = parseKey(rawPublicKey); err != nil {
		return nil, err
	}

	return ret, nil
}

func NewAuthenticateHdlr(option *HandlerOption) (*AuthenticateHdlr, error) {
	v, err := NewAuthorizeParamValidator()
	if err != nil {
		return nil, err
	}
	return &AuthenticateHdlr{
		service:   service.NewAuthenticate(repository.NewUser(option.db, option.idGenerator), option.jwtUtil),
		cookie:    option.cookie,
		session:   option.sessionCreator,
		validator: v,
	}, nil
}

func NewAuthorizeGetHdlr(option *HandlerOption) (*AuthorizeGetHdlr, error) {
	v, err := NewAuthorizeParamValidator()
	if err != nil {
		return nil, err
	}
	return &AuthorizeGetHdlr{
		service:   service.NewAuthorize(repository.NewRelyingParty(option.db)),
		validator: v,
	}, nil
}

func NewCheckHealthHdlr(option *HandlerOption) *CheckHealthHdlr {
	return &CheckHealthHdlr{
		service: service.NewCheckHealth(repository.NewCheckHealth(option.db, option.cache)),
	}
}

func NewConsentHdlr(option *HandlerOption) (*ConsentHdlr, error) {
	v, err := NewAuthorizeParamValidator()
	if err != nil {
		return nil, err
	}
	return &ConsentHdlr{
		service:   service.NewConsent(repository.NewUser(option.db, option.idGenerator)),
		validator: v,
	}, nil
}

func NewRegisterHdlr(option *HandlerOption) (*RegisterHdlr, error) {
	v, err := NewAuthorizeParamValidator()
	if err != nil {
		return nil, err
	}
	return &RegisterHdlr{
		service:   service.NewCreateUser(repository.NewUser(option.db, option.idGenerator)),
		session:   option.sessionRestorer,
		validator: v,
	}, nil
}

func NewTokenHdlr(option *HandlerOption) *TokenHdlr {
	return &TokenHdlr{}
}
