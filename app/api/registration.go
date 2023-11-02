package api

import (
	"encoding/json"
	"net/http"

	"github.com/42milez/go-oidc-server/app/option"

	"github.com/42milez/go-oidc-server/app/iface"

	"github.com/42milez/go-oidc-server/app/repository"
	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

var registerUserHdlr *RegisterHdlr

func NewRegisterHdlr(opt *option.Option) *RegisterHdlr {
	return &RegisterHdlr{
		svc: service.NewRegisterUser(repository.NewUser(opt.DB, opt.IdGen)),
		v:   opt.V,
	}
}

type RegisterHdlr struct {
	svc UserRegisterer
	v   iface.StructValidator
}

func (rr *RegisterHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req RegisterJSONRequestBody

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondJSON500(w, r, err)
		return
	}

	if err := rr.v.Struct(req); err != nil {
		RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
		return
	}

	u, err := rr.svc.RegisterUser(r.Context(), req.Name, req.Password)
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	RespondJSON(w, r, http.StatusOK, &User{
		ID:   u.ID(),
		Name: u.Name(),
	})
}
