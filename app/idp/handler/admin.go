package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/42milez/go-oidc-server/app/idp/model"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

type Create struct {
	Service   AdminCreateService
	Validator *validator.Validate
}

func (p *Create) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req model.AdminCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	if err := p.Validator.Struct(req); err != nil {
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.AuthenticationFailed,
		})
		return
	}

	_, err := p.Service.Create(r.Context(), req.Name, req.Password)

	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}
}

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
