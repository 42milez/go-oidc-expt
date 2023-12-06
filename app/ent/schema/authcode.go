package schema

import (
	"fmt"
	"regexp"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/42milez/go-oidc-server/app/idp/config"
)

// AuthCode holds the schema definition for the AuthCode entity.
type AuthCode struct {
	ent.Schema
}

// Fields of the AuthCode.
func (AuthCode) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			GoType(typedef.AuthCodeID(0)).
			Immutable(),
		field.String("code").
			SchemaType(map[string]string{
				dialect.MySQL: AuthCodeSchemaType(),
			}).
			NotEmpty().
			Match(regexp.MustCompile(fmt.Sprintf("^[0-9a-zA-Z]{%d}$", config.AuthCodeLength))).
			Immutable(),
		field.Uint64("user_id").
			GoType(typedef.UserID(0)).
			Immutable(),
		field.Time("expire_at").
			Default(func() time.Time {
				return time.Now().Add(config.AuthCodeTTL)
			}).
			Immutable(),
		field.Time("used_at").
			Optional().
			Nillable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("modified_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Uint64("relying_party_id").
			GoType(typedef.RelyingPartyID(0)).
			Immutable(),
	}
}

// Edges of the AuthCode.
func (AuthCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("relying_party", RelyingParty.Type).
			Ref("auth_codes").
			Field("relying_party_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Indexes of the AuthCode.
func (AuthCode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("relying_party_id", "code").
			Unique(),
	}
}

func AuthCodeSchemaType() string {
	return fmt.Sprintf("CHAR(%d)", config.AuthCodeLength)
}
