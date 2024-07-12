package xutil

import (
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
		log.Error().Stack().Err(err).Send()
	}
}

func IsEmpty[T string | []byte](v T) bool {
	return len(v) == 0
}

func NewFalse() *bool {
	ret := false
	return &ret
}
