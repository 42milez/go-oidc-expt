package api

import (
	_ "embed"
	"net/http"

	"github.com/42milez/go-oidc-server/app/httpstore"

	"github.com/42milez/go-oidc-server/app/pkg/xtime"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
	"github.com/go-playground/validator/v10"
)

//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen -config gen/config.yml -o api_gen.go spec/spec.yml

//go:embed secret/key/block.key
var rawBlockKey []byte

//go:embed secret/key/hash.key
var rawHashKey []byte

type HandlerImpl struct{}

func (_ *HandlerImpl) Authenticate(w http.ResponseWriter, r *http.Request) {
	authenticateUserHdlr.ServeHTTP(w, r)
}

func (_ *HandlerImpl) Authorize(w http.ResponseWriter, r *http.Request) {
	authorizeGetHdlr.ServeHTTP(w, r)
}

func (_ *HandlerImpl) CheckHealth(w http.ResponseWriter, r *http.Request) {
	checkHealthHdlr.ServeHTTP(w, r)
}

func (_ *HandlerImpl) Consent(w http.ResponseWriter, r *http.Request) {
	consentHdlr.ServeHTTP(w, r)
}

func (_ *HandlerImpl) Register(w http.ResponseWriter, r *http.Request) {
	registerUserHdlr.ServeHTTP(w, r)
}

func (_ *HandlerImpl) Token(w http.ResponseWriter, r *http.Request) {
	tokenHdlr.ServeHTTP(w, r)
}

type HandlerOption struct {
	cache          *datastore.Cache
	clock          xtime.Clocker
	cookie         *httpstore.Cookie
	db             *datastore.Database
	idGenerator    *xid.UniqueID
	tokenGenerator *JWT
	validator      *validator.Validate
}
