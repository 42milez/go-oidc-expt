package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

const nMaxRetry = 5
const initialWaitTime = 2

func NewDatabase(ctx context.Context, cfg *config.Config) (*Database, error) {
	dataSrc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", cfg.DBAdmin, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open(dialect.MySQL, dataSrc)

	if err != nil {
		xutil.CloseConnection(db)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err = pingDB(ctx, db); err != nil {
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

func pingDB(ctx context.Context, db *sql.DB) error {
	retries := 0
	for {
		if err := db.PingContext(ctx); err != nil {
			if retries > nMaxRetry {
				return err
			}
			waitTime := (initialWaitTime << retries) + rand.Intn(1000)/1000
			time.Sleep(time.Duration(waitTime) * time.Second)
			retries++
			continue
		}
		break
	}
	return nil
}

func NewCache(ctx context.Context, cfg *config.Config) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})

	if err := pingCache(ctx, client); err != nil {
		xutil.CloseConnection(client)
		return nil, err
	}

	return &Cache{
		Client: client,
	}, nil
}

func pingCache(ctx context.Context, client *redis.Client) error {
	retries := 0
	for {
		if err := client.Ping(ctx).Err(); err != nil {
			if retries > nMaxRetry {
				return err
			}
			retries++
			waitTime := (initialWaitTime << retries) + rand.Intn(1000)/1000
			time.Sleep(time.Duration(waitTime) * time.Second)
			continue
		}
		break
	}
	return nil
}
