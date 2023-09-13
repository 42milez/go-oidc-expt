package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func NewServer(lis net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		lis: lis,
	}
}

type Server struct {
	srv *http.Server
	lis net.Listener
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.srv.Serve(s.lis); (err != nil) && (!errors.Is(err, http.ErrServerClosed)) {
			log.Error().Err(err).Msg("failed to close")
			return err
		}
		return nil
	})

	<-ctx.Done()

	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Error().Err(err).Msg("failed to shutdown")
	}

	return eg.Wait()
}
