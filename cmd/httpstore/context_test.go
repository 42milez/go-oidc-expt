package httpstore

import (
	"context"
	"testing"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

func TestContext_Read(t *testing.T) {
	t.Parallel()

	type CtxKey struct{}

	want := "TestReadContext_Read"
	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKey{}, want)
	c := &Context{}

	got, ok := c.Read(ctx, CtxKey{}).(string)
	if !ok {
		t.Fatal(xerr.TypeAssertionFailed)
	}
	if want != got {
		t.Errorf("want = %s; got =%s", want, got)
	}
}

func TestContext_Write(t *testing.T) {
	t.Parallel()

	type CtxKey struct{}

	ctx := context.Background()
	want := "TestWriteContext_Write"
	c := &Context{}
	ctx = c.Write(ctx, CtxKey{}, want)

	got, ok := ctx.Value(CtxKey{}).(string)
	if !ok {
		t.Fatal(xerr.TypeAssertionFailed)
	}
	if want != got {
		t.Errorf("want = %s; got =%s", want, got)
	}
}
