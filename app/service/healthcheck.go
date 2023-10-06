package service

import (
	"context"
)

func NewCheckHealth(repo HealthChecker) *CheckHealth {
	return &CheckHealth{
		repo: repo,
	}
}

type CheckHealth struct {
	repo HealthChecker
}

func (ch *CheckHealth) CheckCacheStatus(ctx context.Context) error {
	return ch.repo.PingCache(ctx)
}

func (ch *CheckHealth) CheckDBStatus(ctx context.Context) error {
	return ch.repo.PingDatabase(ctx)
}
