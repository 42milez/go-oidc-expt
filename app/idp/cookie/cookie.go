package cookie

import (
	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/gorilla/securecookie"
	"net/http"
)

type Util struct {
	sc *securecookie.SecureCookie
}

func (p *Util) SetSessionID(w http.ResponseWriter, id string) error {
	name := "sid"
	encoded, err := p.sc.Encode(name, id)

	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  name,
		Value: encoded,
		Path:  "/",
		MaxAge: 0,
		Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}

func NewUtil(cfg *config.Config) *Util {
	return &Util{
		sc: securecookie.New([]byte(cfg.CookieHashKey), []byte(cfg.CookieBlockKey)),
	}
}
