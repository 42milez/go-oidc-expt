package service

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

type Cookie struct {
	sc *securecookie.SecureCookie
}

func (p *Cookie) Read(r *http.Request, name string) (string, error) {
	ck, err := r.Cookie(name)

	if err != nil {
		return "", err
	}

	var ret string

	if err = p.sc.Decode(name, ck.Value, &ret); err != nil {
		return "", err
	}

	return ret, nil
}

func (p *Cookie) Write(w http.ResponseWriter, name, val string, ttl time.Duration) error {
	encoded, err := p.sc.Encode(name, val)

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

func NewCookie(hashKey, blockKey []byte) *Cookie {
	return &Cookie{
		sc: securecookie.New(hashKey, blockKey),
	}
}
