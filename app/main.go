package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/42milez/go-oidc-server/app/api"

	"github.com/42milez/go-oidc-server/app/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var version = "dev"

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Str("role", "idp").Logger()

	var cfg *config.Config
	var err error

	if cfg, err = config.New(); err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to parse env variable")
	}

	if err = run(context.Background(), cfg); err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to shutdown")
	}
}

func run(ctx context.Context, cfg *config.Config) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal().Stack().Err(err).Msgf("failed to listen port %d", cfg.Port)
	}

	log.Info().Msgf("listening on tcp://%s", lis.Addr().String())
	log.Info().Msgf("application starting in %s (version: %s)\n", cfg.Env, version)

	mux, cleanup, err := api.NewMux(ctx, cfg)

	if err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to build routes")
	}

	if cleanup != nil {
		defer cleanup()
	}

	srv := NewServer(lis, mux)

	return srv.Run(ctx)
}
