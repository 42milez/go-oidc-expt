package api

import (
	"errors"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"net/http"

	"github.com/42milez/go-oidc-server/app/api/cookie"
	"github.com/42milez/go-oidc-server/app/api/session"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"
)

func RestoreSession(ck *cookie.Cookie, sess *session.Session) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sid, err := ck.Get(r, config.SessionIDCookieName)

			if errors.Is(err, http.ErrNoCookie) {
				next.ServeHTTP(w, r)
				return
			}

			req, err := sess.Restore(r, typedef.SessionID(sid))

			if err != nil {
				ResponseJson500(w, xerr.UnexpectedErrorOccurred)
				return
			}

			next.ServeHTTP(w, req)
		})
	}
}
