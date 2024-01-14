package api

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/datastore"
	"github.com/42milez/go-oidc-server/app/idp/httpstore"
	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/app/idp/security"
	"github.com/42milez/go-oidc-server/app/idp/service"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
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

	var opt *option.Option

	if opt, err = NewOption(); err != nil {
		return nil, nil, err
	}

	//  DATASTORE
	// --------------------------------------------------

	if err = ConfigureDatastore(ctx, cfg, opt); err != nil {
		return nil, nil, err
	}

	//  HANDLER
	// --------------------------------------------------

	InitHandler(opt)

	//  ROUTER
	// --------------------------------------------------

	mux := chi.NewRouter()

	// Common Middleware Configuration

	mux.Use(middleware.RequestID)
	mux.Use(AccessLogger)
	mux.Use(middleware.Timeout(requestTimeout))
	mux.Use(middleware.Recoverer)

	// Enable debugging

	if cfg.EnableProfiler {
		mux.Mount("/debug", middleware.Profiler())
	}

	// Middleware Configuration on Each Handler

	mw := NewMiddlewareFuncMap()
	restoreSessMW := RestoreSession(opt)

	mw.SetAuthenticateMW(restoreSessMW).SetAuthorizeMW(restoreSessMW).SetConsentMW(restoreSessMW).
		SetRegisterMW(restoreSessMW)

	mux, err = MuxWithOptions(&HandlerImpl{}, &ChiServerOptions{
		BaseRouter:  mux,
		Middlewares: mw,
	}, opt)

	if err != nil {
		return nil, nil, errors.New("failed to setup router")
	}

	return mux, func() {
		xutil.CloseConnection(opt.DB.Client)
		xutil.CloseConnection(opt.Cache.Client)
	}, nil
}

func NewOapiAuthentication(opt *option.Option) openapi3filter.AuthenticationFunc {
	svc := service.NewOapiAuthenticate(repository.NewRelyingParty(opt.DB))
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

	clientID := typedef.ClientID(credentials[0])
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
		return nil, xerr.CredentialNotFoundInHeader
	}

	credentialBase64 := strings.Replace(authHdr, "Basic ", "", -1)
	credentialDecoded, err := base64.StdEncoding.DecodeString(credentialBase64)
	if err != nil {
		return nil, err
	}

	credentials := strings.Split(string(credentialDecoded), ":")

	return credentials, nil
}

func NewOapiErrorHandler() nethttpmiddleware.ErrorHandler {
	return func(w http.ResponseWriter, message string, statusCode int) {
		switch statusCode {
		case http.StatusBadRequest:
			RespondJSON400(w, nil, xerr.InvalidRequest2, nil, nil)
		case http.StatusUnauthorized:
			RespondTokenRequestError(w, nil, xerr.InvalidClient)
		case http.StatusNotFound:
			RespondJSON404(w)
		default:
			RespondJSON500(w, nil, nil)
		}
	}
}

func NewOption() (*option.Option, error) {
	var err error

	opt := &option.Option{
		Cookie: httpstore.NewCookie(security.RawCookieHashKey, security.RawCookieBlockKey, &xtime.RealClocker{}),
	}

	if opt.IDGen, err = xid.GetUniqueIDGenerator(); err != nil {
		return nil, err
	}

	if opt.Token, err = security.NewJWT(xtime.RealClocker{}); err != nil {
		return nil, err
	}

	if opt.V, err = NewOIDCRequestParamValidator(); err != nil {
		return nil, err
	}

	return opt, nil
}

func ConfigureDatastore(ctx context.Context, cfg *config.Config, opt *option.Option) error {
	var err error

	if opt.DB, err = datastore.NewMySQL(ctx, cfg); err != nil {
		return err
	}

	if opt.Cache, err = datastore.NewRedis(ctx, cfg); err != nil {
		return err
	}

	return nil
}

func InitHandler(opt *option.Option) {
	InitAuthentication(opt)
	InitAuthorizationGet(opt)
	InitConfiguration()
	InitConsent(opt)
	InitHealthCheck(opt)
	InitRegistration(opt)
	InitToken(opt)
}
