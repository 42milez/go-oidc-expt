package api

import (
	_ "embed"
	"net/http"
)

//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen -config codegen/config.yml -o api.gen.go spec/spec.yml

//go:embed secret/keypair/private.pem
var rawPrivateKey []byte

//go:embed secret/keypair/public.pem
var rawPublicKey []byte

//go:embed secret/key/block.key
var rawBlockKey []byte

//go:embed secret/key/hash.key
var rawHashKey []byte

var checkHealthHdlr *CheckHealthHdlr
var authenticateUserHdlr *AuthenticateUserHdlr
var registerUserHdlr *RegisterUserHdlr

type HandlerImpl struct{}

func (p *HandlerImpl) CheckHealth(w http.ResponseWriter, r *http.Request) {
	checkHealthHdlr.ServeHTTP(w, r)
}

func (p *HandlerImpl) AuthenticateUser(w http.ResponseWriter, r *http.Request, params *AuthenticateUserParams) {
	authenticateUserHdlr.ServeHTTP(w, r, params)
}

func (p *HandlerImpl) RegisterUser(w http.ResponseWriter, r *http.Request) {
	registerUserHdlr.ServeHTTP(w, r)
}
