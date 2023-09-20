package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/42milez/go-oidc-server/app/typedef"
)

// RelyingParty holds the schema definition for the RelyingParty entity.
type RelyingParty struct {
	ent.Schema
}

// Fields of the RelyingParty.
func (RelyingParty) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			GoType(typedef.RelyingPartyID(0)),
		field.String("client_id").
			GoType(typedef.ClientID("")).
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
	return []ent.Edge{
		edge.To("auth_codes", AuthCode.Type).
			StorageKey(edge.Column("relying_party_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("redirect_uris", RedirectURI.Type).
			StorageKey(edge.Column("relying_party_id")).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
