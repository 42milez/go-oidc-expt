package schema

import (
	"time"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"

	"entgo.io/ent/schema/index"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

const (
	codeType     = "CHAR(20)" // The length of authorization code
	codeLifetime = 10 * time.Minute
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
			GoType(typedef.UserID("")).
			SchemaType(map[string]string{
				dialect.MySQL: userIDType,
			}).
			NotEmpty().
			Immutable(),
		field.Time("expire_at").
			Default(func() time.Time {
				return time.Now().Add(codeLifetime)
			}),
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
