package repository

import (
	"context"
	"encoding/json"
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

func TestSession_Create(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	repo := Session{
		cache: xtestutil.NewCache(t),
	}
	sid := typedef.SessionID("TestSession_Create")

	t.Cleanup(func() {
		repo.cache.Client.Del(ctx, string(sid))
	})

	sess := &entity.Session{
		UserID: typedef.UserID(475924035230777348),
	}

	ok, err := repo.Create(ctx, sid, sess)

	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Error(xerr.SessionIDAlreadyExists)
	}
}

func TestSession_Read(t *testing.T) {
	t.Parallel()

	repo := Session{
		cache: xtestutil.NewCache(t),
	}

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		sid := "TestSession_Read_OK"
		want := &entity.Session{
			UserID: typedef.UserID(475924035230777348),
		}

		if err := repo.cache.Client.SetNX(ctx, sid, want, config.SessionTTL).Err(); err != nil {
			t.Fatal(err)
		}

		t.Cleanup(func() {
			repo.cache.Client.Del(ctx, sid)
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
		sid := "TestSession_Read_NotFound"

		_, err := repo.Read(ctx, typedef.SessionID(sid))

		// TODO: Replace 'redis.Nil' with other error.
		if err == nil || !errors.Is(err, redis.Nil) {
			t.Errorf("want = %+v; got = %+v", redis.Nil, err)
		}
	})
}

func TestSession_Update(t *testing.T) {
	t.Parallel()

	repo := Session{
		cache: xtestutil.NewCache(t),
	}

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		sid := typedef.SessionID("TestSession_Update_OK")
		sess := &entity.Session{
			UserID: typedef.UserID(475924035230777348),
		}

		if err := repo.cache.Client.SetNX(ctx, string(sid), sess, config.SessionTTL).Err(); err != nil {
			t.Fatal(err)
		}

		t.Cleanup(func() {
			repo.cache.Client.Del(ctx, string(sid))
		})

		want := &entity.Session{
			UserID: typedef.UserID(475924035230777348),
		}

		if _, err := repo.Update(ctx, sid, want); err != nil {
			t.Error(err)
		}

		v, err := repo.cache.Client.Get(ctx, string(sid)).Result()

		if err != nil {
			t.Fatal(err)
		}

		got := &entity.Session{}

		if err = json.Unmarshal([]byte(v), got); err != nil {
			t.Fatal(err)
		}

		if d := cmp.Diff(want, got); !xutil.IsEmpty(d) {
			t.Errorf("item not matched (-got +want)\n%s", d)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		sid := "TestSession_Read_NotFound"

		_, err := repo.Read(ctx, typedef.SessionID(sid))

		// TODO: Replace 'redis.Nil' with other error.
		if err == nil || !errors.Is(err, redis.Nil) {
			t.Errorf("want = %+v; got = %+v", redis.Nil, err)
		}
	})
}

func TestSession_Delete(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	repo := Session{
		cache: xtestutil.NewCache(t),
	}
	sid := "TestSession_Delete"
	sess := &entity.Session{
		UserID: typedef.UserID(475924035230777348),
	}

	if _, err := repo.cache.Client.SetNX(ctx, sid, sess, config.SessionTTL).Result(); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		repo.cache.Client.Del(ctx, sid)
	})

	if err := repo.Delete(ctx, typedef.SessionID(sid)); err != nil {
		t.Error(err)
	}
}
