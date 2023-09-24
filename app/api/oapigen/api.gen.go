// Package oapigen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.14.0 DO NOT EDIT.
package oapigen

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

// User defines model for User.
type User struct {
	Id   uint64 `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

// UserName represents a part of user data.
type UserName struct {
	Name string `json:"name" validate:"required"`
}

// UserPassword represents the password of user
type UserPassword struct {
	Password string `json:"password" validate:"required"`
}

// ClientId defines model for ClientId.
type ClientId = string

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
	Sid *SessionId `form:"sid,omitempty" json:"-"`
}

// AuthorizeParams defines parameters for Authorize.
type AuthorizeParams struct {
	// ClientId represents "client_id" parameter
	ClientId ClientId `form:"client_id" json:"client_id" schema:"client_id" validate:"required,alphanum"`

	// Nonce represents "nonce" parameter
	Nonce Nonce `form:"nonce" json:"nonce" schema:"nonce" validate:"required,ascii"`

	// RedirectUri represents "redirect_uri" parameter
	RedirectUri RedirectUri `form:"redirect_uri" json:"redirect_uri" schema:"redirect_uri" validate:"required,printascii"`

	// ResponseType represents "response_type" parameter
	ResponseType ResponseType `form:"response_type" json:"response_type" schema:"response_type" validate:"required,response-type-validator"`

	// Scope represents "scope" parameter
	Scope Scope `form:"scope" json:"scope" schema:"scope" validate:"required,scope-validator"`

	// State represents "state" parameter
	State State `form:"state" json:"state" schema:"state" validate:"required,alphanum"`

	// Display represents "display" parameter
	Display Display `form:"display" json:"display" schema:"display" validate:"required"`

	// MaxAge represents "max_age" parameter
	MaxAge MaxAge `form:"max_age" json:"max_age" schema:"max_age" validate:"required"`

	// Prompt represents "prompt" parameter
	Prompt Prompt `form:"prompt" json:"prompt" schema:"prompt" validate:"required"`

	// Sid Session ID
	Sid *SessionId `form:"sid,omitempty" json:"-"`
}

// ConsentParams defines parameters for Consent.
type ConsentParams struct {
	// Sid Session ID
	Sid *SessionId `form:"sid,omitempty" json:"-"`
}

// TokenFormdataBody defines parameters for Token.
type TokenFormdataBody = struct {
}

// TokenParams defines parameters for Token.
type TokenParams struct {
	// Sid Session ID
	Sid *SessionId `form:"sid,omitempty" json:"-"`
}

// RegisterJSONBody defines parameters for Register.
type RegisterJSONBody struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// AuthenticateJSONRequestBody defines body for Authenticate for application/json ContentType.
type AuthenticateJSONRequestBody AuthenticateJSONBody

// TokenFormdataRequestBody defines body for Token for application/x-www-form-urlencoded ContentType.
type TokenFormdataRequestBody = TokenFormdataBody

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
