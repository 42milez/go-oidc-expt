package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"regexp"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MinLen(6).
			MaxLen(30).
			Match(regexp.MustCompile("^\\D[0-9a-z_]+")).
			Unique().
			NotEmpty(),
		field.String("password").
			MinLen(8).
			MaxLen(100).
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now()).
			Immutable(),
		field.Time("modified_at").
			Default(time.Now()).
			UpdateDefault(time.Now()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
