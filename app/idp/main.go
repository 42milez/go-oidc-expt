package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/42milez/go-oidc-server/app/idp/api"
	"github.com/42milez/go-oidc-server/app/idp/config"

	"github.com/rs/zerolog/pkgerrors"

	"github.com/rs/zerolog"
)

var Version = "dev"

func NewServer(lis net.Listener, mux http.Handler) *Server {
	return &Server{
		lis: lis,
		srv: &http.Server{Handler: mux},
	}
}

func NewBaseLogger(cfg *config.Config) *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixNano
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	ret := zerolog.New(os.Stdout).Level(cfg.LogLevel).With().Timestamp().Str("env", cfg.Env).
		Str("service", config.AppName).Logger()
	return &ret
}

func Run(ctx context.Context, cfg *config.Config, logger *zerolog.Logger) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))

	if err != nil {
		log.Fatalf("failed to listen port %d", cfg.Port)
	}

	log.Printf("listening on tcp://%s", lis.Addr().String())
	log.Printf("application starting in %s (Version: %s)\n", cfg.Env, Version)

	mux, cleanup, err := api.NewMux(ctx, cfg, logger)

	if cleanup != nil {
		defer cleanup()
	}

	if err != nil {
		log.Fatalf("failed to initialize mux: %s", err)
	}

	srv := NewServer(lis, mux)

	return srv.Run(ctx)
}

func main() {
	var cfg *config.Config
	var err error

	if cfg, err = config.New(); err != nil {
		log.Fatalf("failed to get config values: %s", err)
	}

	baseLogger := NewBaseLogger(cfg)

	if err = Run(context.Background(), cfg, baseLogger); err != nil {
		log.Fatalf("failed to run server: %s", err)
	}
}
