package middleware

import (
	"errors"
	"net/http"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/cookie"
	"github.com/42milez/go-oidc-server/app/handler"
	"github.com/42milez/go-oidc-server/app/session"
	"github.com/42milez/go-oidc-server/pkg/xerr"
)

func RestoreSession(ck *cookie.Cookie, sess *session.Session) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionID, err := ck.Get(r, config.SessionIDCookieName)

			if errors.Is(err, http.ErrNoCookie) {
				next.ServeHTTP(w, r)
				return
			}

			req, err := sess.Restore(r, sessionID)

			if err != nil {
				handler.RespondJSON(w, http.StatusUnauthorized, &handler.ErrResponse{
					Error: xerr.UnauthorizedUser,
				})
				return
			}

			next.ServeHTTP(w, req)
		})
	}
}
