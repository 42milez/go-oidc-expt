package fixture

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

func Admin(admin *ent.User) *ent.User {
	ret := &ent.User{
		ID:           xutil.MakeUserID(),
		Name:         "42milez" + strconv.Itoa(rand.Int())[:5],
		Password: "/4L/gQMBAQphcmdvbjJSZXByAf+CAAEIAQdWYXJpYW50AQYAAQdWZXJzaW9uAQQAAQZNZW1vcnkBBgABCkl0ZXJhdGlvbnMBBgABC1BhcmFsbGVsaXNtAQYAAQlLZXlMZW5ndGgBBgABBFNhbHQBCgABDFBhc3N3b3JkSGFzaAEKAAAARv+CAQIBJgH9AQAAAQMBBAEgARBBnF9UwRGCjn9L7BI3g5yiASA6O7i2oGmVB9EAJYYhFvIU+jdvr/UMkQqmeOc0uvsyjAA",
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
	}

	if admin == nil {
		return ret
	}

	if !admin.ID.IsZero() {
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
