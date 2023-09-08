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
	"github.com/oapi-codegen/runtime"
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

// AuthenticateUserParams defines parameters for AuthenticateUser.
type AuthenticateUserParams struct {
	// Sid Session ID
	Sid *string `form:"sid,omitempty" json:"sid,omitempty"`
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
	AuthenticateUser(w http.ResponseWriter, r *http.Request, params *AuthenticateUserParams)

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

		var err error

		// Parameter object where we will unmarshal all parameters from the context
		params := &AuthenticateUserParams{}

		var cookie *http.Cookie
		if cookie, err = r.Cookie("sid"); err == nil {

			var value string
			err = runtime.BindStyledParameter("simple", true, "sid", cookie.Value, &value)
			if err != nil {
				siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sid", Err: err})
				return
			}
			params.Sid = &value

		}

		siw.Handler.AuthenticateUser(w, r.WithContext(ctx), params)
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
		r.Post("/authenticate", wrapper.AuthenticateUser())
	})

	r.Group(func(r chi.Router) {
		if mw := options.Middlewares.raw("CheckHealth"); mw != nil {
			r.Use(mw...)
		}
		r.Get("/health", wrapper.CheckHealth())
	})

	r.Group(func(r chi.Router) {
		if mw := options.Middlewares.raw("RegisterUser"); mw != nil {
			r.Use(mw...)
		}
		r.Post("/register", wrapper.RegisterUser())
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RWX2/bNhD/KsRtj6qldOlQ6GlpUmx52BI07VPhB5Y622xlkiNPSYPA3304ipYlWXHd",
	"diuWF4OQ78/v7n735wGUXTtr0FCA8gGCWuFaxudr761/g8FZE5A/VBiU1460NVCCR+cxsJ5AloQMnLcO",
	"PWkMrThJXfMLP8u1qxFKgAzo3vErkNdmCZsMAklqwkCuyGBh/VoSlNBoQ7+e7hS1IVyih80mA49/N9pj",
	"BeX7rZls63feadgPH1ERu/oDZU2rg7HQCkVrStiFCOhvtcLZXnAToJ8X3wF7Cu27gJ49yLq+WkD5/mGE",
	"QVf8+9UedTXl7QF+9riAEn7Kd4zIEx1yhvKXXCNs5tnjuWsCelFJkrMt/KhzKN1SOOmJU91THifbJCs7",
	"HrEwfy32CTWKNuo+lt1rGcKd9dUXGeGS4BboHkTXs/QtMDv9fagsqs3CsmllDUlF0cs6dheExjnr6bfk",
	"dabsGrKUMji7vhQ3rQBk0HhWWBG5UOZ5TyFPRnJOzDARZ0bgZ4der9GQrIVmnfhmAc7HlUNzeSHOrTGo",
	"SNxpWonfLWRQa4VpdCQ4f16+7WC8fXXB3gj9OlwtbtpGS/BG6KJMzjnUFPO6tM+srtQzbs9YjFv0ocV7",
	"MitmBRu2Do10Gkr4ZXYy4wo4SatYrFw2tEJDWkmK8JwNtM+Bs54UM5WryeTkusfoL6uR1LtEDenlGgl9",
	"iF07tHqDgaGKywvgukIJytpPGndFC7qCLE1iRjVmzrylDgZ6Zav7LS3QxBCkczVj0dbkHwN7fOiZ2o2S",
	"47r9iLHQ9dBmHqk6jJYlBFkxyHif+uQbjL3Q7plYn+dF8VVRfQkiTAC7aZTCEBZNXd8P0FVxeJ0eBaHr",
	"9N22A21uZa0rkUoEuwV3WhSb7EjYw+U7gb9HO2bTQuoaK1E1yNkeYUgRnXxnRNtpJqzvBuIgvJMfHN4U",
	"oBjri2+uXmN42inCqr1qhFWq8czUXaAvflgdx2hmcW+QXPJgGanCnP/LV92Js8SJmXa+QvUprrRWcHTk",
	"DEdbFE4303/YosnDRHISMtHBkvXx5f23nBtLQwC9ErTaMU8p/x6XOlB7uE2vlXOPaaEYvGvPHm1iSfj6",
	"+SDDRCXeJKtpwTy58b/NirCG7CDW/98y2GJ9Kpug5YY/dg88odl4ILLDk7GvCC0v21Nx6iC7wFusreOb",
	"Vrw2t9pbw+/BvVzmeW2VrFc2UPmyeFnksTlGLCK51Gb5qBE+ukMrM9OVm/VP3Alz195WjYqxH7K4b2ne",
	"JaI7vPtzij2lz4M89b6PNstmvvknAAD//z8+0lAoEAAA",
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
