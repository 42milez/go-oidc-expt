package cookie

import (
	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/gorilla/securecookie"
)

type Util struct {
	sc *securecookie.SecureCookie
}

func NewUtil(cfg *config.Config) *Util {
	return &Util{
		sc: securecookie.New([]byte(cfg.CookieHashKey), []byte(cfg.CookieBlockKey)),
	}
}
