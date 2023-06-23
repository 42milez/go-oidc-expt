package schema

import (
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/42milez/go-oidc-server/pkg/xutil"
	"regexp"
	"time"

	"github.com/42milez/go-oidc-server/app/idp/ent/alias"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

const (
	nameMaxLength      = 30
	nameMinLength      = 6
	passwordHashLength = 751
	totpSecretLength   = 160
)

const (
	idType = "CHAR(26)"
	passwordHashType = "VARCHAR(751)"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Account.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(alias.AdminID("")).
			SchemaType(map[string]string{
				dialect.MySQL: idType,
			}).
			Immutable().
			DefaultFunc(func() alias.AdminID {
				return alias.MakeAdminID()
		}),
		field.String("name").
			MaxLen(nameMaxLength).
			MinLen(nameMinLength).
			Match(regexp.MustCompile("^\\D[0-9a-z_]+")).
			Unique().
			NotEmpty(),
		field.String("password_hash").
			GoType(xutil.PasswordHash("")).
			SchemaType(map[string]string{
				dialect.MySQL: passwordHashType,
			}).
			Validate(func(s string) error {
				if len(s) > passwordHashLength {
					return fmt.Errorf("password must be %d characters", passwordHashLength)
				}
				return nil
			}).
			NotEmpty(),
		field.String("totp_secret").
			MaxLen(totpSecretLength).
			Validate(func(s string) error {
				if len(s) != totpSecretLength {
					return fmt.Errorf("totp_secret must be %d characters", totpSecretLength)
				}
				return nil
			}).
			Optional(),
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
