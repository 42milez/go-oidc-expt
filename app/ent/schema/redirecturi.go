package schema

import (
	"time"

	"github.com/42milez/go-oidc-server/app/typedef"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// RedirectUri holds the schema definition for the RedirectUri entity.
type RedirectUri struct {
	ent.Schema
}

// Annotations of the RedirectUri.
func (RedirectUri) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "redirect_uris",
		},
	}
}

// Fields of the RedirectUri.
func (RedirectUri) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			GoType(typedef.RedirectURIID(0)).
			Immutable(),
		field.String("uri").
			NotEmpty(),
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

// Edges of the RedirectUri.
func (RedirectUri) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("relying_party", RelyingParty.Type).
			Ref("redirect_uris").
			Field("relying_party_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Indexes of the RedirectUri.
func (RedirectUri) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("relying_party_id", "uri").
			Unique(),
	}
}
