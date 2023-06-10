package util

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func CloseHTTPConn(resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		log.Error().Err(err).Msg("failed to close http connection")
	}
}
