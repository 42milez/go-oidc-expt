package api

//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config ../../openapi.yml spec/spec.yml

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"

	"github.com/rs/zerolog/log"
)

const (
	errFailedToEncodeHTTPResponse = "failed to encode http response"
	errFailedToWriteHTTPResponse  = "failed to write http response"
	errFailedToDecodeRequestBody  = "failed to decode request body"
	errValidationError            = "validation error"
)

type ErrResponse struct {
	Error   xerr.PublicErr `json:"error"`
	Details []string       `json:"details,omitempty"`
}

func RespondJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		log.Error().Err(err).Msg(errFailedToEncodeHTTPResponse)

		w.WriteHeader(http.StatusInternalServerError)

		resp := ErrResponse{
			Error: xerr.UnexpectedErrorOccurred,
		}

		if err = json.NewEncoder(w).Encode(resp); err != nil {
			log.Error().Err(err).Msg(errFailedToWriteHTTPResponse)
		}

		return
	}

	w.WriteHeader(statusCode)

	if _, err = fmt.Fprintf(w, "%s", bodyBytes); err != nil {
		log.Error().Err(err).Msg(errFailedToWriteHTTPResponse)
	}
}

func ResponseJson500(w http.ResponseWriter, msg xerr.PublicErr) {
	RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
		Error: msg,
	})
}

func Redirect(w http.ResponseWriter, r *http.Request, u string, code int) {
	redirectURL, err := url.Parse(u)

	if err != nil {
		ResponseJson500(w, xerr.UnexpectedErrorOccurred)
		return
	}

	if !xutil.IsEmpty(r.URL.RawQuery) {
		redirectURL, err = url.Parse(fmt.Sprintf("%s&%s", redirectURL, r.URL.RawQuery))
		if err != nil {
			ResponseJson500(w, xerr.UnexpectedErrorOccurred)
			return
		}
	}

	http.Redirect(w, r, redirectURL.String(), code)
}
