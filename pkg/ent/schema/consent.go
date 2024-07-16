package schema

import (
	"time"

	"github.com/42milez/go-oidc-expt/pkg/typedef"

	"entgo.io/ent/schema/edge"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Consent holds the schema definition for the Consent entity.
type Consent struct {
	ent.Schema
}

// Fields of the Consent.
func (Consent) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			GoType(typedef.ConsentID(0)).
			Immutable(),
		field.String("client_id").
			GoType(typedef.ClientID("")).
			Immutable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Uint64("user_id").
			GoType(typedef.UserID(0)).
			Immutable(),
	}
}

func (Consent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("consents").
			Field("user_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Indexes of the Consent
func (Consent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "client_id").
			Unique(),
	}
}
