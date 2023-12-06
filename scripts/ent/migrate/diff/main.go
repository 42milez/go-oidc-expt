package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/42milez/go-oidc-server/app/idp/config"

	"github.com/42milez/go-oidc-server/app/ent/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	workDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	migrationDir := filepath.Join(workDir, "app/ent/migrations")

	if err = os.MkdirAll(migrationDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	dir, err := atlas.NewLocalDir(migrationDir)

	if err != nil {
		log.Fatalf("failed to create atlas migration directory: %+v", err)
	}

	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithDialect(dialect.MySQL),
		schema.WithMigrationMode(schema.ModeReplay),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	if len(os.Args) != 2 {
		log.Fatal("migration name is required. Use: 'go run -mod=mod scripts/ent/migrate/diff/main.go <MIGRATION_NAME>'")
	}

	ctx := context.Background()
	cfg, err := config.New()

	if err != nil {
		log.Fatal(err)
	}

	cfg.DBAdmin = cfg.DBName
	cfg.DBPassword = cfg.DBName

	url := fmt.Sprintf("mysql://%s:%s@%s:%d/%s", cfg.DBAdmin, cfg.DBPassword, cfg.DB1Host, cfg.DB1Port, cfg.DBName)

	if err = migrate.NamedDiff(ctx, url, os.Args[1], opts...); err != nil {
		log.Fatalf("failed to generate migration file: %+v", err)
	}
}
