package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/42milez/go-oidc-server/app/typedef"
	"time"
)

// RelyingParty holds the schema definition for the RelyingParty entity.
type RelyingParty struct {
	ent.Schema
}

// Fields of the RelyingParty.
func (RelyingParty) Fields() []ent.Field {
	return []ent.Field{
		field.String("client_id").
			GoType(typedef.ClientId("")).
			Unique().
			Immutable(),
		field.String("client_secret").
			GoType(typedef.ClientSecret("")).
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("modified_at").
			Default(time.Now),
	}
}

// Edges of the RelyingParty.
func (RelyingParty) Edges() []ent.Edge {
	return nil
}
