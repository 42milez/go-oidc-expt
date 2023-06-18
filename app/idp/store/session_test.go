package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/42milez/go-oidc-server/app/idp/auth"
	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/ent/alias"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/pkg/testutil"
	"github.com/42milez/go-oidc-server/pkg/testutil/fixture"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strconv"
	"testing"
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

func TestSession_ExtractToken(t *testing.T) {
	t.Parallel()

	j, err := auth.NewJWT(testutil.FixedClocker{})

	if err != nil {
		t.Fatalf("failed to create jwt: %v", err)
	}

	admin := fixture.Admin(&ent.Admin{})
	signed, err := j.GenerateAdminAccessToken(admin)
	want, err := j.Parse(signed)

	req := httptest.NewRequest(http.MethodGet, "https://github.com/42milez", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", signed))

	cfg, err := config.New()

	if err != nil {
		t.Error(err)
	}

	ctx := context.Background()
	sess, err := NewAdminSession(ctx, cfg)

	if err != nil {
		t.Fatalf("failed to create session: %v", err)
	}

	if err := sess.SaveID(ctx, want.JwtID(), admin.ID); err != nil {
		t.Fatalf("failed to save id: %v", err)
	}

	t.Cleanup(func() {
		if err := sess.DeleteID(ctx, want.JwtID()); err != nil {
			t.Errorf("failed to delete id: %v", err)
		}
		if err := sess.Close(); err != nil {
			t.Errorf("failed to close connection: %v", err)
		}
	})

	if err != nil {
		t.Fatalf("failed to parse token: %v", err)
	}

	got, err := sess.ExtractToken(ctx, req)

	if err != nil {
		t.Fatalf("failed to extract token: %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("token not matched: want = %v; got = %v", want, got)
	}
}
