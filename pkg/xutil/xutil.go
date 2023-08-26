package xutil

import (
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

func IsEmpty[T string | []byte](v T) bool {
	return len(v) == 0
}
