package api

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/httpstore"
	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/rs/zerolog"

	"github.com/go-chi/chi/v5/middleware"
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

func RestoreSession(opt *option.Option) MiddlewareFunc {
	rs := httpstore.NewCache(opt)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sid, err := opt.Cookie.Read(r, config.SessionIDCookieName)
			if errors.Is(err, http.ErrNoCookie) {
				next.ServeHTTP(w, r)
				return
			}
			sidUint64, err := strconv.ParseUint(sid, 10, 64)
			if err != nil {
				RespondJSON400(w, r, xerr.InvalidRequest, nil, err)
				return
			}
			req, err := rs.Restore(r, typedef.SessionID(sidUint64))
			if err != nil {
				RespondJSON500(w, r, err)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}

func InjectRequestParameter() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req, err := injectRequestParameter(r)
			if err != nil {
				RespondServerError(w, r, err)
			}
			next.ServeHTTP(w, req)
		})
	}
}
