package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"

	"github.com/42milez/go-oidc-server/app/idp/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	migrationDir := filepath.Join(workDir, "app/idp/ent/migrations")
	if err := os.MkdirAll(migrationDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir(migrationDir)
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.MySQL),           // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}
	if len(os.Args) != 2 {
		log.Fatal("migration name is required. Use: 'go run -mod=mod app/idp/ent/cmd/migrate/main.go <name>'")
	}

	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	dbPort := 13306
	dbAdmin := "atlas"
	dbPassword := "atlas"
	dbName := "atlas"
	url := fmt.Sprintf("mysql://%s:%s@%s:%d/%s", dbAdmin, dbPassword, cfg.DBHost, dbPort, dbName)
	if err = migrate.NamedDiff(ctx, url, os.Args[1], opts...); err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
