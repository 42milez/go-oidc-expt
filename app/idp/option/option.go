package option

import (
	datastore2 "github.com/42milez/go-oidc-server/app/idp/datastore"
	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/idp/security"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
	"github.com/go-playground/validator/v10"
)

type Option struct {
	Cache  *datastore2.Cache
	Cookie iface.CookieReadWriter
	DB     *datastore2.Database
	IdGen  *xid.UniqueID
	Token  *security.JWT
	V      *validator.Validate
}
