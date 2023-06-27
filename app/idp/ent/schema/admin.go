package schema

import (
	"fmt"
	"regexp"
	"time"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"entgo.io/ent/dialect"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"

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
	idType           = "CHAR(26)"
	passwordHashType = "VARCHAR(751)"
	totpSecretType   = "CHAR(160)"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Account.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(typedef.AdminID("")).
			SchemaType(map[string]string{
				dialect.MySQL: idType,
			}).
			Immutable().
			DefaultFunc(func() typedef.AdminID {
				return xutil.MakeAdminID()
			}),
		field.String("name").
			MaxLen(nameMaxLength).
			MinLen(nameMinLength).
			Match(regexp.MustCompile("^\\D[0-9a-z_]+")).
			Unique().
			NotEmpty(),
		field.String("password_hash").
			GoType(typedef.PasswordHash("")).
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
			SchemaType(map[string]string{
				dialect.MySQL: totpSecretType,
			}).
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
