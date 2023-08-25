package xtestutil

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/42milez/go-oidc-server/app/ent/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

const TestDBHost = "127.0.0.1"
const TestDBPort = 13306
const TestDBUser = "idp_test"
const TestDBPassword = "idp_test"
const TestDBName = "idp_test"

func OpenDB(t *testing.T) (*ent.Client, *sql.DB) {
	dataSrc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", TestDBUser, TestDBPassword, TestDBHost, TestDBPort,
		TestDBName)

	db, err := sql.Open(dialect.MySQL, dataSrc)
	if err != nil {
		t.Fatalf("failed to open database: %+v", err)
	}

	drv := entsql.OpenDB(dialect.MySQL, db)
	entClient := ent.NewClient(ent.Driver(drv))

	t.Cleanup(func() {
		closeDB(t, entClient)
	})

	return entClient, db
}

func closeDB(t *testing.T, client *ent.Client) {
	if err := client.Close(); err != nil {
		t.Fatalf("failed to close connection: %+v", err)
	}
}

const TestRedisHost = "127.0.0.1"
const TestRedisPort = 16379
const TestRedisPassword = ""
const TestRedisDB = 1

func OpenRedis(t *testing.T) *redis.Client {
	t.Helper()

	opt := redis.Options{
		Addr:     fmt.Sprintf("%s:%d", TestRedisHost, TestRedisPort),
		Password: TestRedisPassword,
		DB:       TestRedisDB,
	}
	client := redis.NewClient(&opt)

	if err := client.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("failed to establish connection: %+v", err)
	}

	t.Cleanup(func() {
		closeRedis(t, client)
	})

	return client
}

func closeRedis(t *testing.T, client *redis.Client) {
	if err := client.Close(); err != nil {
		t.Fatalf("failed to close connection: %+v", err)
	}
}
