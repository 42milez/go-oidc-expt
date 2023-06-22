package handler

import (
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/model"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

type ReadAdmin struct {
	Service ReadAdminService
}

func (p *ReadAdmin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	admin, err := p.Service.ReadAdmin(ctx)
	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
	}

	resp := model.AdminResponse{
		ID:   admin.ID,
		Name: admin.Name,
	}

	RespondJSON(w, http.StatusOK, resp)
}

type ReadAdmins struct {
	Service ReadAdminsService
}

func (p *ReadAdmins) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	admins, err := p.Service.ReadAdmins(ctx)
	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
	}

	resp := make([]model.AdminResponse, len(admins), len(admins))
	for i, admin := range admins {
		resp[i] = model.AdminResponse{
			ID:   admin.ID,
			Name: admin.Name,
		}
	}

	RespondJSON(w, http.StatusOK, resp)
}
