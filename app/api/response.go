package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
)

type ErrResponse struct {
	Status  int            `json:"status"`
	Summary xerr.PublicErr `json:"error"`
	Details []string       `json:"details,omitempty"`
}

func RespondJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		resp := ErrResponse{
			Summary: xerr.UnexpectedErrorOccurred,
		}

		if err = json.NewEncoder(w).Encode(resp); err != nil {
			appLogger.Error().Err(err)
		}

		return
	}

	w.WriteHeader(statusCode)

	if _, err = fmt.Fprintf(w, "%s", bodyBytes); err != nil {
		appLogger.Error().Err(err)
	}
}

func RespondJSON200(w http.ResponseWriter) {
	RespondJSON(w, http.StatusOK, &ErrResponse{
		Status:  http.StatusOK,
		Summary: xerr.OK,
	})
}

func RespondJSON400(w http.ResponseWriter, summary xerr.PublicErr, details []string, err error) {
	body := &ErrResponse{
		Status:  http.StatusBadRequest,
		Summary: summary,
	}
	if details != nil {
		body.Details = details
	}
	if err != nil {
		appLogger.Error().Err(err).Send()
	}
	RespondJSON(w, http.StatusBadRequest, body)
}

func RespondJSON401(w http.ResponseWriter, summary xerr.PublicErr, details []string, err error) {
	body := &ErrResponse{
		Status:  http.StatusUnauthorized,
		Summary: summary,
	}
	if details != nil {
		body.Details = details
	}
	if err != nil {
		appLogger.Error().Err(err).Send()
	}
	RespondJSON(w, http.StatusUnauthorized, body)
}

func RespondJSON500(w http.ResponseWriter, err error) {
	if err != nil {
		appLogger.Error().Err(err).Send()
	}
	RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
		Status:  http.StatusInternalServerError,
		Summary: xerr.UnexpectedErrorOccurred,
	})
}

func RespondJSON503(w http.ResponseWriter, err error) {
	if err != nil {
		appLogger.Error().Err(err).Send()
	}
	RespondJSON(w, http.StatusServiceUnavailable, &ErrResponse{
		Status:  http.StatusServiceUnavailable,
		Summary: xerr.ServiceTemporaryUnavailable,
	})
}

func Redirect(w http.ResponseWriter, r *http.Request, u string, code int) {
	redirectURL, err := url.Parse(u)

	if err != nil {
		RespondJSON500(w, err)
		return
	}

	if !xutil.IsEmpty(r.URL.RawQuery) {
		redirectURL, err = url.Parse(fmt.Sprintf("%s&%s", redirectURL, r.URL.RawQuery))
		if err != nil {
			RespondJSON500(w, err)
			return
		}
	}

	http.Redirect(w, r, redirectURL.String(), code)
}
