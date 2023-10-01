package api

import (
	"encoding/json"
	"net/http"

	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/api/oapigen"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/go-playground/validator/v10"
)

var registerUserHdlr *RegisterHdlr

func NewRegisterHdlr(option *HandlerOption) (*RegisterHdlr, error) {
	return &RegisterHdlr{
		service:   service.NewCreateUser(repository.NewUser(option.db, option.idGenerator)),
		session:   option.sessionRestorer,
		validator: option.validator,
	}, nil
}

type RegisterHdlr struct {
	service   UserCreator
	session   SessionRestorer
	validator *validator.Validate
}

func (rr *RegisterHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req oapigen.RegisterJSONRequestBody

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	if err := rr.validator.Struct(req); err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	_, err := rr.service.CreateUser(r.Context(), req.Name, req.Password)

	if err != nil {
		RespondJSON500(w, r, err)
		return
	}
}
