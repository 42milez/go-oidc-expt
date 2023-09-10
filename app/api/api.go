package api

import (
	_ "embed"
	"net/http"
)

//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen -config codegen/config.yml -o api.gen.go spec/spec.yml

//go:embed secret/key/block.key
var rawBlockKey []byte

//go:embed secret/key/hash.key
var rawHashKey []byte

var authenticateUserHdlr *AuthenticateHdlr
var checkHealthHdlr *CheckHealthHdlr
var registerUserHdlr *RegisterHdlr

type HandlerImpl struct{}

func (_ *HandlerImpl) Authenticate(w http.ResponseWriter, r *http.Request, params *AuthenticateParams) {
	authenticateUserHdlr.ServeHTTP(w, r, params)
}

func (_ *HandlerImpl) CheckHealth(w http.ResponseWriter, r *http.Request) {
	checkHealthHdlr.ServeHTTP(w, r)
}

func (_ *HandlerImpl) Register(w http.ResponseWriter, r *http.Request) {
	registerUserHdlr.ServeHTTP(w, r)
}
