package api

import (
	"encoding/json"
	"net/http"

	"github.com/42milez/go-oidc-expt/cmd/iface"
	"github.com/42milez/go-oidc-expt/cmd/option"
	"github.com/42milez/go-oidc-expt/cmd/repository"
	"github.com/42milez/go-oidc-expt/cmd/service"

	"github.com/42milez/go-oidc-expt/pkg/xerr"
)

var registration *Registration

func InitRegistration(opt *option.Option) {
	if registration == nil {
		registration = &Registration{
			svc: service.NewRegisterUser(repository.NewUser(opt)),
			v:   opt.V,
		}
	}
}

type Registration struct {
	svc UserRegisterer
	v   iface.StructValidator
}

func (rr *Registration) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	RespondJSON(w, r, http.StatusOK, nil, &User{
		ID:   u.ID(),
		Name: u.Name(),
	})
}
