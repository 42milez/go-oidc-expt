package handler

import (
	"encoding/json"
	"net/http"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/go-playground/validator/v10"
)

type SignIn struct {
	Service   SignInService
	Validator *validator.Validate
}

func (p *SignIn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		// TODO: Print error message with logger
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	err := p.Validator.Struct(body)

	if err != nil {
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.AuthenticationFailed,
		})
		return
	}

	token, err := p.Service.SignIn(ctx, body.Name, body.Password)

	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	resp := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	RespondJSON(w, http.StatusOK, resp)
}
