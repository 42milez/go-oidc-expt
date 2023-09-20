package schema

import (
	"fmt"
	"regexp"
	"time"

	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// AuthCode holds the schema definition for the AuthCode entity.
type AuthCode struct {
	ent.Schema
}

// Fields of the AuthCode.
func (AuthCode) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").GoType(typedef.AuthCodeID(0)),
		field.String("code").
			SchemaType(map[string]string{
				dialect.MySQL: AuthCodeSchemaType(),
			}).
			Match(regexp.MustCompile(fmt.Sprintf("^[0-9a-zA-Z]{%d}$", config.AuthCodeLength))).
			NotEmpty().
			Immutable(),
		field.Time("expire_at").
			Default(func() time.Time {
				return time.Now().Add(config.AuthCodeLifetime)
			}).
			Immutable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("used_at").
			Optional().
			Nillable(),
		field.Uint64("relying_party_auth_codes").
			GoType(typedef.RelyingPartyID(0)).
			Immutable(),
	}
}

// Edges of the AuthCode.
func (AuthCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("relying_party", RelyingParty.Type).
			Ref("auth_codes").
			Unique(),
	}
}

// Indexes of the AuthCode.
func (AuthCode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("relying_party_auth_codes", "code").
			Unique(),
	}
}

func AuthCodeSchemaType() string {
	return fmt.Sprintf("CHAR(%d)", config.AuthCodeLength)
}
