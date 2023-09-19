package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/ent/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatal(err)
	}

	workDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	migrationDir := filepath.Join(workDir, "app/ent/migrations")

	if err = os.MkdirAll(migrationDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir(migrationDir)

	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %+v", err)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.MySQL),           // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
	}

	if len(os.Args) != 2 {
		log.Fatal("migration name is required. Use: 'go run -mod=mod app/ent/cmd/migrate/main.go <MIGRATION_NAME>'")
	}

	ctx := context.Background()

	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	dbName := "atlas"
	dbAdmin := dbName
	dbPassword := dbName
	url := fmt.Sprintf("mysql://%s:%s@%s:%d/%s", dbAdmin, dbPassword, cfg.DBHost, cfg.DBPort, dbName)

	if err = migrate.NamedDiff(ctx, url, os.Args[1], opts...); err != nil {
		log.Fatalf("failed generating migration file: %+v", err)
	}
}
