package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/42milez/go-oidc-server/app/typedef"
	"time"
)

// Consent holds the schema definition for the Consent entity.
type Consent struct {
	ent.Schema
}

// Fields of the Consent.
func (Consent) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("user_id").
			GoType(typedef.UserID(0)).
			Immutable(),
		field.String("client_id").
			GoType(typedef.ClientId("")).
			NotEmpty().
			Immutable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Consent.
func (Consent) Edges() []ent.Edge {
	return nil
}
