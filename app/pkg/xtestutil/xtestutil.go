package xtestutil

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/datastore"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xstring"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"

	"github.com/google/go-cmp/cmp"
)

//  Assertion
// --------------------------------------------------

type Response struct {
	StatusCode int
	Location   string
	Body       []byte
}

func AssertResponse(t *testing.T, got *http.Response, want *Response) {
	t.Helper()
	t.Cleanup(func() {
		if err := got.Body.Close(); err != nil {
			t.Errorf("%s", xerr.FailedToCloseResponseBody)
		}
	})
	assertStatus(t, want, got)
	assertLocation(t, want, got)
	assertBody(t, want, got)
}

func AssertResponseJSON(t *testing.T, got *http.Response, want *Response) {
	t.Helper()
	t.Cleanup(func() {
		if err := got.Body.Close(); err != nil {
			t.Errorf("%s", xerr.FailedToCloseResponseBody)
		}
	})
	assertStatus(t, want, got)
	assertLocation(t, want, got)
	assertBodyJSON(t, want, got)
}

func assertStatus(t *testing.T, want *Response, got *http.Response) {
	t.Helper()

	if want.StatusCode != got.StatusCode {
		t.Errorf("status not matched ( want = %d; got = %d )", want.StatusCode, got.StatusCode)
	}
}

func assertLocation(t *testing.T, want *Response, got *http.Response) {
	t.Helper()

	gotLocation, err := got.Location()

	if err != nil && !errors.Is(err, http.ErrNoLocation) {
		t.Error(xerr.FailedToReadResponseLocation)
	}

	if len(want.Location) > 0 && want.Location != gotLocation.String() {
		t.Errorf("location not matched ( want = %s; got = %s )", want.Location, gotLocation)
	}
}

func assertBody(t *testing.T, want *Response, got *http.Response) {
	t.Helper()

	gotBody, err := io.ReadAll(got.Body)

	if err != nil {
		t.Errorf("%s", xerr.FailedToReadResponseBody)
	}

	if xutil.IsEmpty(want.Body) && xutil.IsEmpty(gotBody) {
		return
	}

	wb := strings.Replace(xstring.ByteToString(want.Body), "\n", "", -1)
	gb := strings.Replace(xstring.ByteToString(gotBody), "\n", "", -1)

	if wb != gb {
		t.Errorf("body not matched ( want = %s; got = %s", wb, gb)
	}
}

func assertBodyJSON(t *testing.T, want *Response, got *http.Response) {
	t.Helper()

	gotBody, err := io.ReadAll(got.Body)

	if err != nil {
		t.Errorf("%s", xerr.FailedToReadResponseBody)
	}

	if xutil.IsEmpty(want.Body) && xutil.IsEmpty(gotBody) {
		return
	}

	var wantJSON, gotJSON any

	if err = json.Unmarshal(want.Body, &wantJSON); err != nil {
		t.Errorf("%s", xerr.FailedToUnmarshalJSON)
	}

	if err = json.Unmarshal(gotBody, &gotJSON); err != nil {
		t.Errorf("%s", xerr.FailedToUnmarshalJSON)
	}

	if d := cmp.Diff(wantJSON, gotJSON); !xutil.IsEmpty(d) {
		t.Errorf("%s (-got +want)\n%s", xerr.ResponseBodyNotMatched, d)
	}
}

//  Error
// --------------------------------------------------

const DummyError DummyErr = "DUMMY ERROR"

type DummyErr string

func (v DummyErr) Error() string {
	return string(v)
}

//  Clock
// --------------------------------------------------

type FixedClocker struct{}

func (v FixedClocker) Now() time.Time {
	return time.Date(2000, 12, 31, 23, 59, 59, 0, time.UTC)
}

type FixedTomorrowClocker struct{}

func (v FixedTomorrowClocker) Now() time.Time {
	return FixedClocker{}.Now().Add(24 * time.Hour)
}

//  File I/O
// --------------------------------------------------

func LoadFile(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("%s: %v", xerr.FailedToReadFile, err)
	}
	return data
}

//  Data Store
// --------------------------------------------------

const TestDBHost = "127.0.0.1"
const TestDBPort = 13306
const TestDBUser = "idp_test"
const TestDBPassword = "idp_test"
const TestDBName = "idp_test"

func NewDatabase(t *testing.T) *datastore.Database {
	t.Helper()

	var cfg *config.Config
	var err error

	if cfg, err = config.New(); err != nil {
		t.Fatal(err)
	}

	cfg.DBAdmin = TestDBUser
	cfg.DBPassword = TestDBPassword
	cfg.DBHost = TestDBHost
	cfg.DBPort = TestDBPort
	cfg.DBName = TestDBName

	var db *datastore.Database

	if db, err = datastore.NewDatabase(context.Background(), cfg); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err = db.Client.Close(); err != nil {
			t.Fatal(err)
		}
	})

	return db
}

const TestRedisHost = "127.0.0.1"
const TestRedisPort = 16379
const TestRedisPassword = ""
const TestRedisDB = 1

func NewCache(t *testing.T) *datastore.Cache {
	t.Helper()

	var cfg *config.Config
	var err error

	if cfg, err = config.New(); err != nil {
		t.Fatal(err)
	}

	cfg.RedisHost = TestRedisHost
	cfg.RedisPort = TestRedisPort
	cfg.RedisPassword = TestRedisPassword
	cfg.RedisDB = TestRedisDB

	var cache *datastore.Cache

	if cache, err = datastore.NewCache(context.Background(), cfg); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err = cache.Client.Close(); err != nil {
			t.Fatal(err)
		}
	})

	return cache
}
