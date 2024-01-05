// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.1-0.20231204155340-1f53862bcc64 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/pkg/typedef"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/oapi-codegen/runtime"
)

const (
	BasicAuthScopes = "basicAuth.Scopes"
)

// Defines values for TokenErrorResponseError.
const (
	TokenErrorResponseErrorInvalidClient        TokenErrorResponseError = "invalid_client"
	TokenErrorResponseErrorInvalidGrant         TokenErrorResponseError = "invalid_grant"
	TokenErrorResponseErrorInvalidRequest       TokenErrorResponseError = "invalid_request"
	TokenErrorResponseErrorInvalidScope         TokenErrorResponseError = "invalid_scope"
	TokenErrorResponseErrorUnauthorizedClient   TokenErrorResponseError = "unauthorized_client"
	TokenErrorResponseErrorUnsupportedGrantType TokenErrorResponseError = "unsupported_grant_type"
)

// ErrorResponse represents error response
type ErrorResponse struct {
	Details *[]string        `json:"details,omitempty"`
	Status  uint64           `json:"status"`
	Summary xerr.PublicError `json:"summary"`
}

// Health represents the status of service.
type Health struct {
	Status uint64 `json:"status"`
}

// TokenErrorResponse https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenErrorResponse
type TokenErrorResponse struct {
	Error TokenErrorResponseError `json:"error"`
}

// TokenErrorResponseError defines model for TokenErrorResponse.Error.
type TokenErrorResponseError string

// TokenResponse defines model for TokenResponse.
type TokenResponse struct {
	AccessToken  string            `json:"access_token"`
	ExpiresIn    uint64            `json:"expires_in"`
	IdToken      *string           `json:"id_token,omitempty"`
	RefreshToken string            `json:"refresh_token"`
	TokenType    typedef.TokenType `json:"token_type"`
}

// User defines model for User.
type User struct {
	ID   typedef.UserID `json:"id" validate:"required"`
	Name string         `json:"name" validate:"required"`
}

// UserName represents a part of user data.
type UserName struct {
	Name string `json:"name" validate:"required"`
}

// UserPassword represents the password of user
type UserPassword struct {
	Password string `json:"password" validate:"required"`
}

// ClientID defines model for ClientID.
type ClientID = string

// Code defines model for Code.
type Code = string

// Display defines model for Display.
type Display = string

// GrantType defines model for GrantType.
type GrantType = string

// MaxAge defines model for MaxAge.
type MaxAge = uint64

// Nonce defines model for Nonce.
type Nonce = string

// Prompt defines model for Prompt.
type Prompt = string

// RedirectUri defines model for RedirectUri.
type RedirectUri = string

// ResponseType defines model for ResponseType.
type ResponseType = string

// Scope defines model for Scope.
type Scope = string

// SessionId defines model for SessionId.
type SessionId = string

// State defines model for State.
type State = string

// InternalServerError represents error response
type InternalServerError = ErrorResponse

// InvalidRequest represents error response
type InvalidRequest = ErrorResponse

// UnauthorizedRequest represents error response
type UnauthorizedRequest = ErrorResponse

// AuthenticateJSONBody defines parameters for Authenticate.
type AuthenticateJSONBody struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// AuthenticateParams defines parameters for Authenticate.
type AuthenticateParams struct {
	// Sid Session ID
	Sid *SessionId `form:"sid,omitempty" json:"-" url:"sid"`
}

