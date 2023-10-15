// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
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
	IdToken      string            `json:"id_token"`
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
	Code         *string `form:"code" json:"code"`
	GrantType    string  `form:"grant_type" json:"grant_type"`
	RedirectUri  *string `form:"redirect_uri" json:"redirect_uri"`
	RefreshToken *string `form:"refresh_token" json:"refresh_token"`
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
