package api

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/api/validation"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/api/session"

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

// ServeHTTP authorizes a request that asking to access the resources belonging to a user
//
//	@summary		TBD
//	@description	TBD
//	@id				AuthorizeGet.ServeHTTP
//	@tags			OIDC
//	@accept			json
//	@produce		json
//	@param			name		query		string	true	"TBD"
//	@param			password	query		string	true	"TBD"
//	@success		200			{string}	string
//	@failure		500			{object}	model.ErrorResponse
//	@failure		500			{object}	model.ErrorResponse
//	@router			/v1/authorize [get]
func (p *AuthorizeGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := schema.NewDecoder()
	q := &model.AuthorizeRequest{}

	if err := decoder.Decode(q, r.URL.Query()); err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	if err := p.validator.Struct(q); err != nil {
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.InvalidRequest,
		})
		return
	}

	// TODO: Redirect unauthenticated user to the authentication endpoint with the posted parameters
	// ...

	// TODO: Redirect authenticated user to the consent endpoint with the posted parameters
	// ...

	userID, ok := session.GetUserID(r.Context())

	if !ok {
		RespondJSON(w, http.StatusUnauthorized, &ErrResponse{
			Error: xerr.UnauthorizedRequest,
		})
		return
	}

	location, err := p.Service.Authorize(r.Context(), userID, q)

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

// ServeHTTP authorizes a request that asking to access the resources belonging to a user
//
//	@summary		TBD
//	@description	TBD
//	@id				AuthorizePost.ServeHTTP
//	@tags			OIDC
//	@accept			json
//	@produce		json
//	@param			name		formData	string	true	"TBD"
//	@param			password	formData	string	true	"TBD"
//	@success		200			{string}	string
//	@failure		500			{object}	model.ErrorResponse
//	@failure		500			{object}	model.ErrorResponse
//	@router			/v1/authorize [post]
func (p *AuthorizePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NOT IMPLEMENTED YET
}
