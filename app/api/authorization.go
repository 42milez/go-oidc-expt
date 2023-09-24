package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/api/oapigen"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

type AuthorizeGetHdlr struct {
	service   Authorizer
	validator *validator.Validate
}

func (ag *AuthorizeGetHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := schema.NewDecoder()
	q := &oapigen.AuthorizeParams{}

	if err := decoder.Decode(q, r.URL.Query()); err != nil {
		RespondJSON500(w, err)
		return
	}

	if err := ag.validator.Struct(q); err != nil {
		RespondJSON400(w, xerr.InvalidRequest, nil, err)
		return
	}

	// TODO: Redirect unauthenticated user to the authentication endpoint with the posted parameters
	// ...

	// TODO: Redirect authenticated user to the consent endpoint with the posted parameters
	// ...

	location, err := ag.service.Authorize(r.Context(), q.ClientId, q.RedirectUri, q.State)

	if err != nil {
		RespondJSON400(w, xerr.InvalidRequest, nil, err)
		return
	}

	http.Redirect(w, r, location, http.StatusFound)
}

func NewAuthorizePost() *AuthorizePost {
	return &AuthorizePost{}
}

type AuthorizePost struct {
	Service   Authorizer
	Validator *validator.Validate
}

func (p *AuthorizePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NOT IMPLEMENTED YET
}
