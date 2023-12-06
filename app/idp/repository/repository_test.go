package repository

import (
	"testing"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
)

func TestNewCheckHealth(t *testing.T) {
	t.Parallel()
	if ch := NewCheckHealth(xtestutil.NewDatabase(t, nil), xtestutil.NewCache(t)); ch == nil {
		t.Fatal(xerr.FailedToInitialize)
	}
}

func TestNewSession(t *testing.T) {
	t.Parallel()
	if sess := NewCache(xtestutil.NewCache(t)); sess == nil {
		t.Fatal(xerr.FailedToInitialize)
	}
}

func TestNewUser(t *testing.T) {
	t.Parallel()
	idGen, err := xid.GetUniqueIDGenerator()
	if err != nil {
		t.Fatal(err)
	}
	if user := NewUser(xtestutil.NewDatabase(t, nil), idGen); user == nil {
		t.Fatal(xerr.FailedToInitialize)
	}
}
