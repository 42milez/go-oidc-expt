package api

import (
	"encoding/json"
	"net/http"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/rs/zerolog/log"

	"github.com/go-playground/validator/v10"
)

type RegisterUserHdlr struct {
	service   UserCreator
	session   SessionRestorer
	validator *validator.Validate
}

func (p *RegisterUserHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req RegisterUserJSONRequestBody

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	if err := p.validator.Struct(req); err != nil {
		log.Error().Err(err).Msg(errValidationError)
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.InvalidRequest,
		})
		return
	}

	_, err := p.service.CreateUser(r.Context(), req.Name, req.Password)

	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}
}

type SelectUser struct {
	Service UserSelector
}

func (p *SelectUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := p.Service.SelectUser(ctx)

	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
	}

	resp := User{
		Id:   uint64(user.ID),
		Name: user.Name,
	}

	RespondJSON(w, http.StatusOK, resp)
}
