package datastore

import (
	"context"
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
	"github.com/cenkalti/backoff/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

func NewMySQL(ctx context.Context, cfg *config.Config) (*Database, error) {
	dataSrc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", cfg.DBAdmin, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open(dialect.MySQL, dataSrc)

	if err != nil {
		xutil.CloseConnection(db)
		return nil, err
	}

	pingOp := func() error {
		return db.PingContext(ctx)
	}

	if err = ping(ctx, pingOp); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(dbConnectionMaxLifetime)
	db.SetMaxIdleConns(dbMaxIdleConnection)
	db.SetMaxOpenConns(dbMaxOpenConnection)

	drv := entsql.OpenDB(dbDialect, db)
	client := ent.NewClient(ent.Driver(drv))

	if cfg.Debug {
		client = client.Debug()
	}

	return &Database{
		Client: client,
		db:     db,
	}, nil
}

func NewRedis(ctx context.Context, cfg *config.Config) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})

	pingOp := func() error {
		return client.Ping(ctx).Err()
	}

	if err := ping(ctx, pingOp); err != nil {
		xutil.CloseConnection(client)
		return nil, err
	}

	return &Cache{
		Client: client,
	}, nil
}

func ping(ctx context.Context, op func() error) error {
	b := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
	if err := backoff.Retry(op, b); err != nil {
		return err
	}
	return nil
}
