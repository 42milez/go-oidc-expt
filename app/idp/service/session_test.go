package service

import (
	"context"
	"testing"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
	"github.com/42milez/go-oidc-server/pkg/xerr"
)

const userULID = "01H3M514Q0KGDS7PCKE9VVEMT4"

func TestSession_SetID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	sess := &Session{}

	want := typedef.UserID(userULID)

	ctx = sess.SetID(ctx, want)

	got, ok := ctx.Value(IDKey{}).(typedef.UserID)

	if !ok {
		t.Fatalf("%s", xerr.FailedToReadContextValue)
	}

	if want != got {
		t.Errorf("want = %+v; got = %+v", want, got)
	}
}

func TestSession_GetID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	sess := &Session{}

	want := typedef.UserID(userULID)

	ctx = context.WithValue(ctx, IDKey{}, want)

	got, ok := sess.GetID(ctx)

	if !ok {
		t.Fatalf("%s", xerr.FailedToReadContextValue)
	}

	if want != got {
		t.Errorf("want = %+v; got = %+v", want, got)
	}
}
