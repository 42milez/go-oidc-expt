package schema

import (
	"fmt"
	"regexp"
	"time"

	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"entgo.io/ent/dialect"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

const (
	nameMaxLength         = 30
	nameMinLength         = 6
	passwordHashMaxLength = 1000
	totpSecretLength      = 160
)

const (
	userIDType       = "CHAR(26)"
	passwordHashType = "VARCHAR(1000)"
	totpSecretType   = "CHAR(160)"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(typedef.UserID("")).
			SchemaType(map[string]string{
				dialect.MySQL: userIDType,
			}).
			Immutable().
			DefaultFunc(func() typedef.UserID {
				return xutil.MakeUserID()
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
				if len(s) > passwordHashMaxLength {
					return fmt.Errorf("password must be %d characters", passwordHashMaxLength)
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

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("auth_codes", AuthCode.Type).
			StorageKey(edge.Column("user_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("redirect_uris", RedirectURI.Type).
			StorageKey(edge.Column("user_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
