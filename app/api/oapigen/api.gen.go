// Package oapigen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
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

// User defines model for User.
type User struct {
	Id   typedef.UserID `json:"id" validate:"required"`
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
