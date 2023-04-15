package handler

import (
	"context"
)

type CheckHealthService interface {
	PingCache(ctx context.Context) error
	PingDB(ctx context.Context) error
}
