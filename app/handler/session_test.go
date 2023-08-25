package handler

import (
	"context"
	"testing"

	"github.com/42milez/go-oidc-server/app/ent/typedef"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

const userULID = "01H3M514Q0KGDS7PCKE9VVEMT4"

func TestGetUserID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	want := typedef.UserID(userULID)

	ctx = context.WithValue(ctx, UserIDKey{}, want)
	got, ok := GetUserID(ctx)

	if !ok {
		t.Fatalf("%s", xerr.FailedToReadContextValue)
	}

	if want != got {
		t.Errorf("want = %+v; got = %+v", want, got)
	}
}
