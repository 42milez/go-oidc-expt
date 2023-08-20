package schema

import (
	"entgo.io/ent/schema/index"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

const (
	codeType = "CHAR(20)" // The length of authorization code
)

// AuthCode holds the schema definition for the AuthCode entity.
type AuthCode struct {
	ent.Schema
}

// Fields of the AuthCode.
func (AuthCode) Fields() []ent.Field {
	return []ent.Field{
		// TODO: Set length
		field.String("code").
			SchemaType(map[string]string{
				dialect.MySQL: codeType,
			}).
			NotEmpty().
			Immutable(),
		field.String("user_id").
			Optional().
			Immutable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Indexes of the AuthCode.
func (AuthCode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "code").
			Unique(),
	}
}
