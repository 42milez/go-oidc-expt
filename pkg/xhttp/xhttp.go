package xhttp

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func CloseHTTPConn(resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		log.Error().Stack().Err(err).Send()
	}
}
