package testutil

import (
	"os"
	"testing"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

func LoadFile(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("%s: %v", xerr.FailedToReadFile, err)
	}
	return data
}
