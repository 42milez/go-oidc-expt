// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.14.0 DO NOT EDIT.
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

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// ErrorResponse represents error
type ErrorResponse struct {
	Detail string `json:"detail"`
	Status uint64 `json:"status"`
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

// Display defines model for Display.
type Display = string

// IdTokenHint defines model for IdTokenHint.
type IdTokenHint = string

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

// InternalServerError represents error
type InternalServerError = ErrorResponse

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
	// ClientId represents "client_id" parameter described in https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	ClientId ClientId `form:"client_id" json:"client_id" schema:"client_id" validate:"required,alphanum"`

	// Nonce represents "nonce" parameter described in https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	Nonce Nonce `form:"nonce" json:"nonce" schema:"nonce" validate:"required,ascii"`

	// RedirectUri represents "redirect_uri" parameter described in https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	RedirectUri RedirectUri `form:"redirect_uri" json:"redirect_uri" schema:"redirect_uri" validate:"required,printascii"`

	// ResponseType represents "response_type" parameter described in https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	ResponseType ResponseType `form:"response_type" json:"response_type" schema:"response_type" validate:"required,response-type-validator"`

	// Scope represents "scope" parameter described in https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	Scope Scope `form:"scope" json:"scope" schema:"scope" validate:"required,scope-validator"`

	// State represents "state" parameter described in https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	State State `form:"state" json:"state" schema:"state" validate:"required,alphanum"`

	// Display represents "display" parameter described in https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	Display Display `form:"display" json:"display" schema:"display" validate:"required"`

	// IdTokenHint represents "id_token_hint" parameter described in https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	IdTokenHint IdTokenHint `form:"id_token_hint" json:"id_token_hint" schema:"id_token_hint" validate:"required,alpha"`

	// MaxAge represents "max_age" parameter described in https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	MaxAge MaxAge `form:"max_age" json:"max_age" schema:"max_age" validate:"required"`

	// Prompt represents "prompt" parameter described in https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
	Prompt Prompt `form:"prompt" json:"prompt" schema:"prompt" validate:"required"`

	// Sid Session ID
	Sid *SessionId `form:"sid,omitempty" json:"-"`
}

// ConsentParams defines parameters for Consent.
type ConsentParams struct {
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

// RegisterJSONRequestBody defines body for Register for application/json ContentType.
type RegisterJSONRequestBody RegisterJSONBody

// --------------------------------------------------
//  Interface
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

	// POST: /register
	Register(w http.ResponseWriter, r *http.Request)
}

// --------------------------------------------------
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
		if mw := options.Middlewares.raw("Register"); mw != nil {
			r.Use(mw...)
		}
		r.Post("/register", si.Register)
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
