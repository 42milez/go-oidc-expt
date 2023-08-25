package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/42milez/go-oidc-server/app/ent/typedef"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/rs/zerolog/log"
)

const (
	errFailedToEncodeHTTPResponse = "failed to encode http response"
	errFailedToWriteHTTPResponse  = "failed to write http response"
	errFailedToDecodeRequestBody  = "failed to decode request body"
	errValidationError            = "validation error"
)

type ErrResponse struct {
	Error   xerr.HTTPErr `json:"error"`
	Details []string     `json:"details,omitempty"`
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

func ResponseJsonWithInternalServerError(w http.ResponseWriter) {
	RespondJSON(w, http.StatusInternalServerError, &ErrResponse{
		Error: xerr.UnexpectedErrorOccurred,
	})
}

func GetUserID(ctx context.Context) (typedef.UserID, bool) {
	ret, ok := ctx.Value(UserIDKey{}).(typedef.UserID)
	return ret, ok
}
