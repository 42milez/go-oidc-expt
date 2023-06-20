package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type SignIn struct {
	Service SignInService
	Validator *validator.Validate
}

func (p *SignIn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body struct {
		Name string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: "",
		}, http.StatusInternalServerError)
		return
	}

	err := p.Validator.Struct(body)

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: "",
		}, http.StatusBadRequest)
		return
	}

	token, err := p.Service.SignIn(ctx, body.Name, body.Password)

	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: "",
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
