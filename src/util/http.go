package util

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func HttpClose(resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		log.Error().Err(err).Send()
	}
}
