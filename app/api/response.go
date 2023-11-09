package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
)

type Response struct {
	Status  int            `json:"status"`
	Summary xerr.PublicErr `json:"summary"`
	Details []string       `json:"details,omitempty"`
}

func RespondJSON(w http.ResponseWriter, r *http.Request, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		LogError(r, err, nil)

		w.WriteHeader(http.StatusInternalServerError)

		resp := Response{
			Status:  http.StatusInternalServerError,
			Summary: xerr.UnexpectedErrorOccurred,
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
	RespondJSON(w, r, http.StatusOK, &Response{
		Status:  http.StatusOK,
		Summary: xerr.OK,
	})
}

func RespondJSON400(w http.ResponseWriter, r *http.Request, summary xerr.PublicErr, details []string, err error) {
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
	RespondJSON(w, r, http.StatusBadRequest, body)
}

func RespondTokenRequestError(w http.ResponseWriter, err xerr.TokenRequestErr) {
	body := &struct {
		Error xerr.TokenRequestErr `json:"error,string"`
	}{
		Error: err,
	}
	RespondJSON(w, nil, http.StatusBadRequest, body)
}

func RespondJSON401(w http.ResponseWriter, r *http.Request, summary xerr.PublicErr, details []string, err error) {
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
	RespondJSON(w, r, http.StatusUnauthorized, body)
}

func RespondJSON404(w http.ResponseWriter) {
	RespondJSON(w, nil, http.StatusNotFound, nil)
}

func RespondJSON500(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		e := errors.WithStack(err)
		appLogger.Error().Stack().Err(e).Send()
	}
	RespondJSON(w, r, http.StatusInternalServerError, &Response{
		Status:  http.StatusInternalServerError,
		Summary: xerr.UnexpectedErrorOccurred,
	})
}

func RespondJSON503(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		appLogger.Error().Err(err).Send()
	}
	RespondJSON(w, r, http.StatusServiceUnavailable, &Response{
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
