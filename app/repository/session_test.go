package repository

import (
	"context"
	"errors"
	"os"
	"strconv"
	"testing"

	"github.com/42milez/go-oidc-server/pkg/xtestutil"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/ent/typedef"

	"github.com/redis/go-redis/v9"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

const userULID = "01H3M514Q0KGDS7PCKE9VVEMT4"

func TestNewSession(t *testing.T) {
	t.Parallel()

	if err := os.Setenv("REDIS_DB", strconv.Itoa(xtestutil.TestRedisDB)); err != nil {
		t.Error(err)
	}

	cfg, err := config.New()

	if err != nil {
		t.Error(err)
	}

	ctx := context.Background()

	sess, err := NewCacheClient(ctx, cfg)

	if err != nil {
		t.Fatalf("%s: %+v", xerr.FailedToInitialize, err)
	}

	if err = sess.Close(); err != nil {
		t.Errorf("%s: %+v", xerr.FailedToCloseConnection, err)
	}
}

func TestSession_SaveID(t *testing.T) {
	t.Parallel()

	client := xtestutil.OpenRedis(t)
	repo := Session{
		Cache: client,
	}
	ctx := context.Background()
	key := "TestEpStore_SaveID"

	t.Cleanup(func() {
		client.Del(ctx, key)
	})

	id := typedef.UserID(userULID)

	if err := repo.SaveUserID(ctx, key, id); err != nil {
		t.Error(err)
	}
}

func TestSession_LoadID(t *testing.T) {
	t.Parallel()

	client := xtestutil.OpenRedis(t)
	repo := Session{
		Cache: client,
	}

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		key := "TestEpStore_Load_OK"
		id := typedef.UserID(userULID)

		if err := client.Set(ctx, key, id, sessionTTL).Err(); err != nil {
			t.Fatal(err)
		}

		t.Cleanup(func() {
			client.Del(ctx, key)
		})

		got, err := repo.LoadUserID(ctx, key)

		if err != nil {
			t.Fatal(err)
		}

		if got != id {
			t.Errorf("want = %s; got = %s", id, got)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		key := "TestEpStore_Load_NotFound"

		_, err := repo.LoadUserID(ctx, key)

		if err == nil || !errors.Is(err, redis.Nil) {
			t.Errorf("want = %+v; got = %+v", redis.Nil, err)
		}
	})
}

func TestSession_Delete(t *testing.T) {
	t.Parallel()

	client := xtestutil.OpenRedis(t)
	repo := Session{
		Cache: client,
	}
	ctx := context.Background()
	key := "TestEpStore_Delete"

	id := typedef.UserID(userULID)

	if err := repo.SaveUserID(ctx, key, id); err != nil {
		t.Fatal(err)
	}

	if err := repo.Delete(ctx, key); err != nil {
		t.Error(err)
	}
}
