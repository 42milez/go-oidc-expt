package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/rs/zerolog/log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDialect             = "mysql"
	maxOpenConnection     = 100
	maxIdleConnection     = 10
	connectionMaxLifetime = time.Hour
)

func NewDB(ctx context.Context, cfg *config.Config) (*ent.Client, *sql.DB, func(), error) {
	dataSrc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", cfg.DBAdmin, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open(dialect.MySQL, dataSrc)

	if err != nil {
		return nil, nil, nil, xerr.WrapErr(xerr.FailToEstablishConnection, err)
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		if err = db.Close(); err != nil {
			log.Error().Err(err).Msg(xerr.FailedToCloseConnection.Error())
		}
		return nil, nil, nil, err
	}

	db.SetConnMaxLifetime(connectionMaxLifetime)
	db.SetMaxIdleConns(maxIdleConnection)
	db.SetMaxOpenConns(maxOpenConnection)

	drv := entsql.OpenDB(dbDialect, db)
	entClient := ent.NewClient(ent.Driver(drv))

	return entClient, db, func() {
		if err = entClient.Close(); err != nil {
			log.Error().Err(err).Msg(xerr.FailedToCloseConnection.Error())
		}
	}, nil
}
