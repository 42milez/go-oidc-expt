package main

import (
	"context"
	"fmt"
	"github.com/42milez/go-oidc-server/src/util"
	"golang.org/x/sync/errgroup"
	"io"
	"net"
	"net/http"
	"testing"
)

func TestServer_Run(t *testing.T) {
	// Bind dynamic port
	// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports
	lis, err := net.Listen("tcp", "localhost:0")

	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)
	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "Path requested: %s", r.URL.Path[1:]); err != nil {
			t.Error(err)
		}
	})

	eg.Go(func() error {
		s := NewServer(lis, mux)
		return s.Run(ctx)
	})

	path := "test"
	url := fmt.Sprintf("http://%s/%s", lis.Addr().String(), path)
	t.Logf("try request to %q", url)

	resp, err := http.Get(url)
	defer util.HttpClose(resp)

	if err != nil {
		t.Fatalf("failed to get: %+v", err)
	}

	got, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Fatalf("failed to read body: %+v", err)
	}

	want := fmt.Sprintf("Path requested: %s", path)

	if string(got) != want {
		t.Errorf("got = %q; want = %q", got, want)
	}

	cancel()

	if err := eg.Wait(); err != nil {
		t.Fatalf("failed to shutdown: %+v", err)
	}
}
