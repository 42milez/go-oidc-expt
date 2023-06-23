package schema

import (
	"fmt"
	"github.com/42milez/go-oidc-server/app/idp/ent/alias"
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

const (
	nameMaxLength     = 30
	nameMinLength     = 6
	passwordMaxLength = 100
	passwordMinLength = 8
	totpSecretLength  = 160
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Account.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", alias.AdminID{}).
			Unique().
			Default(alias.MakeAdminID),
		field.String("name").
			MaxLen(nameMaxLength).
			MinLen(nameMinLength).
			Match(regexp.MustCompile("^\\D[0-9a-z_]+")).
			Unique().
			NotEmpty(),
		field.String("password").
			MaxLen(passwordMaxLength).
			MinLen(passwordMinLength).
			NotEmpty(),
		field.String("totp_secret").
			MaxLen(totpSecretLength).
			Validate(func(s string) error {
				if len(s) != totpSecretLength {
					return fmt.Errorf("totp_secret must be %d characters", totpSecretLength)
				}
				return nil
			}).
			Nillable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("modified_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Account.
func (Admin) Edges() []ent.Edge {
	return nil
}
