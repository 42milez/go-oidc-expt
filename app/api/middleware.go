package api

import (
	"errors"
	"net/http"

	"github.com/42milez/go-oidc-server/app/service"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"
)

func RestoreSession(cookie *service.Cookie, sess SessionRestorer) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sid, err := cookie.Read(r, config.SessionIDCookieName)
			if errors.Is(err, http.ErrNoCookie) {
				next.ServeHTTP(w, r)
				return
			}
			req, err := sess.Restore(r, typedef.SessionID(sid))
			if err != nil {
				RespondJson500(w, xerr.UnexpectedErrorOccurred)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}
