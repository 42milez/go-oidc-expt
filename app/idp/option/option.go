package option

import (
	"github.com/42milez/go-oidc-server/app/idp/datastore"
	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/idp/security"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
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