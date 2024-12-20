// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/predicate"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/redirecturi"
)

// RedirectURIDelete is the builder for deleting a RedirectURI entity.
type RedirectURIDelete struct {
	config
	hooks    []Hook
	mutation *RedirectURIMutation
}

// Where appends a list predicates to the RedirectURIDelete builder.
func (rud *RedirectURIDelete) Where(ps ...predicate.RedirectURI) *RedirectURIDelete {
	rud.mutation.Where(ps...)
	return rud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rud *RedirectURIDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rud.sqlExec, rud.mutation, rud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rud *RedirectURIDelete) ExecX(ctx context.Context) int {
	n, err := rud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rud *RedirectURIDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(redirecturi.Table, sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeUint64))
	if ps := rud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rud.mutation.done = true
	return affected, err
}

// RedirectURIDeleteOne is the builder for deleting a single RedirectURI entity.
type RedirectURIDeleteOne struct {
	rud *RedirectURIDelete
}

// Where appends a list predicates to the RedirectURIDelete builder.
func (rudo *RedirectURIDeleteOne) Where(ps ...predicate.RedirectURI) *RedirectURIDeleteOne {
	rudo.rud.mutation.Where(ps...)
	return rudo
}

// Exec executes the deletion query.
func (rudo *RedirectURIDeleteOne) Exec(ctx context.Context) error {
	n, err := rudo.rud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{redirecturi.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rudo *RedirectURIDeleteOne) ExecX(ctx context.Context) {
	if err := rudo.Exec(ctx); err != nil {
		panic(err)
	}
}
