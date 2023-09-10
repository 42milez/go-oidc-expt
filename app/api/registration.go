package api

import (
	"encoding/json"
	"net/http"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/rs/zerolog/log"

	"github.com/go-playground/validator/v10"
)

type RegisterHdlr struct {
	service   UserCreator
	session   SessionRestorer
	validator *validator.Validate
}

func (ru *RegisterHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req RegisterJSONRequestBody

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	if err := ru.validator.Struct(req); err != nil {
		log.Error().Err(err).Msg(errValidationError)
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.InvalidRequest,
		})
		return
	}

	_, err := ru.service.CreateUser(r.Context(), req.Name, req.Password)

	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}
}
