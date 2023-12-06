package datastore

import (
	"context"
	"database/sql"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/ent/ent"
)

const (
	dbDialect               = "mysql"
	dbMaxOpenConnection     = 100
	dbMaxIdleConnection     = 10
	dbConnectionMaxLifetime = time.Hour
)

type Database struct {
	Client *ent.Client
	db     *sql.DB
}

func (d *Database) Ping(ctx context.Context) error {
	if err := d.db.PingContext(ctx); err != nil {
		return err
	}
	return nil
}
