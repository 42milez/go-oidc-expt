package handler

import (
	"github.com/42milez/go-oidc-server/app/idp/session"
	"net/http"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

func RestoreSession(sess *session.Util) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req, err := sess.FillContext(r)
			if err != nil {
				RespondJSON(w, http.StatusUnauthorized, &ErrResponse{
					Error: xerr.UnauthorizedUser,
				})
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}
