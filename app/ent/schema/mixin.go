package schema

import (
	"context"
	"fmt"

	"github.com/42milez/go-oidc-server/app/pkg/xid"
	"github.com/42milez/go-oidc-server/app/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/typedef"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/42milez/go-oidc-server/app/ent/ent/hook"
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
		SetID(id typedef.UserID)
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			is, ok := m.(IDSetter)

			if !ok {
				return nil, fmt.Errorf("unexpected mutation %T", m)
			}

			idGen, err := xid.GetUniqueID()

			if err != nil {
				return nil, err
			}

			id, err := idGen.NextID()

			if err != nil {
				return nil, err
			}

			is.SetID(typedef.UserID(id))

			return next.Mutate(ctx, m)
		})
	}
}
