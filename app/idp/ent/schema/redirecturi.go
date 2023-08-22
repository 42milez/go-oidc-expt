package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
)

// RedirectURI holds the schema definition for the RedirectURI entity.
type RedirectURI struct {
	ent.Schema
}

// Annotations of the RedirectURI.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "redirect_uris"},
	}
}

// Fields of the RedirectURI.
func (RedirectURI) Fields() []ent.Field {
	return []ent.Field{
		field.String("uri").
			NotEmpty(),
		field.String("user_id").
			GoType(typedef.UserID("")).
			SchemaType(map[string]string{
				dialect.MySQL: userIDType,
			}).
			NotEmpty().
			Immutable(),
		field.Time("created_at").
			Default(time.Now).
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
