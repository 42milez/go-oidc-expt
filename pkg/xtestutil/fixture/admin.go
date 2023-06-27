package fixture

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

func Admin(admin *ent.Admin) *ent.Admin {
	ret := &ent.Admin{
		ID:           xutil.MakeAdminID(),
		Name:         "42milez" + strconv.Itoa(rand.Int())[:5],
		PasswordHash: "dP+BAwEBCmFyZ29uMlJlcHIB/4IAAQcBB1ZhcmlhbnQBDAABB1ZlcnNpb24BBAABBk1lbW9yeQEGAAEKSXRlcmF0aW9ucwEGAAELUGFyYWxsZWxpc20BBgABBFNhbHQBCgABDFBhc3N3b3JkSGFzaAEKAAAA/gGf/4IBCGFyZ29uMmlkASYB/QEAAAEDAQQB/4C55El7IkUMi4jedf0VNTMmx56zTbNcYQzWeMsLLgLUblXQNRQPoPAIs2UcBR/B0S20w5t2q75kueq/3r172G0u6h6WGiRNQjwJc46erfK49igzvQnbEIKnL0i2EssRJZ6KqQGcTViX7w2LWq8o0xAziyhv4LlNO+BOlFO1pCXJFQH+AQDeYSQf95j4BCP1huSk+kVvfJgbm8LIpWwGMe7s2m8EjTZIi0PoQmh31tdKjxnbePEBywkXa0ZDvjrPdW4Iuwu41Bn/NdFAnX6HHQUcmJgHPjES51WkdMxCiIWHXkFNma1tnISOJeCwL2mUfO5o3HH2agtKSafh2rDPo1H2lB0iRAfdZ/tIz5UBThBvMCHPEcAW3Odl0eXGPP7hIqHe/4p/qfZsD2golL518Duuy9R9pWZdBR9Q9WvkxHa/7V7NIsx5PDuH/0EAxgHiqZc4aYm61765QcOc7SQYbW+9ZmoKlr6/7KqQE2x2JgmkcUDPXfMyMt8iKVfzu/I1c5HuRDxLAA",
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

	if admin.PasswordHash != "" {
		ret.PasswordHash = admin.PasswordHash
	}

	if !admin.CreatedAt.IsZero() {
		ret.CreatedAt = admin.CreatedAt
	}

	if !admin.ModifiedAt.IsZero() {
		ret.ModifiedAt = admin.ModifiedAt
	}

	return ret
}
