package xtestutil

import (
	"os"
	"testing"
	"time"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

const DummyError DummyErr = "DUMMY ERROR"

type DummyErr string

func (v DummyErr) Error() string {
	return string(v)
}

type FixedClocker struct{}

func (v FixedClocker) Now() time.Time {
	return time.Date(2000, 12, 31, 23, 59, 59, 0, time.UTC)
}

type FixedTomorrowClocker struct{}

func (v FixedTomorrowClocker) Now() time.Time {
	return FixedClocker{}.Now().Add(24 * time.Hour)
}

func LoadFile(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("%s: %v", xerr.FailedToReadFile, err)
	}
	return data
}
