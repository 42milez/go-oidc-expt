package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/42milez/go-oidc-server/app/idp/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var version = "dev"

//	@title						go-oidc-server
//	@version					1.0
//	@description				TBD
//	@tag.name					TBD
//	@tag.description			TBD
//	@tag.docs.url				TBD
//	@tag.docs.description		TBD
//	@termsOfService				TBD
//	@contact.name				TBD
//	@contact.url				TBD
//	@contact.email				TBD
//	@license.name				MIT
//	@license.url				TBD
//	@host						TBD
//	@BasePath					/v1
//	@accept						json
//	@produce					json
//	@query.collection.format	TBD
//	@schemes					http https
//	@externalDocs.description	TBD
//	@externalDocs.url			TBD

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Str("role", "idp").Logger()

	cfg, err := config.New()

	if err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to parse env variable")
	}

	if cfg.CI {
		log.Logger.Level(zerolog.Disabled)
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

	mux, cleanup, err := NewMux(ctx, cfg)

	if err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to build routes")
	}

	if cleanup != nil {
		defer cleanup()
	}

	srv := NewServer(lis, mux)

	return srv.Run(ctx)
}
