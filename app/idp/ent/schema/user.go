package schema

import (
	"errors"
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
		field.String("totp_secret").
			Validate(func(s string) error {
				if len(s) != 160 {
					return errors.New("totp_secret must be 160 characters")
				}
				return nil
			}).
			Nillable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("modified_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
