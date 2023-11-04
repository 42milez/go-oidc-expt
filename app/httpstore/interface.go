package httpstore

import (
	"context"
	"time"
)

//go:generate go run -mod=mod go.uber.org/mock/mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

type CacheReader interface {
	Read(ctx context.Context, key string) (string, error)
}

type CacheHashReader interface {
	ReadHash(ctx context.Context, key string, field string) (string, error)
}

type CacheHashAllReader interface {
	ReadHashAll(ctx context.Context, key string) (map[string]string, error)
}

type CacheWriter interface {
	Write(ctx context.Context, key string, value any, ttl time.Duration) (bool, error)
}

type CacheHashWriter interface {
	WriteHash(ctx context.Context, key string, values map[string]string, ttl time.Duration) (bool, error)
}

type CacheReadWriter interface {
	CacheReader
	CacheHashReader
	CacheHashAllReader
	CacheWriter
	CacheHashWriter
}
