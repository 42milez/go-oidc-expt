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
		t.Error(err)
	}

	if err := repo.Close(); err != nil {
		t.Error(err)
	}
}

func TestAdminSession_Save(t *testing.T) {
	t.Parallel()

	client := testutil.OpenRedis(t)
	repo := AdminSession{
		client: client,
	}
	ctx := context.Background()
	key := "TestAdminSession_Save"

	t.Cleanup(func() {
		if err := repo.DeleteID(ctx, key); err != nil {
			t.Error(err)
		}
	})

	id := alias.AdminID(123)

	if err := repo.SaveID(ctx, key, id); err != nil {
		t.Errorf("want = ( no error ); got = %v", err)
	}
}

func TestAdminSession_LoadID(t *testing.T) {
	t.Parallel()

	client := testutil.OpenRedis(t)
	repo := AdminSession{
		client: client,
	}

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		key := "TestAdminSession_LoadID_OK"
		id := alias.AdminID(123)

		client.Set(ctx, key, uint64(id), sessionTTL)

		t.Cleanup(func() {
			client.Del(ctx, key)
		})

		got, err := repo.LoadID(ctx, key)
		if err != nil {
			t.Fatalf("want = ( no error ); got = %+v", err)
		}
		if got != id {
			t.Errorf("want = %d; got = %d", id, got)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		key := "TestAdminSession_LoadID_NotFound"

		got, err := repo.LoadID(ctx, key)
		if err == nil || !errors.Is(err, ErrNotFound) {
			t.Errorf("want = %v; got = %v ( value: %d )", ErrNotFound, err, got)
		}
	})
}
