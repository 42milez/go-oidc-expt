package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Authorization struct {
	Service   Authorizer
	Validator *validator.Validate
}

func (p *Authorization) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NOT IMPLEMENTED YET
}
