package option

import (
	"github.com/42milez/go-oidc-expt/cmd/datastore"
	"github.com/42milez/go-oidc-expt/cmd/iface"
	"github.com/42milez/go-oidc-expt/cmd/security"
	"github.com/42milez/go-oidc-expt/pkg/xid"
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
