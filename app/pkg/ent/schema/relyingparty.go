package schema

import (
	"fmt"
	"regexp"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"entgo.io/ent/schema/index"
	"github.com/42milez/go-oidc-server/app/idp/config"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RelyingParty holds the schema definition for the RelyingParty entity.
type RelyingParty struct {
	ent.Schema
}

// Fields of the RelyingParty.
func (RelyingParty) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			GoType(typedef.RelyingPartyID(0)).
			Immutable(),
		field.String("client_id").
			Match(regexp.MustCompile(fmt.Sprintf("^[0-9a-zA-Z]{%d}$", config.ClientIDLength))).
			Unique().
			Immutable(),
		field.String("client_secret").
			Match(regexp.MustCompile(fmt.Sprintf("^[0-9a-zA-Z]{%d}$", config.ClientSecretLength))).
			NotEmpty().
			Sensitive(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("modified_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the RelyingParty.
func (RelyingParty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("auth_codes", AuthCode.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("redirect_uris", RedirectUri.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (RelyingParty) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("client_id", "client_secret").
			Unique(),
	}
}
