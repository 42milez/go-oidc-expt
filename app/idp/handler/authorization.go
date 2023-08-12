package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

func NewAuthorizeGet() *AuthorizeGet {
	return &AuthorizeGet{}
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
//	@param			name		body		string	true	"TBD"
//	@param			password	body		string	true	"TBD"
//	@success		200			{object}	model.Authenticate
//	@failure		500			{object}	model.BadRequest
//	@failure		500			{object}	model.InternalServerError
//	@router			/v1/authorization [get]
func (p *AuthorizeGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NOT IMPLEMENTED YET
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
//	@param			name		body		string	true	"TBD"
//	@param			password	body		string	true	"TBD"
//	@success		200			{object}	model.Authenticate
//	@failure		500			{object}	model.BadRequest
//	@failure		500			{object}	model.InternalServerError
//	@router			/v1/authorization [post]
func (p *AuthorizePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NOT IMPLEMENTED YET
}

// authorizationCore
//func authorizationCore(w http.ResponseWriter, r *http.Request) {
//	// NOT IMPLEMENTED YET
//}