// AuthorizeParams defines parameters for Authorize.
type AuthorizeParams struct {
	// ClientId represents "client_id" parameter
	ClientID ClientID `form:"client_id" json:"client_id" schema:"client_id" url:"client_id" validate:"required,alphanum"`

	// Nonce represents "nonce" parameter
	Nonce Nonce `form:"nonce" json:"nonce" schema:"nonce" url:"nonce" validate:"required,alphanum"`

	// RedirectUri represents "redirect_uri" parameter
	RedirectUri RedirectUri `form:"redirect_uri" json:"redirect_uri" schema:"redirect_uri" url:"redirect_uri" validate:"required,url_encoded"`

	// ResponseType represents "response_type" parameter
	ResponseType ResponseType `form:"response_type" json:"response_type" schema:"response_type" url:"response_type" validate:"required,response-type-validator"`

	// Scope represents "scope" parameter
	Scope Scope `form:"scope" json:"scope" schema:"scope" url:"scope" validate:"required,scope-validator"`

	// State represents "state" parameter
	State State `form:"state" json:"state" schema:"state" url:"state" validate:"required,alphanum"`

	// Display represents "display" parameter
	Display Display `form:"display" json:"display" schema:"display" url:"display" validate:"required,display-validator"`

	// MaxAge represents "max_age" parameter
	MaxAge MaxAge `form:"max_age" json:"max_age" schema:"max_age" url:"max_age" validate:"required,numeric"`

	// Prompt represents "prompt" parameter
	Prompt Prompt `form:"prompt" json:"prompt" schema:"prompt" url:"prompt" validate:"required,prompt-validator"`

	// Sid Session ID
	Sid *SessionId `form:"sid,omitempty" json:"-" url:"sid"`
}

// ConsentParams defines parameters for Consent.
type ConsentParams struct {
	// Sid Session ID
	Sid *SessionId `form:"sid,omitempty" json:"-" url:"sid"`
}

// TokenFormdataBody defines parameters for Token.
type TokenFormdataBody struct {
	// Code represents "code" parameter
	Code *string `form:"code" json:"code" schema:"code" validate:"omitempty,alphanum"`

	// GrantType represents "grant_type" parameter
	GrantType string `form:"grant_type" json:"grant_type" schema:"grant_type" validate:"required,grant-type-validator"`

	// RedirectUri represents "redirect_uri" parameter
	RedirectUri *string `form:"redirect_uri" json:"redirect_uri" schema:"redirect_uri" validate:"omitempty,url_encoded"`

	// RefreshToken represents "refresh_token" parameter
	RefreshToken *string `form:"refresh_token" json:"refresh_token" schema:"refresh_token" validate:"omitempty,jwt"`
}

// TokenParams defines parameters for Token.
type TokenParams struct {
	// Sid Session ID
	Sid *SessionId `form:"sid,omitempty" json:"-" url:"sid"`
}

// RegisterJSONBody defines parameters for Register.
type RegisterJSONBody struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// AuthenticateJSONRequestBody defines body for Authenticate for application/json ContentType.
type AuthenticateJSONRequestBody AuthenticateJSONBody

// TokenFormdataRequestBody defines body for Token for application/x-www-form-urlencoded ContentType.
type TokenFormdataRequestBody TokenFormdataBody

// RegisterJSONRequestBody defines body for Register for application/json ContentType.
type RegisterJSONRequestBody RegisterJSONBody

//  Interface
// --------------------------------------------------

// HandlerInterface represents all server handlers.
type HandlerInterface interface {

	// POST: /authentication
	Authenticate(w http.ResponseWriter, r *http.Request)

	// GET: /authorization
	Authorize(w http.ResponseWriter, r *http.Request)

	// POST: /consent
	Consent(w http.ResponseWriter, r *http.Request)

	// GET: /health
	CheckHealth(w http.ResponseWriter, r *http.Request)

	// POST: /token
	Token(w http.ResponseWriter, r *http.Request)

	// POST: /user/registration
	Register(w http.ResponseWriter, r *http.Request)
}

//  Middleware
// --------------------------------------------------

