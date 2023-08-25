package cookie

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

type Cookie struct {
	sc *securecookie.SecureCookie
}

func (p *Cookie) Set(w http.ResponseWriter, name, val string) error {
	encoded, err := p.sc.Encode(name, val)

	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    encoded,
		Path:     "/",
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}

func NewCookie(hashKey, blockKey string) *Cookie {
	return &Cookie{
		sc: securecookie.New([]byte(hashKey), []byte(blockKey)),
	}
}
