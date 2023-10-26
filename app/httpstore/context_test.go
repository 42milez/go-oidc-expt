package httpstore

import (
	"context"
	"testing"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

func TestReadContext_Read(t *testing.T) {
	t.Parallel()

	type CtxKey struct{}

	want := "TestReadContext_Read"
	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxKey{}, want)
	rc := &ReadContext{}

	got, ok := rc.Read(ctx, CtxKey{}).(string)
	if !ok {
		t.Fatal(xerr.TypeAssertionFailed)
	}
	if want != got {
		t.Errorf("want = %s; got =%s", want, got)
	}
}

func TestWriteContext_Write(t *testing.T) {
	t.Parallel()

	type CtxKey struct{}

	ctx := context.Background()
	want := "TestWriteContext_Write"
	wc := &WriteContext{}
	ctx = wc.Write(ctx, CtxKey{}, want)

	got, ok := ctx.Value(CtxKey{}).(string)
	if !ok {
		t.Fatal(xerr.TypeAssertionFailed)
	}
	if want != got {
		t.Errorf("want = %s; got =%s", want, got)
	}
}
