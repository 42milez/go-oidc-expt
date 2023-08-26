package testutil

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/42milez/go-oidc-server/pkg/xstring"

	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/42milez/go-oidc-server/pkg/xutil"
	"github.com/google/go-cmp/cmp"
)

const (
	ErrResponseBodyNotMatched = "response body not matched"
)

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

	if err := json.Unmarshal(want.Body, &wantJSON); err != nil {
		t.Errorf("%s", xerr.FailedToUnmarshalJSON)
	}

	if err := json.Unmarshal(gotBody, &gotJSON); err != nil {
		t.Errorf("%s", xerr.FailedToUnmarshalJSON)
	}

	if d := cmp.Diff(wantJSON, gotJSON); !xutil.IsEmpty(d) {
		t.Errorf("%s (-got +want)\n%s", ErrResponseBodyNotMatched, d)
	}
}

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
