package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/redis/go-redis/v9"
)

func TestSession_Write(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	repo := Session{
		cache: xtestutil.NewCache(t),
	}
	sid := "484481116225601901"
	uid := typedef.UserID(475924035230777348)

	t.Cleanup(func() {
		repo.cache.Client.Del(ctx, sid)
	})

	ok, err := repo.Write(ctx, sid, uid, config.SessionTTL)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Error(xerr.FailedToWriteSession)
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
		sid := "484481116225667437"
		wantUserId := "475924035230777348"

		if err := repo.cache.Client.SetNX(ctx, sid, wantUserId, config.SessionTTL).Err(); err != nil {
			t.Fatal(err)
		}

		t.Cleanup(func() {
			repo.cache.Client.Del(ctx, sid)
		})

		gotUserId, err := repo.Read(ctx, sid)

		if err != nil {
			t.Fatal(err)
		}

		if wantUserId != gotUserId {
			t.Errorf("want = %s; got = %s", wantUserId, gotUserId)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		sid := "TestSession_Read_NotFound"

		_, err := repo.Read(ctx, sid)

		// TODO: Replace 'redis.Nil' with other error.
		if err == nil || !errors.Is(err, redis.Nil) {
			t.Errorf("want = %+v; got = %+v", redis.Nil, err)
		}
	})
}
