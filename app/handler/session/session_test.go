package session

import (
	"context"
	"testing"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

const userULID typedef.UserID = 475924034190589956

func TestGetUserID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	want := userULID

	ctx = context.WithValue(ctx, UserIDKey{}, want)
	got, ok := GetUserID(ctx)

	if !ok {
		t.Fatalf("%s", xerr.FailedToReadContextValue)
	}

	if want != got {
		t.Errorf("want = %+v; got = %+v", want, got)
	}
}
