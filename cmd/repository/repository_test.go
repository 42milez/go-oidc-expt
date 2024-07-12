package repository

import (
	"testing"

	"github.com/42milez/go-oidc-server/cmd/option"

	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/42milez/go-oidc-server/pkg/xid"
	"github.com/42milez/go-oidc-server/pkg/xtestutil"
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
	opt := &option.Option{
		DB:    xtestutil.NewDatabase(t, nil),
		IDGen: idGen,
	}
	if user := NewUser(opt); user == nil {
		t.Fatal(xerr.FailedToInitialize)
	}
}
