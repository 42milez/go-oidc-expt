package xutil

import (
	"net/http"

	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/rs/zerolog/log"
)

type ClosableClient interface {
	Close() error
}

func CloseConnection(client ClosableClient) {
	if client == nil {
		return
	}
	if err := client.Close(); err != nil {
		log.Error().Err(err).Msg(xerr.FailedToCloseConnection.Error())
	}
}

func CloseHTTPConn(resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		log.Error().Err(err).Msg("failed to close http connection")
	}
}
