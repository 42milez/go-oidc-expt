package handler

import (
	"errors"
	"net/http"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/handler/cookie"
	"github.com/42milez/go-oidc-server/app/handler/session"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"
)

func RestoreSession(ck *cookie.Cookie, sess *session.Session) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sid, err := ck.Get(r, config.SessionIDCookieName)

			if errors.Is(err, http.ErrNoCookie) {
				RespondJSON(w, http.StatusUnauthorized, &ErrResponse{
					Error: xerr.UnauthorizedRequest,
				})
				return
			}

			req, err := sess.Restore(r, typedef.SessionID(sid))

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			next.ServeHTTP(w, req)
		})
	}
}
