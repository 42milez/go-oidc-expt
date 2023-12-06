package httpstore

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/42milez/go-oidc-server/app/idp/config"

	"github.com/42milez/go-oidc-server/app/pkg/xtestutil"
	"github.com/gorilla/securecookie"
)

func TestCookie_Read(t *testing.T) {
	t.Parallel()

	hashKey := []byte("DiqrnXaMKuv4REtWbISL6iCiPJfhE0mVYIb966SimcUvBYTWutgxZDwVWdFOZn6T")
	blockKey := []byte("d0S2V8EhtPrgA4Ethow8PQAPC8sP6Dbf")

	sc := securecookie.New(hashKey, blockKey)

	name := "sid"
	want := "67ebdc7b-0c57-4002-a5c4-ddf36d2a9a47"

	encoded, err := sc.Encode(name, want)

	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodGet, "https://example.com", nil)
	req.AddCookie(&http.Cookie{
		Name:  name,
		Value: encoded,
	})

	cookie := NewCookie(hashKey, blockKey, xtestutil.FixedClocker{})
	got, err := cookie.Read(req, name)

	if err != nil {
		t.Error(err)
	}

	if got != want {
		t.Errorf("got = %s; want = %s", got, want)
	}
}

func TestCookie_Write(t *testing.T) {
	t.Parallel()

	hashKey := []byte("DiqrnXaMKuv4REtWbISL6iCiPJfhE0mVYIb966SimcUvBYTWutgxZDwVWdFOZn6T")
	blockKey := []byte("d0S2V8EhtPrgA4Ethow8PQAPC8sP6Dbf")

	cookie := NewCookie(hashKey, blockKey, xtestutil.FixedClocker{})

	name := "sid"
	value := "67ebdc7b-0c57-4002-a5c4-ddf36d2a9a47"

	if err := cookie.Write(httptest.NewRecorder(), name, value, config.SessionIDCookieTTL); err != nil {
		t.Error(err)
	}
}
