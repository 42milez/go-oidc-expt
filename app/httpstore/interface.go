package httpstore

import (
	"context"
	"time"
)

//go:generate go run -mod=mod go.uber.org/mock/mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

type SessionReader interface {
	Read(ctx context.Context, key string) (string, error)
}

type SessionWriter interface {
	Write(ctx context.Context, key string, value any, ttl time.Duration) (bool, error)
}

type IdGenerator interface {
	NextID() (uint64, error)
}
