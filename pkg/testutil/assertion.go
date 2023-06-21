package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/42milez/go-oidc-server/pkg/xutil"
	"github.com/google/go-cmp/cmp"
)

const (
	ErrResponseBodyNotMatched = "response body not matched"
)

func AssertResponse(t *testing.T, got *http.Response, wantStatusCode int, wantBody []byte) {
	t.Helper()
	t.Cleanup(func() {
		if err := got.Body.Close(); err != nil {
			t.Fatalf("%s", xerr.FailedToCloseResponseBody)
		}
	})

	gotBody, err := io.ReadAll(got.Body)

	if err != nil {
		t.Fatalf("%s", xerr.FailedToResponseBody)
	}

	if wantStatusCode != got.StatusCode {
		t.Fatalf("want = %d; got = %d", wantStatusCode, got.StatusCode)
	}

	if xutil.IsEmpty(wantBody) && xutil.IsEmpty(gotBody) {
		return
	}

	AssertJSON(t, wantBody, gotBody)
}

func AssertJSON(t *testing.T, wantBody, gotBody []byte) {
	t.Helper()

	var wantJSON, gotJSON any

	if err := json.Unmarshal(wantBody, &wantJSON); err != nil {
		t.Fatalf("%s", xerr.FailedToUnmarshalJSON)
	}

	if err := json.Unmarshal(gotBody, &gotJSON); err != nil {
		t.Fatalf("%s", xerr.FailedToUnmarshalJSON)
	}

	if d := cmp.Diff(wantJSON, gotJSON); !xutil.IsEmpty(d) {
		t.Fatalf("%s (-got +want)\n%s", ErrResponseBodyNotMatched, d)
	}
}
