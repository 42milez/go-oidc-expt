package service

import (
	"context"
)

type CheckHealth struct {
	repo HealthChecker
}

func (p *CheckHealth) CheckCacheStatus(ctx context.Context) error {
	return p.repo.PingCache(ctx)
}

func (p *CheckHealth) CheckDBStatus(ctx context.Context) error {
	return p.repo.PingDB(ctx)
}
