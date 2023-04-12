package store

import (
	"context"
	"database/sql"
	"github.com/42milez/go-oidc-server/src/clock"
	"github.com/42milez/go-oidc-server/src/config"
)

func New(ctx context.Context, cfg *config.Config) (*sql.DB, func(), error) {
	return nil, func() {}, nil
}

type Repository struct {
	Clocker clock.Clocker
}
