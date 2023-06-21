package handler

import (
	"encoding/json"
	"github.com/42milez/go-oidc-server/pkg/xerr"
	"net/http"

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
		RespondJSON(ctx, w, &ErrResponse{
			// TODO: change the type of Message from string to error
			Message: xerr.ErrInternalServerError.Error(),
		}, http.StatusInternalServerError)
		return
	}

	err := p.Validator.Struct(body)

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: xerr.ErrFailedToAuthenticate.Error(),
		}, http.StatusBadRequest)
		return
	}

	token, err := p.Service.SignIn(ctx, body.Name, body.Password)

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: xerr.ErrInternalServerError.Error(),
		}, http.StatusInternalServerError)
		return
	}

	resp := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	RespondJSON(r.Context(), w, resp, http.StatusOK)
}
