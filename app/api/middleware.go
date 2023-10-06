package api

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog"

	"github.com/42milez/go-oidc-server/app/api/oapigen"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"
)

var accessLogger zerolog.Logger

func AccessLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		reqBodyBuf, err := io.ReadAll(r.Body)

		if err != nil {
			RespondJSON500(w, r, err)
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(reqBodyBuf))

		respBodyBuf := &bytes.Buffer{}
		respWriter := middleware.NewWrapResponseWriter(w, 1)
		respWriter.Tee(respBodyBuf)

		defer func() {
			duration, _ := strconv.ParseFloat(fmt.Sprintf("%f", time.Since(startTime).Seconds()), 64)
			accessLogger.Info().
				Str("request_id", middleware.GetReqID(r.Context())).
				Str("http_method", r.Method).
				Str("path", r.URL.String()).
				Str("remote_addr", r.RemoteAddr).
				Str("user_agent", r.UserAgent()).
				Str("request", string(reqBodyBuf)).
				Int("status", respWriter.Status()).
				Str("response", respBodyBuf.String()).
				Float64("duration", duration).
				Send()
		}()

		next.ServeHTTP(w, r)
	})
}

func RestoreSession(option *HandlerOption) oapigen.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sid, err := option.cookie.Read(r, config.SessionIDCookieName)
			if errors.Is(err, http.ErrNoCookie) {
				next.ServeHTTP(w, r)
				return
			}
			req, err := option.sessionRestorer.Restore(r, typedef.SessionID(sid))
			if err != nil {
				RespondJSON500(w, r, err)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}
