package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/42milez/go-oidc-server/pkg/xerr"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/ent/alias"
	"github.com/42milez/go-oidc-server/pkg/testutil"
)

func TestNewAdminSession(t *testing.T) {
	t.Parallel()

	if err := os.Setenv("REDIS_DB", strconv.Itoa(testutil.TestRedisDB)); err != nil {
		t.Error(err)
	}

	cfg, err := config.New()

	if err != nil {
		t.Error(err)
	}

	ctx := context.Background()

	repo, err := NewAdminSession(ctx, cfg)

	if err != nil {
		t.Fatalf("%s: %v", xerr.FailedToInitialize, err)
	}

	if err = repo.Close(); err != nil {
		t.Errorf("%s: %v", xerr.FailedToCloseConnection, err)
	}
}

func TestSession_SaveID(t *testing.T) {
	t.Parallel()

	client := testutil.OpenRedis(t)
	repo := Session[alias.AdminID]{
		client: client,
	}
	ctx := context.Background()
	key := "TestSession_SaveID"

	t.Cleanup(func() {
		client.Del(ctx, key)
	})

	id := alias.AdminID(123)

	if err := repo.SaveID(ctx, key, id); err != nil {
		t.Error(err)
	}
}

func TestSession_Load(t *testing.T) {
	t.Parallel()

	client := testutil.OpenRedis(t)
	repo := Session[alias.AdminID]{
		client: client,
	}

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		key := "TestSession_Load_OK"
		id := alias.AdminID(123)

		client.Set(ctx, key, uint64(id), sessionTTL)

		t.Cleanup(func() {
			client.Del(ctx, key)
		})

		got, err := repo.Load(ctx, key)

		if err != nil {
			t.Fatal(err)
		}

		if got != id {
			t.Errorf("want = %d; got = %d", id, got)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		key := "TestSession_Load_NotFound"

		_, err := repo.Load(ctx, key)

		if err == nil || !errors.Is(err, ErrFailedToLoad) {
			t.Errorf("want = %s; got = %v", fmt.Sprintf("%s ( %s ): redis: nil", ErrFailedToLoad, key), err)
		}
	})
}

func TestSession_Delete(t *testing.T) {
	t.Parallel()

	repo := Session[alias.AdminID]{
		client: testutil.OpenRedis(t),
	}
	ctx := context.Background()
	key := "TestSession_Delete"

	id := alias.AdminID(123)

	if err := repo.SaveID(ctx, key, id); err != nil {
		t.Fatal(err)
	}

	if err := repo.Delete(ctx, key); err != nil {
		t.Error(err)
	}
}

func TestSession_SetID(t *testing.T) {
	t.Parallel()

	if err := os.Setenv("REDIS_DB", strconv.Itoa(testutil.TestRedisDB)); err != nil {
		t.Error(err)
	}

	cfg, err := config.New()

	if err != nil {
		t.Error(err)
	}

	ctx := context.Background()
	sess, err := NewAdminSession(ctx, cfg)

	if err != nil {
		t.Fatalf("%s: %v", xerr.FailedToInitialize, err)
	}

	want := alias.AdminID(123)
	ctx = sess.SetID(ctx, want)

	got, ok := ctx.Value(IDKey{}).(alias.AdminID)

	if !ok {
		t.Fatalf("%s", xerr.FailedToReadContextValue)
	}

	if want != got {
		t.Errorf("want = %v; got = %v", want, got)
	}
}

func TestSession_GetID(t *testing.T) {
	t.Parallel()

	if err := os.Setenv("REDIS_DB", strconv.Itoa(testutil.TestRedisDB)); err != nil {
		t.Error(err)
	}

	cfg, err := config.New()

	if err != nil {
		t.Error(err)
	}

	ctx := context.Background()
	sess, err := NewAdminSession(ctx, cfg)

	if err != nil {
		t.Fatalf("%s: %v", xerr.FailedToInitialize, err)
	}

	want := alias.AdminID(123)
	ctx = context.WithValue(ctx, IDKey{}, want)

	got, ok := sess.GetID(ctx)

	if !ok {
		t.Fatalf("%s", xerr.FailedToReadContextValue)
	}

	if want != got {
		t.Errorf("want = %v; got = %v", want, got)
	}
}
