package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

func NewDatabase(ctx context.Context, cfg *config.Config) (*Database, error) {
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

	drv := entsql.OpenDB(dbDialect, db)
	client := ent.NewClient(ent.Driver(drv))

	return &Database{
		Client: client,
		db:     db,
	}, nil
}

func NewCache(ctx context.Context, cfg *config.Config) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})

	if err := client.Ping(ctx).Err(); err != nil {
		xutil.CloseConnection(client)
		return nil, xerr.FailedToReachHost.Wrap(err)
	}

	return &Cache{
		Client: client,
	}, nil
}
