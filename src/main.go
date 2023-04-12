package main

import (
	"context"
	"fmt"
	"github.com/42milez/go-oidc-server/src/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net"
	"os"
)

var version = "dev"

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Str("role", "app").Logger()

	if err := run(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed to shutdown")
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()

	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse env variable")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))

	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen port %d", cfg.Port)
	}

	log.Info().Msgf("application starting in %s", cfg.Env)
	log.Info().Msgf("listening on tcp://%s:%d", lis.Addr().String(), cfg.Port)

	mux, cleanup, err := NewMux(ctx, cfg)
	defer cleanup()

	if err != nil {
		log.Fatal().Err(err).Msg("failed to build routes")
	}

	srv := NewServer(lis, mux)

	return srv.Run(ctx)
}
