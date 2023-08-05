package repository

//import (
//	"context"
//	"errors"
//	"fmt"
//	"os"
//	"strconv"
//	"testing"
//
//	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
//
//	"github.com/42milez/go-oidc-server/pkg/xerr"
//
//	"github.com/42milez/go-oidc-server/app/idp/config"
//	"github.com/42milez/go-oidc-server/pkg/xtestutil"
//)
//
//const userULID = "01H3M514Q0KGDS7PCKE9VVEMT4"
//
//func TestNewEpStore(t *testing.T) {
//	t.Parallel()
//
//	if err := os.Setenv("REDIS_DB", strconv.Itoa(xtestutil.TestRedisDB)); err != nil {
//		t.Error(err)
//	}
//
//	cfg, err := config.New()
//
//	if err != nil {
//		t.Error(err)
//	}
//
//	ctx := context.Background()
//
//	sess, err := NewRedisClient(ctx, cfg)
//
//	if err != nil {
//		t.Fatalf("%s: %+v", xerr.FailedToInitialize, err)
//	}
//
//	if err = sess.Close(); err != nil {
//		t.Errorf("%s: %+v", xerr.FailedToCloseConnection, err)
//	}
//}
//
//func TestEpStore_SaveID(t *testing.T) {
//	t.Parallel()
//
//	client := xtestutil.OpenRedis(t)
//	repo := EpStore{
//		client: client,
//	}
//	ctx := context.Background()
//	key := "TestEpStore_SaveID"
//
//	t.Cleanup(func() {
//		client.Del(ctx, key)
//	})
//
//	id := typedef.UserID(userULID)
//
//	if err := repo.saveID(ctx, key, id); err != nil {
//		t.Error(err)
//	}
//}
//
//func TestEpStore_Load(t *testing.T) {
//	t.Parallel()
//
//	client := xtestutil.OpenRedis(t)
//	repo := EpStore{
//		client: client,
//	}
//
//	t.Run("OK", func(t *testing.T) {
//		t.Parallel()
//
//		ctx := context.Background()
//		key := "TestEpStore_Load_OK"
//		id := typedef.UserID(userULID)
//
//		if err := client.Set(ctx, key, id, sessionTTL).Err(); err != nil {
//			t.Fatal(err)
//		}
//
//		t.Cleanup(func() {
//			client.Del(ctx, key)
//		})
//
//		got, err := repo.load(ctx, key)
//
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		if got != id {
//			t.Errorf("want = %s; got = %s", id, got)
//		}
//	})
//
//	t.Run("NotFound", func(t *testing.T) {
//		t.Parallel()
//
//		ctx := context.Background()
//		key := "TestEpStore_Load_NotFound"
//
//		_, err := repo.load(ctx, key)
//
//		if err == nil || !errors.Is(err, ErrFailedToLoadItem) {
//			t.Errorf("want = %s; got = %+v", fmt.Sprintf("%s ( %s ): redis: nil", ErrFailedToLoadItem, key), err)
//		}
//	})
//}
//
//func TestEpStore_Delete(t *testing.T) {
//	t.Parallel()
//
//	repo := EpStore{
//		client: xtestutil.OpenRedis(t),
//	}
//	ctx := context.Background()
//	key := "TestEpStore_Delete"
//
//	id := typedef.UserID(userULID)
//
//	if err := repo.saveID(ctx, key, id); err != nil {
//		t.Fatal(err)
//	}
//
//	if err := repo.delete(ctx, key); err != nil {
//		t.Error(err)
//	}
//}
//
//func TestEpStore_SetID(t *testing.T) {
//	t.Parallel()
//
//	if err := os.Setenv("REDIS_DB", strconv.Itoa(xtestutil.TestRedisDB)); err != nil {
//		t.Error(err)
//	}
//
//	cfg, err := config.New()
//
//	if err != nil {
//		t.Error(err)
//	}
//
//	ctx := context.Background()
//	sess, err := NewRedisClient(ctx, cfg)
//
//	if err != nil {
//		t.Fatalf("%s: %+v", xerr.FailedToInitialize, err)
//	}
//
//	t.Cleanup(func() {
//		if err = sess.Close(); err != nil {
//			t.Errorf("%s: %+v", xerr.FailedToCloseConnection, err)
//		}
//	})
//
//	want := typedef.UserID(userULID)
//	ctx = sess.setID(ctx, want)
//
//	got, ok := ctx.Value(IDKey{}).(typedef.UserID)
//
//	if !ok {
//		t.Fatalf("%s", xerr.FailedToReadContextValue)
//	}
//
//	if want != got {
//		t.Errorf("want = %+v; got = %+v", want, got)
//	}
//}
//
//func TestEpStore_GetID(t *testing.T) {
//	t.Parallel()
//
//	if err := os.Setenv("REDIS_DB", strconv.Itoa(xtestutil.TestRedisDB)); err != nil {
//		t.Error(err)
//	}
//
//	cfg, err := config.New()
//
//	if err != nil {
//		t.Error(err)
//	}
//
//	ctx := context.Background()
//	sess, err := NewRedisClient(ctx, cfg)
//
//	if err != nil {
//		t.Fatalf("%s: %+v", xerr.FailedToInitialize, err)
//	}
//
//	t.Cleanup(func() {
//		if err = sess.Close(); err != nil {
//			t.Errorf("%s: %+v", xerr.FailedToCloseConnection, err)
//		}
//	})
//
//	want := typedef.UserID(userULID)
//	ctx = context.WithValue(ctx, IDKey{}, want)
//
//	got, ok := sess.getID(ctx)
//
//	if !ok {
//		t.Fatalf("%s", xerr.FailedToReadContextValue)
//	}
//
//	if want != got {
//		t.Errorf("want = %+v; got = %+v", want, got)
//	}
//}
