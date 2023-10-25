// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/typedef"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

const (
	BasicAuthScopes = "basicAuth.Scopes"
)

// ErrorResponse represents error response
type ErrorResponse struct {
	Details *[]string      `json:"details,omitempty"`
	Status  uint64         `json:"status"`
	Summary xerr.PublicErr `json:"summary"`
}

// Health represents the status of service.
type Health struct {
	Status uint64 `json:"status"`
}

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

//  INTERFACE
// --------------------------------------------------

// HandlerInterface represents all server handlers.
type HandlerInterface interface {

	// POST: /authenticate
	Authenticate(w http.ResponseWriter, r *http.Request)

	// GET: /authorize
	Authorize(w http.ResponseWriter, r *http.Request)

	// POST: /consent
	Consent(w http.ResponseWriter, r *http.Request)

	// GET: /health
	CheckHealth(w http.ResponseWriter, r *http.Request)

	// POST: /token
	Token(w http.ResponseWriter, r *http.Request)

	// POST: /user/register
	Register(w http.ResponseWriter, r *http.Request)
}

//  MIDDLEWARE
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

//  HANDLER AND OTHERS
// --------------------------------------------------

type ChiServerOptions struct {
	BaseRouter       *chi.Mux
	Middlewares      *MiddlewareFuncMap
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// MuxWithOptions creates http.Handler with additional options
func MuxWithOptions(si HandlerInterface, options *ChiServerOptions) *chi.Mux {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}

	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	r.Group(func(r chi.Router) {
		if mw := options.Middlewares.raw("Authenticate"); mw != nil {
			r.Use(mw...)
		}
		r.Post("/authenticate", si.Authenticate)
	})

	r.Group(func(r chi.Router) {
		if mw := options.Middlewares.raw("Authorize"); mw != nil {
			r.Use(mw...)
		}
		r.Get("/authorize", si.Authorize)
	})

	r.Group(func(r chi.Router) {
		if mw := options.Middlewares.raw("Consent"); mw != nil {
			r.Use(mw...)
		}
		r.Post("/consent", si.Consent)
	})

	r.Group(func(r chi.Router) {
		if mw := options.Middlewares.raw("CheckHealth"); mw != nil {
			r.Use(mw...)
		}
		r.Get("/health", si.CheckHealth)
	})

	r.Group(func(r chi.Router) {
		if mw := options.Middlewares.raw("Token"); mw != nil {
			r.Use(mw...)
		}
		r.Post("/token", si.Token)
	})

	r.Group(func(r chi.Router) {
		if mw := options.Middlewares.raw("Register"); mw != nil {
			r.Use(mw...)
		}
		r.Post("/user/register", si.Register)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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
