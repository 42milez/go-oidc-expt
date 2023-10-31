package httpstore

import (
	"context"
	"time"
)

//go:generate go run -mod=mod go.uber.org/mock/mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

type SessionBasicReader interface {
	Read(ctx context.Context, key string) (string, error)
}

type SessionHashReader interface {
	ReadHash(ctx context.Context, key string, field string) (string, error)
}

type SessionReader interface {
	SessionBasicReader
	SessionHashReader
}

type SessionBasicWriter interface {
	Write(ctx context.Context, key string, value any, ttl time.Duration) (bool, error)
}

type SessionHashWriter interface {
	WriteHash(ctx context.Context, key string, values map[string]string, ttl time.Duration) (bool, error)
}

type SessionWriter interface {
	SessionBasicWriter
	SessionHashWriter
}

type IdGenerator interface {
	NextID() (uint64, error)
}
