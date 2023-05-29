package service

import (
	"context"
	"database/sql"
	"errors"
)

type CheckHealth struct {
	DB *sql.DB
}

func (p *CheckHealth) PingCache(ctx context.Context) error {
	return errors.New("not implemented")
}

func (p *CheckHealth) PingDB(ctx context.Context) error {
	if err := p.DB.PingContext(ctx); err != nil {
		return err
	}
	return nil
}
