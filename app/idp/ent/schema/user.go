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
	userIDLength          = 26
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
				dialect.MySQL: UserIDSchemaType(),
			}).
			DefaultFunc(func() typedef.UserID {
				return xutil.MakeUserID()
			}).
			Immutable(),
		field.String("name").
			Match(regexp.MustCompile(fmt.Sprintf("^[0-9a-z_]{%d,%d}$", nameMinLength, nameMaxLength))).
			Unique().
			NotEmpty(),
		field.String("password_hash").
			SchemaType(map[string]string{
				dialect.MySQL: PasswordHashSchemaType(),
			}).
			Validate(func(s string) error {
				if len(s) > passwordHashMaxLength {
					return fmt.Errorf("password must be less than or equal to %d characters", passwordHashMaxLength)
				}
				return nil
			}).
			NotEmpty(),
		field.String("totp_secret").
			SchemaType(map[string]string{
				dialect.MySQL: TotoSecretSchemaType(),
			}).
			// TOTP secret is encoded with base32 encoding.
			// https://datatracker.ietf.org/doc/html/rfc4648#page-8
			Match(regexp.MustCompile(fmt.Sprintf("^[A-Z2-7=]{%d}$", totpSecretLength))).
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

func PasswordHashSchemaType() string {
	return fmt.Sprintf("CHAR(%d)", passwordHashMaxLength)
}

func TotoSecretSchemaType() string {
	return fmt.Sprintf("CHAR(%d)", totpSecretLength)
}

func UserIDSchemaType() string {
	return fmt.Sprintf("CHAR(%d)", userIDLength)
}
