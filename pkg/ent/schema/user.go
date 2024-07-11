package schema

import (
	"fmt"
	"regexp"
	"time"

	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

const (
	nameMaxLength           = 30
	nameMinLength           = 6
	hashedPasswordLength    = 300
	encodedTotpSecretLength = 100
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Unique().
			Match(regexp.MustCompile(fmt.Sprintf("^[0-9a-z_]{%d,%d}$", nameMinLength, nameMaxLength))),
		field.String("password").
			SchemaType(map[string]string{
				dialect.MySQL: PasswordSchemaType(),
			}).
			NotEmpty().
			Sensitive(),
		// TOTP secret is encoded with base32 encoding.
		// https://datatracker.ietf.org/doc/html/rfc4648#page-8
		field.String("totp_secret").
			SchemaType(map[string]string{
				dialect.MySQL: TotpSecretSchemaType(),
			}).
			Optional().
			Match(regexp.MustCompile(fmt.Sprintf("^[A-Z2-7=]{%d}$", encodedTotpSecretLength))).
			Sensitive(),
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
		edge.To("consents", Consent.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func PasswordSchemaType() string {
	return fmt.Sprintf("VARCHAR(%d)", hashedPasswordLength)
}

func TotpSecretSchemaType() string {
	return fmt.Sprintf("CHAR(%d)", encodedTotpSecretLength)
}
