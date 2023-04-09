package main

import (
	"context"
	"fmt"
	"github.com/42milez/go-oidc-server/src/config"
	"os"
)

var version string

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(version)
}

func run(ctx context.Context) error {
	_, err := config.New()
	if err != nil {
		return err
	}
	return nil
}
