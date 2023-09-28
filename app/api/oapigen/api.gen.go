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
type TokenFormdataBody struct {
	Code         *string `form:"code" json:"code"`
	GrantType    string  `form:"grant_type" json:"grant_type"`
	RedirectUri  *string `form:"redirect_uri" json:"redirect_uri"`
	RefreshToken *string `form:"refresh_token" json:"refresh_token"`
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
