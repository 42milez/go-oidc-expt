package handler

import (
	"net/http"

	"github.com/42milez/go-oidc-server/src/entity"
)

type ReadUser struct {
	Service ReadUserService
}

func (p *ReadUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := p.Service.Read(ctx)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}

	resp := entity.UserResponse{
		ID:   user.ID,
		Name: user.Name,
	}

	RespondJSON(ctx, w, resp, http.StatusOK)
}

type ReadUsers struct {
	Service ReadUsersService
}

func (p *ReadUsers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := p.Service.ReadBulk(ctx)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}

	resp := make([]entity.UserResponse, len(users), len(users))
	for i, user := range users {
		resp[i] = entity.UserResponse{
			ID:   user.ID,
			Name: user.Name,
		}
	}

	RespondJSON(ctx, w, resp, http.StatusOK)
}
