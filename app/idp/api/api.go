package api

import (
	_ "embed"
	"net/http"
)

//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen -config gen/config.yml -o api_gen.go spec/spec.yml

type HandlerImpl struct{}

func (_ *HandlerImpl) Authenticate(w http.ResponseWriter, r *http.Request) {
	authentication.ServeHTTP(w, r)
}

func (_ *HandlerImpl) Authorize(w http.ResponseWriter, r *http.Request) {
	authorizationGet.ServeHTTP(w, r)
}

func (_ *HandlerImpl) CheckHealth(w http.ResponseWriter, r *http.Request) {
	healthCheck.ServeHTTP(w, r)
}

func (_ *HandlerImpl) Consent(w http.ResponseWriter, r *http.Request) {
	consent.ServeHTTP(w, r)
}

func (_ *HandlerImpl) Register(w http.ResponseWriter, r *http.Request) {
	registration.ServeHTTP(w, r)
}

func (_ *HandlerImpl) Token(w http.ResponseWriter, r *http.Request) {
	token.ServeHTTP(w, r)
}
