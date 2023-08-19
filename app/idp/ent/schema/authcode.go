package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// AuthCode holds the schema definition for the AuthCode entity.
type AuthCode struct {
	ent.Schema
}

// Fields of the AuthCode.
func (AuthCode) Fields() []ent.Field {
	return []ent.Field{
		// TODO: Set length
		field.String("auth_code").
			Immutable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the AuthCode.
func (AuthCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("auth_codes").
			Unique().
			Field("user_name"),
	}
}

// Indexes of the AuthCode
func (AuthCode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("auth_code", "user_name").
			Unique(),
	}
}
