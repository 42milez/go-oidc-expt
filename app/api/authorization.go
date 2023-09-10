package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/api/validation"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

func NewAuthorizeGet() (*AuthorizeGet, error) {
	v, err := validation.NewAuthorizeValidator()

	if err != nil {
		return nil, err
	}

	ret := &AuthorizeGet{
		validator: v,
	}

	return ret, nil
}

type AuthorizeGet struct {
	Service   Authorizer
	validator *validator.Validate
}

func (ag *AuthorizeGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := schema.NewDecoder()
	q := &model.AuthorizeRequest{}

	if err := decoder.Decode(q, r.URL.Query()); err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	if err := ag.validator.Struct(q); err != nil {
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.InvalidRequest,
		})
		return
	}

	// TODO: Redirect unauthenticated user to the authentication endpoint with the posted parameters
	// ...

	// TODO: Redirect authenticated user to the consent endpoint with the posted parameters
	// ...

	userID, ok := service.GetUserID(r.Context())

	if !ok {
		RespondJSON(w, http.StatusUnauthorized, &ErrResponse{
			Error: xerr.UnauthorizedRequest,
		})
		return
	}

	location, err := ag.Service.Authorize(r.Context(), userID, q)

	if err != nil {
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.InvalidParameter,
		})
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
