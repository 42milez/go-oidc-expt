package api

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5/middleware"

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
	"github.com/rs/zerolog"
)

const basicAuthSchemeName = "basicAuth"

func NewMux(ctx context.Context, cfg *config.Config, logger *zerolog.Logger) (http.Handler, func(), error) {
	var err error

	appLogger = logger.With().Str(config.LoggerTagKey, config.AppLoggerTagValue).Logger()
	accessLogger = logger.With().Str(config.LoggerTagKey, config.AccessLoggerTagValue).Logger()

	//  HANDLER OPTION
	// --------------------------------------------------

	var option *HandlerOption

	if option, err = NewHandlerOption(); err != nil {
		return nil, nil, err
	}

	//  DATASTORE
	// --------------------------------------------------

	if err = ConfigureDatastore(ctx, cfg, option); err != nil {
		return nil, nil, err
	}

	//  HANDLER
	// --------------------------------------------------

	if err = ConfigureHandler(option); err != nil {
		return nil, nil, err
	}

	//  SERVICE
	// --------------------------------------------------

	ConfigureService(option)

	//  ROUTER
	// --------------------------------------------------

	mux := chi.NewRouter()

	// Common Middleware Configuration

	//mwLogger := logger.With().Str(config.LoggerTagKey, config.AccessLoggerTagValue).Logger()
	//
	//mux.Use(httplog.RequestLogger(mwLogger))

	mux.Use(middleware.RequestID)
	mux.Use(AccessLogger)
	mux.Use(middleware.Recoverer)

	// OpenAPI Validation Middleware

	var swag *openapi3.T

	if swag, err = oapigen.GetSwagger(); err != nil {
		return nil, nil, err
	}

	swag.Servers = nil

	mux.Use(chimw.OapiRequestValidatorWithOptions(swag, &chimw.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: NewOapiAuthentication(),
		},
		ErrorHandler: NewOapiErrorHandler(),
	}))

	// Middleware Configuration on Each Handler

	mw := oapigen.NewMiddlewareFuncMap()
	rs := RestoreSession(option)

	mw.SetAuthenticateMW(rs).SetAuthorizeMW(rs).SetConsentMW(rs).SetRegisterMW(rs).SetTokenMW(rs)

	mux = oapigen.MuxWithOptions(&HandlerImpl{}, &oapigen.ChiServerOptions{
		BaseRouter:  mux,
		Middlewares: mw,
	})

	return mux, func() {
		xutil.CloseConnection(option.db.Client)
		xutil.CloseConnection(option.cache.Client)
	}, nil
}

func NewOapiAuthentication() openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return oapiAuthenticate(ctx, input)
	}
}

func oapiAuthenticate(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	if input.SecuritySchemeName != basicAuthSchemeName {
		return xerr.UnknownSecurityScheme
	}

	authHdr := input.RequestValidationInput.Request.Header.Get("Authentication")

	if xutil.IsEmpty(authHdr) {
		return xerr.UnauthorizedRequest
	}

	return nil
}

func NewOapiErrorHandler() chimw.ErrorHandler {
	return func(w http.ResponseWriter, message string, statusCode int) {
		switch statusCode {
		case http.StatusBadRequest:
			RespondJSON400(w, nil, xerr.InvalidRequest, nil, nil)
		case http.StatusUnauthorized:
			RespondJSON401(w, nil, xerr.UnauthorizedRequest, nil, nil)
		case http.StatusNotFound:
			RespondJSON404(w)
		default:
			RespondJSON500(w, nil, nil)
		}
	}
}

func NewHandlerOption() (*HandlerOption, error) {
	var idGen *xid.UniqueID
	var err error

	if idGen, err = xid.GetUniqueID(); err != nil {
		return nil, err
	}

	option := &HandlerOption{
		cookie:      service.NewCookie(rawHashKey, rawBlockKey, xtime.RealClocker{}),
		idGenerator: idGen,
	}

	if option.jwtUtil, err = NewJWT(xtime.RealClocker{}); err != nil {
		return nil, err
	}

	if option.validator, err = NewAuthorizeParamValidator(); err != nil {
		return nil, err
	}

	return option, nil
}

func ConfigureDatastore(ctx context.Context, cfg *config.Config, option *HandlerOption) error {
	var err error

	if option.db, err = datastore.NewDatabase(ctx, cfg); err != nil {
		return err
	}

	if option.cache, err = datastore.NewCache(ctx, cfg); err != nil {
		return err
	}

	return nil
}

func ConfigureHandler(option *HandlerOption) error {
	var err error

	checkHealthHdlr = NewCheckHealthHdlr(option)
	tokenHdlr = NewTokenHdlr(option)

	if authenticateUserHdlr, err = NewAuthenticateHdlr(option); err != nil {
		return err
	}

	if authorizeGetHdlr, err = NewAuthorizeGetHdlr(option); err != nil {
		return err
	}

	if consentHdlr, err = NewConsentHdlr(option); err != nil {
		return err
	}

	if registerUserHdlr, err = NewRegisterHdlr(option); err != nil {
		return err
	}

	return nil
}

func ConfigureService(option *HandlerOption) {
	option.sessionCreator = service.NewCreateSession(repository.NewSession(option.cache))
	option.sessionRestorer = service.NewRestoreSession(repository.NewSession(option.cache))
	option.sessionUpdater = service.NewUpdateSession(repository.NewSession(option.cache))
}
