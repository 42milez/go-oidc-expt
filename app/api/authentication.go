package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/42milez/go-oidc-server/app/api/oapigen"

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

func (ah *AuthenticateHdlr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var reqBody oapigen.AuthenticateJSONRequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		RespondJSON500(w, err)
		return
	}

	if err := ah.validator.Struct(reqBody); err != nil {
		RespondJSON400(w, xerr.InvalidRequest, nil, err)
		return
	}

	userID, err := ah.service.Authenticate(r.Context(), reqBody.Name, reqBody.Password)

	if err != nil {
		if errors.Is(err, xerr.UserNotFound) {
			RespondJSON401(w, xerr.InvalidUsernameOrPassword, nil, err)
			return
		} else if errors.Is(err, xerr.PasswordNotMatched) {
			RespondJSON401(w, xerr.InvalidUsernameOrPassword, nil, err)
			return
		} else {
			RespondJSON500(w, err)
			return
		}
	}

	sessionID, err := ah.session.Create(r.Context(), &entity.Session{
		UserID: userID,
	})

	if err != nil {
		RespondJSON500(w, err)
		return
	}

	if err = ah.cookie.Write(w, sessionIDCookieName, sessionID, config.SessionIDCookieTTL); err != nil {
		RespondJSON500(w, err)
		return
	}

	Redirect(w, r, config.ConsentURL, http.StatusFound)
}
