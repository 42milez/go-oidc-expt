package httpstore

import (
	"context"
)

type Context struct{}

func (c *Context) Read(ctx context.Context, key any) any {
	return ctx.Value(key)
}

func (c *Context) Write(ctx context.Context, key any, val any) context.Context {
	return context.WithValue(ctx, key, val)
}
