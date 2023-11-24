package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	lis net.Listener
	srv *http.Server
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := s.srv.Serve(s.lis); (err != nil) && (!errors.Is(err, http.ErrServerClosed)) {
			log.Fatal(err)
			return err
		}
		return nil
	})

	<-ctx.Done()

	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}

	return eg.Wait()
}
