package api

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"
	"time"

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
const requestTimeout = 30 * time.Second

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

	//  SERVICE
	// --------------------------------------------------

	ConfigureService(option)

	//  HANDLER
	// --------------------------------------------------

	if err = ConfigureHandler(option); err != nil {
		return nil, nil, err
	}

	//  ROUTER
	// --------------------------------------------------

	mux := chi.NewRouter()

	// Common Middleware Configuration

	mux.Use(middleware.RequestID)
	mux.Use(AccessLogger)
	mux.Use(middleware.Timeout(requestTimeout))
	mux.Use(middleware.Recoverer)

	// OpenAPI Validation Middleware

	var swag *openapi3.T

	if swag, err = oapigen.GetSwagger(); err != nil {
		return nil, nil, err
	}

	swag.Servers = nil

	mux.Use(chimw.OapiRequestValidatorWithOptions(swag, &chimw.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: NewOapiAuthentication(option.db),
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

func NewOapiAuthentication(db *datastore.Database) openapi3filter.AuthenticationFunc {
	svc := service.NewOapiAuthenticate(repository.NewRelyingParty(db))
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return oapiBasicAuthenticate(ctx, input, svc)
	}
}

func oapiBasicAuthenticate(ctx context.Context, input *openapi3filter.AuthenticationInput, svc CredentialValidator) error {
	if input.SecuritySchemeName != basicAuthSchemeName {
		return xerr.UnknownSecurityScheme
	}

	credentials, err := extractCredential(input.RequestValidationInput.Request)

	if err != nil {
		return err
	}

	clientID := credentials[0]
	clientSecret := credentials[1]

	if err = svc.ValidateCredential(ctx, clientID, clientSecret); err != nil {
		LogError(input.RequestValidationInput.Request, err, nil)
		return err
	}

	return nil
}

func extractCredential(r *http.Request) ([]string, error) {
	authHdr := r.Header.Get("Authorization")

	if xutil.IsEmpty(authHdr) {
		return nil, xerr.UnauthorizedRequest
	}

	credentialBase64 := strings.Replace(authHdr, "Basic ", "", -1)
	credentialDecoded, err := base64.StdEncoding.DecodeString(credentialBase64)

	if err != nil {
		return nil, xerr.UnexpectedErrorOccurred
	}

	credentials := strings.Split(string(credentialDecoded), ":")

	return credentials, nil
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
		clock:       &xtime.RealClocker{},
		cookie:      service.NewCookie(rawHashKey, rawBlockKey, xtime.RealClocker{}),
		idGenerator: idGen,
	}

	if option.tokenGenerator, err = NewJWT(xtime.RealClocker{}); err != nil {
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
