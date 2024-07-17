package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/42milez/go-oidc-expt/cmd/config"

	"github.com/pkg/errors"

	"github.com/42milez/go-oidc-expt/pkg/xerr"
	"github.com/42milez/go-oidc-expt/pkg/xutil"
)

type Response struct {
	Status  int              `json:"status"`
	Summary xerr.PublicError `json:"summary"`
	Details []string         `json:"details,omitempty"`
}

func RespondJSON(w http.ResponseWriter, r *http.Request, statusCode int, headers map[string]string, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	for k, v := range headers {
		w.Header().Set(k, v)
	}

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		LogError(r, err, nil)

		w.WriteHeader(http.StatusInternalServerError)

		resp := Response{
			Status:  http.StatusInternalServerError,
			Summary: xerr.UnexpectedErrorOccurred2,
		}

		if err = json.NewEncoder(w).Encode(resp); err != nil {
			LogError(r, err, nil)
		}

		return
	}

	w.WriteHeader(statusCode)

	if _, err = fmt.Fprintf(w, "%s", bodyBytes); err != nil {
		LogError(r, err, nil)
	}
}

func RespondJSON200(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, r, http.StatusOK, nil, &Response{
		Status:  http.StatusOK,
		Summary: xerr.OK,
	})
}

func RespondJSON400(w http.ResponseWriter, r *http.Request, summary xerr.PublicError, details []string, err error) {
	body := &Response{
		Status:  http.StatusBadRequest,
		Summary: summary,
	}
	if details != nil {
		body.Details = details
	}
	if err != nil {
		appLogger.Error().Err(err).Send()
	}
	RespondJSON(w, r, http.StatusBadRequest, nil, body)
}

func RespondJSON401(w http.ResponseWriter, r *http.Request, summary xerr.PublicError, details []string, err error) {
	body := &Response{
		Status:  http.StatusUnauthorized,
		Summary: summary,
	}
	if details != nil {
		body.Details = details
	}
	if err != nil {
		appLogger.Error().Err(err).Send()
	}
	RespondJSON(w, r, http.StatusUnauthorized, nil, body)
}

func RespondJSON404(w http.ResponseWriter) {
	RespondJSON(w, nil, http.StatusNotFound, nil, nil)
}

func RespondJSON500(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		e := errors.WithStack(err)
		appLogger.Error().Stack().Err(e).Send()
	}
	RespondJSON(w, r, http.StatusInternalServerError, nil, &Response{
		Status:  http.StatusInternalServerError,
		Summary: xerr.UnexpectedErrorOccurred2,
	})
}

func RespondJSON503(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		appLogger.Error().Err(err).Send()
	}
	RespondJSON(w, r, http.StatusServiceUnavailable, nil, &Response{
		Status:  http.StatusServiceUnavailable,
		Summary: xerr.ServiceTemporaryUnavailable,
	})
}

func Redirect(w http.ResponseWriter, r *http.Request, path string, code int) {
	cfg, err := config.New()
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	redirectURL, err := url.Parse(cfg.IdpHost + path)
	if err != nil {
		RespondJSON500(w, r, err)
		return
	}

	if !xutil.IsEmpty(r.URL.RawQuery) {
		redirectURL, err = url.Parse(fmt.Sprintf("%s?%s", redirectURL, r.URL.RawQuery))
		if err != nil {
			RespondJSON500(w, r, err)
			return
		}
	}

	http.Redirect(w, r, redirectURL.String(), code)
}

type OIDCError struct {
	Error            xerr.OIDCError `json:"error,string"`
	ErrorDescription string         `json:"error_description,omitempty"`
	ErrorUri         string         `json:"error_uri,omitempty"`
}

func RespondAuthorizationRequestError(w http.ResponseWriter, r *http.Request, redirectUri, state string, err xerr.OIDCError) {
	uri, e := url.Parse(redirectUri)
	if e != nil {
		RespondServerError(w, r, err)
	}
	q := uri.Query()
	q.Set("error", err.Error())
	q.Set("state", state)
	uri.RawQuery = q.Encode()
	http.Redirect(w, r, uri.String(), http.StatusFound)
}

func RespondTokenRequestError(w http.ResponseWriter, r *http.Request, err xerr.OIDCError) {
	body := &OIDCError{
		Error: err,
	}
	RespondJSON(w, r, http.StatusBadRequest, nil, body)
}

func RespondServerError(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		e := errors.WithStack(err)
		appLogger.Error().Stack().Err(e).Send()
	}
	body := &struct {
		Error xerr.OIDCError `json:"error,string"`
	}{
		Error: xerr.ServerError,
	}
	RespondJSON(w, r, http.StatusInternalServerError, nil, body)
}
