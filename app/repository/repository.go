package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/ent/ent"

	"github.com/42milez/go-oidc-server/pkg/xutil"
	"github.com/redis/go-redis/v9"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDialect               = "mysql"
	dbMaxOpenConnection     = 100
	dbMaxIdleConnection     = 10
	dbConnectionMaxLifetime = time.Hour
)

func NewDBClient(ctx context.Context, cfg *config.Config) (*sql.DB, error) {
	dataSrc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", cfg.DBAdmin, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open(dialect.MySQL, dataSrc)

	if err != nil {
		xutil.CloseConnection(db)
		return nil, xerr.FailToEstablishConnection.Wrap(err)
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(dbConnectionMaxLifetime)
	db.SetMaxIdleConns(dbMaxIdleConnection)
	db.SetMaxOpenConns(dbMaxOpenConnection)

	return db, nil
}

func NewEntClient(db *sql.DB) *ent.Client {
	drv := entsql.OpenDB(dbDialect, db)
	return ent.NewClient(ent.Driver(drv))
}

func NewCacheClient(ctx context.Context, cfg *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})

	if err := client.Ping(ctx).Err(); err != nil {
		xutil.CloseConnection(client)
		return nil, xerr.FailedToReachHost.Wrap(err)
	}

	return client, nil
}
