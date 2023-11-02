package option

import (
	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/iface"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
	"github.com/42milez/go-oidc-server/app/security"
	"github.com/go-playground/validator/v10"
)

type Option struct {
	Cache  *datastore.Cache
	Cookie iface.CookieReadWriter
	DB     *datastore.Database
	IdGen  *xid.UniqueID
	Token  *security.JWT
	V      *validator.Validate
}
