package option

import (
	"github.com/42milez/go-oidc-server/cmd/datastore"
	"github.com/42milez/go-oidc-server/cmd/iface"
	"github.com/42milez/go-oidc-server/cmd/security"
	"github.com/42milez/go-oidc-server/pkg/xid"
	"github.com/go-playground/validator/v10"
)

type Option struct {
	Cache  *datastore.Cache
	Cookie iface.CookieReadWriter
	DB     *datastore.Database
	IDGen  *xid.UniqueID
	Token  *security.JWT
	V      *validator.Validate
}
