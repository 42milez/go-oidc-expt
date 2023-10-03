package api

import (
	"github.com/rs/zerolog"
	"testing"
)

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	m.Run()
}