// HandlerInterfaceWrapper converts contexts to parameters.
type HandlerInterfaceWrapper struct {
	Handler          HandlerInterface
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

func NewMiddlewareFuncMap() *MiddlewareFuncMap {
	return &MiddlewareFuncMap{
		m: make(map[string][]MiddlewareFunc),
	}
}

type MiddlewareFuncMap struct {
	m map[string][]MiddlewareFunc
}

func (mfm *MiddlewareFuncMap) raw(key string) []func(http.Handler) http.Handler {
	ret := make([]func(http.Handler) http.Handler, len(mfm.m[key]), len(mfm.m[key]))
	v, ok := mfm.m[key]
	if !ok {
		return nil
	}
	for i, f := range v {
		ret[i] = f
	}
	return ret
}

func (mfm *MiddlewareFuncMap) SetAuthenticateMW(mf ...MiddlewareFunc) *MiddlewareFuncMap {
	mfm.append("Authenticate", mf...)
	return mfm
}

func (mfm *MiddlewareFuncMap) SetAuthorizeMW(mf ...MiddlewareFunc) *MiddlewareFuncMap {
	mfm.append("Authorize", mf...)
	return mfm
}

func (mfm *MiddlewareFuncMap) SetConsentMW(mf ...MiddlewareFunc) *MiddlewareFuncMap {
	mfm.append("Consent", mf...)
	return mfm
}

func (mfm *MiddlewareFuncMap) SetCheckHealthMW(mf ...MiddlewareFunc) *MiddlewareFuncMap {
	mfm.append("CheckHealth", mf...)
	return mfm
}

func (mfm *MiddlewareFuncMap) SetTokenMW(mf ...MiddlewareFunc) *MiddlewareFuncMap {
	mfm.append("Token", mf...)
	return mfm
}

func (mfm *MiddlewareFuncMap) SetRegisterMW(mf ...MiddlewareFunc) *MiddlewareFuncMap {
	mfm.append("Register", mf...)
	return mfm
}

func (mfm *MiddlewareFuncMap) append(key string, mf ...MiddlewareFunc) {
	for _, v := range mf {
		mfm.m[key] = append(mfm.m[key], v)
	}
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

func injectRequestParameter(r *http.Request) (*http.Request, error) {
	switch r.URL.Path {
	case "/authentication":
		return unmarshalAuthenticateParameter(r)
	case "/authorization":
		return unmarshalAuthorizeParameter(r)
	case "/consent":
		return unmarshalConsentParameter(r)
	case "/health":
		return unmarshalCheckHealthParameter(r)
	case "/token":
		return unmarshalTokenParameter(r)
	case "/user/registration":
		return unmarshalRegisterParameter(r)
	default:
		return nil, xerr.InvalidPath
	}
}

func unmarshalAuthenticateParameter(r *http.Request) (*http.Request, error) {
	var ctx context.Context
	var err error

	// ==================================================
	// Unmarshal Parameter: BEGIN

	if ctx == nil {
		ctx = r.Context()
	}

	params := &AuthenticateParams{}

	// --------------------------------------------------
	// Cookie Parameter: BEGIN

	var cookie *http.Cookie

	if cookie, err = r.Cookie("sid"); err == nil {
		var value SessionId
		err = runtime.BindStyledParameter("simple", true, "sid", cookie.Value, &value)
		if err != nil {
			return nil, &InvalidParamFormatError{
				ParamName: "sid",
				Err:       err,
			}
		}
		params.Sid = &value
	}

	// Cookie Parameter: END
	// --------------------------------------------------

	ctx = context.WithValue(ctx, typedef.RequestParamKey{}, params)

	// Unmarshal Parameter: END
	// ==================================================

	return r.Clone(ctx), nil
}

func unmarshalAuthorizeParameter(r *http.Request) (*http.Request, error) {
	var ctx context.Context
	var err error

	// ==================================================
	// Unmarshal Parameter: BEGIN

	if ctx == nil {
		ctx = r.Context()
	}

	params := &AuthorizeParams{}

	// --------------------------------------------------
	//  Query Parameter: BEGIN

	if ctx == nil {
		ctx = r.Context()
	}

	// Required query parameter "client_id"
	err = runtime.BindQueryParameter("form", true, true, "client_id", r.URL.Query(), &params.ClientID)
	if err != nil {
		return nil, &InvalidParamFormatError{
			ParamName: "client_id",
			Err:       err,
		}
	}

	// Required query parameter "nonce"
	err = runtime.BindQueryParameter("form", true, true, "nonce", r.URL.Query(), &params.Nonce)
	if err != nil {
		return nil, &InvalidParamFormatError{
			ParamName: "nonce",
			Err:       err,
		}
	}

	// Required query parameter "redirect_uri"
	err = runtime.BindQueryParameter("form", true, true, "redirect_uri", r.URL.Query(), &params.RedirectUri)
	if err != nil {
		return nil, &InvalidParamFormatError{
			ParamName: "redirect_uri",
			Err:       err,
		}
	}

	// Required query parameter "response_type"
	err = runtime.BindQueryParameter("form", true, true, "response_type", r.URL.Query(), &params.ResponseType)
	if err != nil {
		return nil, &InvalidParamFormatError{
			ParamName: "response_type",
			Err:       err,
		}
	}

	// Required query parameter "scope"
	err = runtime.BindQueryParameter("form", true, true, "scope", r.URL.Query(), &params.Scope)
	if err != nil {
		return nil, &InvalidParamFormatError{
			ParamName: "scope",
			Err:       err,
		}
	}

	// Required query parameter "state"
	err = runtime.BindQueryParameter("form", true, true, "state", r.URL.Query(), &params.State)
	if err != nil {
		return nil, &InvalidParamFormatError{
			ParamName: "state",
			Err:       err,
		}
	}

	// Required query parameter "display"
	err = runtime.BindQueryParameter("form", true, true, "display", r.URL.Query(), &params.Display)
	if err != nil {
		return nil, &InvalidParamFormatError{
			ParamName: "display",
			Err:       err,
		}
	}

	// Required query parameter "max_age"
	err = runtime.BindQueryParameter("form", true, true, "max_age", r.URL.Query(), &params.MaxAge)
	if err != nil {
		return nil, &InvalidParamFormatError{
			ParamName: "max_age",
			Err:       err,
		}
	}

	// Required query parameter "prompt"
	err = runtime.BindQueryParameter("form", true, true, "prompt", r.URL.Query(), &params.Prompt)
	if err != nil {
		return nil, &InvalidParamFormatError{
			ParamName: "prompt",
			Err:       err,
		}
	}

	//  Query Parameter: END
	// --------------------------------------------------

	// --------------------------------------------------
	// Cookie Parameter: BEGIN

	var cookie *http.Cookie

	if cookie, err = r.Cookie("sid"); err == nil {
		var value SessionId
		err = runtime.BindStyledParameter("simple", true, "sid", cookie.Value, &value)
		if err != nil {
			return nil, &InvalidParamFormatError{
				ParamName: "sid",
				Err:       err,
			}
		}
		params.Sid = &value
	}

	// Cookie Parameter: END
	// --------------------------------------------------

	ctx = context.WithValue(ctx, typedef.RequestParamKey{}, params)

	// Unmarshal Parameter: END
	// ==================================================

	return r.Clone(ctx), nil
}

func unmarshalConsentParameter(r *http.Request) (*http.Request, error) {
	var ctx context.Context
	var err error

	// ==================================================
	// Unmarshal Parameter: BEGIN

	if ctx == nil {
		ctx = r.Context()
	}

	params := &ConsentParams{}

	// --------------------------------------------------
	// Cookie Parameter: BEGIN

	var cookie *http.Cookie

	if cookie, err = r.Cookie("sid"); err == nil {
		var value SessionId
		err = runtime.BindStyledParameter("simple", true, "sid", cookie.Value, &value)
		if err != nil {
			return nil, &InvalidParamFormatError{
				ParamName: "sid",
				Err:       err,
			}
		}
		params.Sid = &value
	}

	// Cookie Parameter: END
	// --------------------------------------------------

	ctx = context.WithValue(ctx, typedef.RequestParamKey{}, params)

	// Unmarshal Parameter: END
	// ==================================================

	return r.Clone(ctx), nil
}

func unmarshalCheckHealthParameter(r *http.Request) (*http.Request, error) {
	return r, nil
}

func unmarshalTokenParameter(r *http.Request) (*http.Request, error) {
	var ctx context.Context
	var err error

	// ==================================================
	//  Security Definition: BEGIN

	if ctx == nil {
		ctx = r.Context()
	}

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	//  Security Definition: END
	// ==================================================

	// ==================================================
	// Unmarshal Parameter: BEGIN

	if ctx == nil {
		ctx = r.Context()
	}

	params := &TokenParams{}

	// --------------------------------------------------
	// Cookie Parameter: BEGIN

	var cookie *http.Cookie

	if cookie, err = r.Cookie("sid"); err == nil {
		var value SessionId
		err = runtime.BindStyledParameter("simple", true, "sid", cookie.Value, &value)
		if err != nil {
			return nil, &InvalidParamFormatError{
				ParamName: "sid",
				Err:       err,
			}
		}
		params.Sid = &value
	}

	// Cookie Parameter: END
	// --------------------------------------------------

	ctx = context.WithValue(ctx, typedef.RequestParamKey{}, params)

	// Unmarshal Parameter: END
	// ==================================================

	return r.Clone(ctx), nil
}

func unmarshalRegisterParameter(r *http.Request) (*http.Request, error) {
	return r, nil
}

//  Handler and others
// --------------------------------------------------

type ChiServerOptions struct {
	BaseRouter       *chi.Mux
	Middlewares      *MiddlewareFuncMap
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// MuxWithOptions creates http.Handler with additional options
func MuxWithOptions(hi HandlerInterface, option *ChiServerOptions, appOption *option.Option) (*chi.Mux, error) {
	r := option.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}

	if option.ErrorHandlerFunc == nil {
		option.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	swag, err := GetSwagger()
	if err != nil {
		return nil, err
	}
	swag.Servers = nil

	oapiValidator := nethttpmiddleware.OapiRequestValidatorWithOptions(swag, &nethttpmiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: NewOapiAuthentication(appOption),
		},
		ErrorHandler: NewOapiErrorHandler(),
	})

	r.Group(func(r chi.Router) {
		r.Use(oapiValidator)
		r.Use(InjectRequestParameter())
		if mw := option.Middlewares.raw("Authenticate"); mw != nil {
			r.Use(mw...)
		}
		r.Post("/authentication", hi.Authenticate)
	})

	r.Group(func(r chi.Router) {
		r.Use(oapiValidator)
		r.Use(InjectRequestParameter())
		if mw := option.Middlewares.raw("Authorize"); mw != nil {
			r.Use(mw...)
		}
		r.Get("/authorization", hi.Authorize)
	})

	r.Group(func(r chi.Router) {
		r.Use(oapiValidator)
		r.Use(InjectRequestParameter())
		if mw := option.Middlewares.raw("Consent"); mw != nil {
			r.Use(mw...)
		}
		r.Post("/consent", hi.Consent)
	})

	r.Group(func(r chi.Router) {
		r.Use(oapiValidator)
		r.Use(InjectRequestParameter())
		if mw := option.Middlewares.raw("CheckHealth"); mw != nil {
			r.Use(mw...)
		}
		r.Get("/health", hi.CheckHealth)
	})

	r.Group(func(r chi.Router) {
		r.Use(oapiValidator)
		r.Use(InjectRequestParameter())
		if mw := option.Middlewares.raw("Token"); mw != nil {
			r.Use(mw...)
		}
		r.Post("/token", hi.Token)
	})

	r.Group(func(r chi.Router) {
		r.Use(oapiValidator)
		r.Use(InjectRequestParameter())
		if mw := option.Middlewares.raw("Register"); mw != nil {
			r.Use(mw...)
		}
		r.Post("/user/registration", hi.Register)
	})

	return r, nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
