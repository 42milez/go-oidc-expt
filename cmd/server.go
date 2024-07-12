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
	"time"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	lis net.Listener
	srv *http.Server
}

func (s *Server) Run(appCtx context.Context) error {
	eg, egCtx := errgroup.WithContext(appCtx)

	shutdownServer := func() error {
		notifyCtx, stop := signal.NotifyContext(egCtx, os.Interrupt, syscall.SIGTERM)
		defer stop()

		<-notifyCtx.Done()
		log.Print(notifyCtx.Err())

		withTimeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := s.srv.Shutdown(withTimeoutCtx); err != nil {
			return err
		}

		log.Print("server stopped")

		return nil
	}

	eg.Go(func() error {
		return shutdownServer()
	})

	if err := s.srv.Serve(s.lis); (err != nil) && (!errors.Is(err, http.ErrServerClosed)) {
		return err
	}

	return eg.Wait()
}
