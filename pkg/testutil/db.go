package testutil

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"

	entsql "entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB(t *testing.T) (*ent.Client, *sql.DB) {
	dbUser := "test"
	dbPassword := "test"
	dbHost := "127.0.0.1"
	dbPort := 13306
	dbName := "test"
	dataSrc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", dbUser, dbPassword, dbHost, dbPort, dbName)

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
		t.Errorf("failed to close database connection: %+v", err)
	}
}
