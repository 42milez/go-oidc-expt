package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/entity"
	"github.com/google/go-cmp/cmp"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/redis/go-redis/v9"
)

func TestSession_SaveID(t *testing.T) {
	t.Parallel()

	cache := xtestutil.NewCache(t)
	repo := Session{
		cache: cache,
	}
	ctx := context.Background()
	sid := "TestSession_SaveID"

	t.Cleanup(func() {
		cache.Client.Del(ctx, sid)
	})

	sess := &entity.Session{
		UserID: typedef.UserID(475924035230777348),
	}

	ok, err := repo.Create(ctx, typedef.SessionID(sid), sess)

	if err != nil {
		t.Error(err)
	}

	if !ok {
		t.Error(xerr.SessionIDAlreadyExists)
	}
}

func TestSession_LoadID(t *testing.T) {
	t.Parallel()

	cache := xtestutil.NewCache(t)
	repo := Session{
		cache: cache,
	}

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		sid := "TestSession_LoadID_OK"
		want := &entity.Session{
			UserID: typedef.UserID(475924035230777348),
		}

		if err := cache.Client.Set(ctx, sid, want, config.SessionTTL).Err(); err != nil {
			t.Fatal(err)
		}

		t.Cleanup(func() {
			cache.Client.Del(ctx, sid)
		})

		got, err := repo.Read(ctx, typedef.SessionID(sid))

		if err != nil {
			t.Fatal(err)
		}

		if d := cmp.Diff(want, got); !xutil.IsEmpty(d) {
			t.Errorf("item not matched (-got +want)\n%s", d)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		sid := "TestSession_LoadID_NotFound"

		_, err := repo.Read(ctx, typedef.SessionID(sid))

		if err == nil || !errors.Is(err, redis.Nil) {
			t.Errorf("want = %+v; got = %+v", redis.Nil, err)
		}
	})
}

func TestSession_Delete(t *testing.T) {
	t.Parallel()

	cache := xtestutil.NewCache(t)
	repo := Session{
		cache: cache,
	}
	ctx := context.Background()
	sid := "TestSession_Delete"
	sess := &entity.Session{
		UserID: typedef.UserID(475924035230777348),
	}

	ok, err := cache.Client.SetNX(ctx, sid, sess, config.SessionTTL).Result()

	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Error(xerr.SessionIDAlreadyExists)
	}

	if err = repo.Delete(ctx, typedef.SessionID(sid)); err != nil {
		t.Error(err)
	}
}
