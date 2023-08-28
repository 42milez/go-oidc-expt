package schema

import (
	"fmt"
	"regexp"
	"time"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/ent/typedef"

	"entgo.io/ent/schema/index"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
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
				dialect.MySQL: AuthCodeSchemaType(),
			}).
			Match(regexp.MustCompile(fmt.Sprintf("^[0-9a-zA-Z]{%d}$", config.AuthCodeLength))).
			NotEmpty().
			Immutable(),
		field.Time("expire_at").
			Default(func() time.Time {
				return time.Now().Add(config.AuthCodeLifetime)
			}).
			Immutable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Uint64("user_id").
			GoType(typedef.UserID(0)).
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

func AuthCodeSchemaType() string {
	return fmt.Sprintf("CHAR(%d)", config.AuthCodeLength)
}
