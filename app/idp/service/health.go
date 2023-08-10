package service

import (
	"context"
)

type CheckHealth struct {
	Repo HealthChecker
}

func (p *CheckHealth) CheckCacheStatus(ctx context.Context) error {
	return p.Repo.PingCache(ctx)
}

func (p *CheckHealth) CheckDBStatus(ctx context.Context) error {
	return p.Repo.PingDB(ctx)
}
