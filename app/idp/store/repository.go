package store

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/42milez/go-oidc-server/pkg/util"
	"time"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/rs/zerolog/log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

func New(ctx context.Context, cfg *config.Config) (*ent.Client, *sql.DB, func(), error) {
	dataSrc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", cfg.DBAdmin, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open(dialect.MySQL, dataSrc)
	if err != nil {
		return nil, nil, func() {}, err
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, nil, func() {
			if err := db.Close(); err != nil {
				log.Error().Err(err).Msg("failed to close database connection")
			}
		}, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	drv := entsql.OpenDB("mysql", db)
	entClient := ent.NewClient(ent.Driver(drv))

	return entClient, db, func() {
		if err := entClient.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close database connection")
		}
	}, nil
}

type Repository struct {
	Clocker util.Clocker
}
