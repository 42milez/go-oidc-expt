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
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

// UserName represents a part of user data.
type UserName struct {
	Name string `json:"name"`
}

// UserPassword represents the password of user
type UserPassword struct {
	Password string `json:"password"`
}

// AuthenticateUserJSONBody defines parameters for AuthenticateUser.
type AuthenticateUserJSONBody struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// RegisterUserJSONBody defines parameters for RegisterUser.
type RegisterUserJSONBody struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// AuthenticateUserJSONRequestBody defines body for AuthenticateUser for application/json ContentType.
type AuthenticateUserJSONRequestBody AuthenticateUserJSONBody

// RegisterUserJSONRequestBody defines body for RegisterUser for application/json ContentType.
type RegisterUserJSONRequestBody RegisterUserJSONBody

// --------------------------------------------------
//  Interface
// --------------------------------------------------

// HandlerInterface represents all server handlers.
type HandlerInterface interface {

	// POST: /authenticate
	AuthenticateUser(w http.ResponseWriter, r *http.Request)

	// GET: /health
	CheckHealth(w http.ResponseWriter, r *http.Request)

	// POST: /register
	RegisterUser(w http.ResponseWriter, r *http.Request)
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

func (mfm *MiddlewareFuncMap) SetAuthenticateUserMW(mf []MiddlewareFunc) {
	mfm.m["AuthenticateUser"] = mf
}

func (mfm *MiddlewareFuncMap) SetCheckHealthMW(mf []MiddlewareFunc) {
	mfm.m["CheckHealth"] = mf
}

func (mfm *MiddlewareFuncMap) SetRegisterUserMW(mf []MiddlewareFunc) {
	mfm.m["RegisterUser"] = mf
}

// AuthenticateUser operation middleware
func (siw *HandlerInterfaceWrapper) AuthenticateUser() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		siw.Handler.AuthenticateUser(w, r.WithContext(ctx))
	})
}

// CheckHealth operation middleware
func (siw *HandlerInterfaceWrapper) CheckHealth() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		siw.Handler.CheckHealth(w, r.WithContext(ctx))
	})
}

// RegisterUser operation middleware
func (siw *HandlerInterfaceWrapper) RegisterUser() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		siw.Handler.RegisterUser(w, r.WithContext(ctx))
	})
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
	BaseURL          string
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

	wrapper := HandlerInterfaceWrapper{
		Handler:          si,
		ErrorHandlerFunc: options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		if mw := options.Middlewares.raw("AuthenticateUser"); mw != nil {
			r.Use(mw...)
		}
		r.Post(options.BaseURL+"/authenticate", wrapper.AuthenticateUser())
	})

	r.Group(func(r chi.Router) {
		if mw := options.Middlewares.raw("CheckHealth"); mw != nil {
			r.Use(mw...)
		}
		r.Get(options.BaseURL+"/health", wrapper.CheckHealth())
	})

	r.Group(func(r chi.Router) {
		if mw := options.Middlewares.raw("RegisterUser"); mw != nil {
			r.Use(mw...)
		}
		r.Post(options.BaseURL+"/register", wrapper.RegisterUser())
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RXTXPbNhD9Kxi0R4aU0qST4amOnWl9aO2Jk1NGBxhciUhBAAWWcjwa/ffOAqBCSvJH",
	"2njqXjwcaD8e3r5drDdc2s5ZAwYDrzc8yBY6ET/feW/9ewjOmgB00ECQXjlU1vCae3AeAvkxIEtecOet",
	"A48KQjJHoTR9wRfROQ285r2BLw4kQpOcmJWy9x4aXnC8dWQS0Cuz4tuCBxTYh0mA17NZwZfWdwIpmjL4",
	"86uvrsogrMDz7bbgHv7qFQWuPw2BigHSYudhrz+DREr2GwiN7b3XxBZYCsXskgXwayWhPLj3Edgv/w3s",
	"Y2g/BvCUQWh9seT1p80eBtXQ32/OqJpj2Tb8Rw9LXvMfqq9iqbJSKoLyh+iAbxfF3dz1ATxrBIpygB99",
	"7qNbMCc8EtUj532yTY4yklgAT6ezQ0nt3Tb63sXupQjhxvrmQUW4bDgAPYDoRpH+Ccyd/yFUMlVmaSm0",
	"tAaFxJili43HQ++c9fhLzlpK2/EiU8ZPLs/ZVTLgBe89ObSILtRVNXKocpCKiJkScWIYdbNXHRgUminy",
	"id9kQHxcODDnZ+zUGgMS2Y3Clv1qecG1kpCnSobz+/mHHYwPb88oG4LvwsXyKjVahreHLtpUxKHCyOvK",
	"vrCqkS+oPWMx1uBDwjsvZ+WMAlsHRjjFa/5TOS+pAk5gG4tViR5bMKikwAjP2YCHGjgZWZFSqZokTqp7",
	"vP15s2f1MUmDCgsB39rmdigamJhAOKfJUllTfQ6UZZjG00Z/XC8+oml3Ct8uopCmNyQLhpZN+BgLE30P",
	"UanpgYjsvZzNvulWD0HkR4Bd9VJCCMte69sJuiaOllffEcL0BTyCZVRgUvxSKA0Na3og5pRZC60alkue",
	"0c2fGbphDjHrd6MsQn39bIjcXxrKOCJRrAKNx6krX9BvVbt7zVdwpH1PW5B/xumdDPfe82kXR+O8Hjyh",
	"3nOGI+RkZGwHS+jvXaCHkxuLUwCjEiTvyFPm38NKBUw7yvEJeuohz04DN+mFVyaWhB76axGOVOJ9jvp/",
	"naUDK8watJO7Pr/JOmD9L8ZqqrN/7FB9/Uyw3T+nxo48qSTtKCFKcZrlDNagraNlir0za+Wtoe/JolZX",
	"lbZS6NYGrN/M3syqDUnpUmC7pb1HeCWudZLP8EPqw6XoNf1bsJ7zgoPpO8K3nidU+1CuUKyUWd0Jg/bF",
	"kGxK1bhyvJ09CaBLb5teRv7vw/QUWBa7cu721vHsI6z5eFLt0fnea7VdbP8OAAD//7J29YiCDwAA",
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
