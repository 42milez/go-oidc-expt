package repository

import (
	"testing"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
)

func TestNewCheckHealth(t *testing.T) {
	t.Parallel()
	if ch := NewCheckHealth(xtestutil.NewDatabase(t), xtestutil.NewCache(t)); ch == nil {
		t.Error(xerr.FailedToInitialize)
	}
}

func TestNewSession(t *testing.T) {
	t.Parallel()
	if sess := NewSession(xtestutil.NewCache(t)); sess == nil {
		t.Error(xerr.FailedToInitialize)
	}
}

func TestNewUser(t *testing.T) {
	t.Parallel()
	idGen, err := xid.GetUniqueID()
	if err != nil {
		t.Fatal(err)
	}
	if user := NewUser(xtestutil.NewDatabase(t), idGen); user == nil {
		t.Error(xerr.FailedToInitialize)
	}
}
