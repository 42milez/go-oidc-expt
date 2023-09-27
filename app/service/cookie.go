package service

import (
	"net/http"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/xtime"

	"github.com/gorilla/securecookie"
)

func NewCookie(hashKey, blockKey []byte, clock xtime.Clocker) *Cookie {
	return &Cookie{
		clock: clock,
		sc:    securecookie.New(hashKey, blockKey),
	}
}

type Cookie struct {
	clock xtime.Clocker
	sc    *securecookie.SecureCookie
}

func (c *Cookie) Read(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)

	if err != nil {
		return "", err
	}

	var ret string

	if err = c.sc.Decode(name, cookie.Value, &ret); err != nil {
		return "", err
	}

	return ret, nil
}

func (c *Cookie) Write(w http.ResponseWriter, name, val string, ttl time.Duration) error {
	encoded, err := c.sc.Encode(name, val)

	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    encoded,
		Path:     "/",
		MaxAge:   int(time.Now().Add(ttl).Unix()),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}
