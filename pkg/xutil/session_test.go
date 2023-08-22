package xutil

import (
	"context"
	"testing"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
	"github.com/42milez/go-oidc-server/pkg/xerr"
)

const userULID = "01H3M514Q0KGDS7PCKE9VVEMT4"

func TestSession_GetUserID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	sess := &Session{}

	want := typedef.UserID(userULID)

	ctx = context.WithValue(ctx, IDKey{}, want)

	got, ok := sess.GetUserID(ctx)

	if !ok {
		t.Fatalf("%s", xerr.FailedToReadContextValue)
	}

	if want != got {
		t.Errorf("want = %+v; got = %+v", want, got)
	}
}
