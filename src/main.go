package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/42milez/go-oidc-server/src/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var version = "dev"

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Str("role", "app").Logger()

	if err := run(context.Background()); err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to shutdown")
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to parse env variable")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal().Stack().Err(err).Msgf("failed to listen port %d", cfg.Port)
	}

	log.Info().Msgf("application starting in %s (version: %s)\n", cfg.Env, version)
	log.Info().Msgf("listening on tcp://%s", lis.Addr().String())

	mux, cleanup, err := NewMux(ctx, cfg)
	defer cleanup()
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to build routes")
	}

	srv := NewServer(lis, mux)

	return srv.Run(ctx)
}
