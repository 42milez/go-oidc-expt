package store

import (
	"context"
	"errors"
	"os"
	"strconv"
	"testing"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/ent/alias"
	"github.com/42milez/go-oidc-server/pkg/testutil"
)

func TestNewSession(t *testing.T) {
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
		t.Fatalf("failed to create session: %v", err)
	}

	if err := repo.Close(); err != nil {
		t.Errorf("failed to close connection: %v", err)
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
		t.Errorf("failed to save id ( key = %s, id = %d ): want = ( no error ); got = %v", key, id, err)
	}
}

func TestSession_LoadID(t *testing.T) {
	t.Parallel()

	client := testutil.OpenRedis(t)
	repo := Session[alias.AdminID]{
		client: client,
	}

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		key := "TestSession_LoadID_OK"
		id := alias.AdminID(123)

		client.Set(ctx, key, uint64(id), sessionTTL)

		t.Cleanup(func() {
			client.Del(ctx, key)
		})

		got, err := repo.LoadID(ctx, key)
		if err != nil {
			t.Fatalf("faild to load id ( key = %s ): want = ( no error ); got = %v", key, err)
		}
		if got != id {
			t.Errorf("id not matched: want = %d; got = %d", id, got)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		key := "TestSession_LoadID_NotFound"

		_, err := repo.LoadID(ctx, key)
		if err == nil || !errors.Is(err, ErrNotFound) {
			t.Errorf("unexpected behavior on loading id: want = %v; got = %v", ErrNotFound, err)
		}
	})
}

func TestSession_DeleteID(t *testing.T) {
	t.Parallel()

	repo := Session[alias.AdminID]{
		client: testutil.OpenRedis(t),
	}
	ctx := context.Background()
	key := "TestSession_DeleteID"

	id := alias.AdminID(123)

	if err := repo.SaveID(ctx, key, id); err != nil {
		t.Fatalf("failed to save id ( key = %s, id = %d ): want = ( no error ); got = %v", key, id, err)
	}

	if err := repo.DeleteID(ctx, key); err != nil {
		t.Errorf("failed to delete id ( key = %s ): want = ( no error ); got = %v", key, err)
	}
}
