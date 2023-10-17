package httpstore

import (
	"context"
)

type ReadContext struct{}

func (rc *ReadContext) Read(ctx context.Context, key any) any {
	return ctx.Value(key)
}

type WriteContext struct{}

func (wc *WriteContext) Write(ctx context.Context, key any, val any) context.Context {
	return context.WithValue(ctx, key, val)
}
