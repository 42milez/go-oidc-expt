package schema

import (
	"time"

	"github.com/42milez/go-oidc-server/app/ent/typedef"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// RedirectURI holds the schema definition for the RedirectURI entity.
type RedirectURI struct {
	ent.Schema
}

// Annotations of the RedirectURI.
func (RedirectURI) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "redirect_uris",
		},
	}
}

// Fields of the RedirectURI.
func (RedirectURI) Fields() []ent.Field {
	return []ent.Field{
		field.String("uri").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("modified_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("user_id").
			GoType(typedef.UserID("")).
			SchemaType(map[string]string{
				dialect.MySQL: UserIDSchemaType(),
			}).
			NotEmpty().
			Immutable(),
	}
}

// Indexes of the RedirectURI.
func (RedirectURI) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "uri").
			Unique(),
	}
}
