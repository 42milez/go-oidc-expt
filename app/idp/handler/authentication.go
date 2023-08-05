package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/auth"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/app/idp/service"
	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/rs/zerolog/log"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/go-playground/validator/v10"
)

func NewAuthenticate(entClient *ent.Client) (*Authenticate, error) {
	jwtUtil, err := auth.NewJWTUtil(&xutil.RealClocker{})

	if err != nil {
		return nil, fmt.Errorf("%w: %w", xerr.FailedToInitialize, err)
	}

	return &Authenticate{
		Service: &service.Authenticate{
			Repo: &repository.User{
				Clock: &xutil.RealClocker{},
				DB:    entClient,
			},
			TokenGenerator: jwtUtil,
		},
		Validator: validator.New(),
	}, nil
}

type Authenticate struct {
	Service   Authenticator
	Validator *validator.Validate
}

func (p *Authenticate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Error().Err(err).Msg(errFailedToDecodeRequestBody)
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	if err := p.Validator.Struct(body); err != nil {
		log.Error().Err(err).Msg(errValidationError)
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.AuthenticationFailed,
		})
		return
	}

	token, err := p.Service.Authenticate(ctx, body.Name, body.Password)

	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		})
		return
	}

	resp := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	RespondJSON(w, http.StatusOK, resp)
}
