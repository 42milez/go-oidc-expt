package api

import (
	"encoding/json"
	"net/http"

	"github.com/42milez/go-oidc-server/app/api/oapigen"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/go-playground/validator/v10"
)

type RegisterHdlr struct {
	service   UserCreator
	session   SessionRestorer
	validator *validator.Validate
}

func (rh *RegisterHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req oapigen.RegisterJSONRequestBody

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondJSON500(w, err)
		return
	}

	if err := rh.validator.Struct(req); err != nil {
		RespondJSON400(w, xerr.InvalidRequest, nil, err)
		return
	}

	_, err := rh.service.CreateUser(r.Context(), req.Name, req.Password)

	if err != nil {
		RespondJSON500(w, err)
		return
	}
}
