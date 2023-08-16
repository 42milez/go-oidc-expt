package handler

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/validation"

	"github.com/42milez/go-oidc-server/app/idp/model"
	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

func NewAuthorizeGet() (*AuthorizeGet, error) {
	v, err := validation.NewAuthorizeValidator()

	if err != nil {
		return nil, err
	}

	ret := &AuthorizeGet{
		Validator: v,
	}

	return ret, nil
}

type AuthorizeGet struct {
	Service   Authorizer
	Validator *validator.Validate
}

// ServeHTTP authorizes a request that asking to access the resources belonging to a user
//
//	@summary		TBD
//	@description	TBD
//	@id				AuthorizeGet.ServeHTTP
//	@tags			authorization
//	@accept			json
//	@produce		json
//	@param			name		query		string	true	"TBD"
//	@param			password	query		string	true	"TBD"
//	@success		200			{string}	string
//	@failure		500			{object}	model.BadRequest
//	@failure		500			{object}	model.InternalServerError
//	@router			/v1/authorization [get]
func (p *AuthorizeGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := schema.NewDecoder()
	q := &model.AuthorizeRequest{}

	if err := decoder.Decode(q, r.URL.Query()); err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	if err := p.Validator.Struct(q); err != nil {
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.InvalidRequest,
		})
		return
	}

	location, err := p.Service.Authorize(r.Context(), q)

	if err != nil {
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.InvalidParameter,
		})
		return
	}

	// TODO: temporary
	http.Redirect(w, r, location, http.StatusFound)

	// TODO: Redirect unauthenticated user to the authentication endpoint
	// ...

	// TODO: Redirect authenticated user to the consent endpoint
	// ...
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
//	@tags			authorization
//	@accept			json
//	@produce		json
//	@param			name		formData	string	true	"TBD"
//	@param			password	formData	string	true	"TBD"
//	@success		200			{string}	string
//	@failure		500			{object}	model.BadRequest
//	@failure		500			{object}	model.InternalServerError
//	@router			/v1/authorization [post]
func (p *AuthorizePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NOT IMPLEMENTED YET
}
