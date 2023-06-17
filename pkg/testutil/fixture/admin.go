package fixture

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/42milez/go-oidc-server/app/idp/ent/alias"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

func Admin(admin *ent.Admin) *ent.Admin {
	ret := &ent.Admin{
		ID:         alias.AdminID(rand.Uint64()),
		Name:       "42milez" + strconv.Itoa(rand.Int())[:5],
		Password:   "42milez",
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	if admin == nil {
		return ret
	}

	if admin.ID != 0 {
		ret.ID = admin.ID
	}

	if admin.Name != "" {
		ret.Name = admin.Name
	}

	if admin.Password != "" {
		ret.Password = admin.Password
	}

	if !admin.CreatedAt.IsZero() {
		ret.CreatedAt = admin.CreatedAt
	}

	if !admin.ModifiedAt.IsZero() {
		ret.ModifiedAt = admin.ModifiedAt
	}

	return ret
}
