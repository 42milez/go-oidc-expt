package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"fmt"
	"github.com/42milez/go-oidc-server/app/ent/ent/hook"
	"github.com/42milez/go-oidc-server/app/ent/typedef"
	"github.com/42milez/go-oidc-server/pkg/xid"
	"github.com/42milez/go-oidc-server/pkg/xutil"
)

// BaseMixin to be shared will all different schemas.
type BaseMixin struct {
	mixin.Schema
}

// Fields of the Mixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			GoType(typedef.UserID(0)).
			Immutable().
			Annotations(entsql.Annotation{
				Incremental: xutil.NewFalse(),
			}),
	}
}

// Hooks of the Mixin.
// https://entgo.io/docs/faq/#how-to-use-a-custom-generator-of-ids
func (BaseMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(IDHook(), ent.OpCreate),
	}
}

func IDHook() ent.Hook {
	type IDSetter interface {
		SetID(uint64)
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			is, ok := m.(IDSetter)

			if !ok {
				return nil, fmt.Errorf("unexpected mutation %T", m)
			}

			id, err := xid.UID.NextID()

			if err != nil {
				return nil, err
			}

			is.SetID(id)

			return next.Mutate(ctx, m)
		})
	}
}
