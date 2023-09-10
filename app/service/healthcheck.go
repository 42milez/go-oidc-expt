package service

import (
	"context"
)

type CheckHealth struct {
	repo HealthChecker
}

func (ch *CheckHealth) CheckCacheStatus(ctx context.Context) error {
	return ch.repo.PingCache(ctx)
}

func (ch *CheckHealth) CheckDBStatus(ctx context.Context) error {
	return ch.repo.PingDatabase(ctx)
}
