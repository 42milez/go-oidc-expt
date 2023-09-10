package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/go-playground/validator/v10"
)

type AuthenticateHdlr struct {
	service   Authenticator
	cookie    CookieWriter
	session   SessionCreator
	validator *validator.Validate
}

const sessionIDCookieName = config.SessionIDCookieName

func (au *AuthenticateHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request, params *AuthenticateParams) {
	var reqBody AuthenticateJSONRequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Error().Err(err).Msg(errFailedToDecodeRequestBody)
		RespondJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	if err := au.validator.Struct(reqBody); err != nil {
		log.Error().Err(err).Msg(errValidationError)
		RespondJSON(w, http.StatusBadRequest, &ErrResponse{
			Error: xerr.InvalidRequest,
		})
		return
	}

	userID, err := au.service.Authenticate(r.Context(), reqBody.Name, reqBody.Password)

	if err != nil {
		if errors.Is(err, xerr.UserNotFound) {
			RespondJSON(w, http.StatusUnauthorized, &ErrResponse{
				Error: xerr.InvalidUsernameOrPassword,
			})
			return
		} else if errors.Is(err, xerr.PasswordNotMatched) {
			RespondJSON(w, http.StatusUnauthorized, &ErrResponse{
				Error: xerr.InvalidUsernameOrPassword,
			})
			return
		} else {
			RespondJson500(w, xerr.UnexpectedErrorOccurred)
			return
		}
	}

	sessionID, err := au.session.Create(r.Context(), &entity.Session{
		UserID: userID,
	})

	if err != nil {
		RespondJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	if err = au.cookie.Write(w, sessionIDCookieName, sessionID, config.SessionIDCookieTTL); err != nil {
		RespondJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	Redirect(w, r, config.ConsentURL, http.StatusFound)
}
