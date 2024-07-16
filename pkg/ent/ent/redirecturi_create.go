// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/redirecturi"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-expt/pkg/typedef"
)

// RedirectURICreate is the builder for creating a RedirectURI entity.
type RedirectURICreate struct {
	config
	mutation *RedirectURIMutation
	hooks    []Hook
}

// SetURI sets the "uri" field.
func (ruc *RedirectURICreate) SetURI(s string) *RedirectURICreate {
	ruc.mutation.SetURI(s)
	return ruc
}

// SetCreatedAt sets the "created_at" field.
func (ruc *RedirectURICreate) SetCreatedAt(t time.Time) *RedirectURICreate {
	ruc.mutation.SetCreatedAt(t)
	return ruc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruc *RedirectURICreate) SetNillableCreatedAt(t *time.Time) *RedirectURICreate {
	if t != nil {
		ruc.SetCreatedAt(*t)
	}
	return ruc
}

// SetModifiedAt sets the "modified_at" field.
func (ruc *RedirectURICreate) SetModifiedAt(t time.Time) *RedirectURICreate {
	ruc.mutation.SetModifiedAt(t)
	return ruc
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (ruc *RedirectURICreate) SetNillableModifiedAt(t *time.Time) *RedirectURICreate {
	if t != nil {
		ruc.SetModifiedAt(*t)
	}
	return ruc
}

// SetRelyingPartyID sets the "relying_party_id" field.
func (ruc *RedirectURICreate) SetRelyingPartyID(tpi typedef.RelyingPartyID) *RedirectURICreate {
	ruc.mutation.SetRelyingPartyID(tpi)
	return ruc
}

// SetID sets the "id" field.
func (ruc *RedirectURICreate) SetID(tu typedef.RedirectURIID) *RedirectURICreate {
	ruc.mutation.SetID(tu)
	return ruc
}

// SetRelyingParty sets the "relying_party" edge to the RelyingParty entity.
func (ruc *RedirectURICreate) SetRelyingParty(r *RelyingParty) *RedirectURICreate {
	return ruc.SetRelyingPartyID(r.ID)
}

// Mutation returns the RedirectURIMutation object of the builder.
func (ruc *RedirectURICreate) Mutation() *RedirectURIMutation {
	return ruc.mutation
}

// Save creates the RedirectURI in the database.
func (ruc *RedirectURICreate) Save(ctx context.Context) (*RedirectURI, error) {
	ruc.defaults()
	return withHooks(ctx, ruc.sqlSave, ruc.mutation, ruc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ruc *RedirectURICreate) SaveX(ctx context.Context) *RedirectURI {
	v, err := ruc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ruc *RedirectURICreate) Exec(ctx context.Context) error {
	_, err := ruc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruc *RedirectURICreate) ExecX(ctx context.Context) {
	if err := ruc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruc *RedirectURICreate) defaults() {
	if _, ok := ruc.mutation.CreatedAt(); !ok {
		v := redirecturi.DefaultCreatedAt()
		ruc.mutation.SetCreatedAt(v)
	}
	if _, ok := ruc.mutation.ModifiedAt(); !ok {
		v := redirecturi.DefaultModifiedAt()
		ruc.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruc *RedirectURICreate) check() error {
	if _, ok := ruc.mutation.URI(); !ok {
		return &ValidationError{Name: "uri", err: errors.New(`ent: missing required field "RedirectURI.uri"`)}
	}
	if v, ok := ruc.mutation.URI(); ok {
		if err := redirecturi.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "RedirectURI.uri": %w`, err)}
		}
	}
	if _, ok := ruc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RedirectURI.created_at"`)}
	}
	if _, ok := ruc.mutation.ModifiedAt(); !ok {
		return &ValidationError{Name: "modified_at", err: errors.New(`ent: missing required field "RedirectURI.modified_at"`)}
	}
	if _, ok := ruc.mutation.RelyingPartyID(); !ok {
		return &ValidationError{Name: "relying_party_id", err: errors.New(`ent: missing required field "RedirectURI.relying_party_id"`)}
	}
	if _, ok := ruc.mutation.RelyingPartyID(); !ok {
		return &ValidationError{Name: "relying_party", err: errors.New(`ent: missing required edge "RedirectURI.relying_party"`)}
	}
	return nil
}

func (ruc *RedirectURICreate) sqlSave(ctx context.Context) (*RedirectURI, error) {
	if err := ruc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ruc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ruc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = typedef.RedirectURIID(id)
	}
	ruc.mutation.id = &_node.ID
	ruc.mutation.done = true
	return _node, nil
}

func (ruc *RedirectURICreate) createSpec() (*RedirectURI, *sqlgraph.CreateSpec) {
	var (
		_node = &RedirectURI{config: ruc.config}
		_spec = sqlgraph.NewCreateSpec(redirecturi.Table, sqlgraph.NewFieldSpec(redirecturi.FieldID, field.TypeUint64))
	)
	if id, ok := ruc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ruc.mutation.URI(); ok {
		_spec.SetField(redirecturi.FieldURI, field.TypeString, value)
		_node.URI = value
	}
	if value, ok := ruc.mutation.CreatedAt(); ok {
		_spec.SetField(redirecturi.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ruc.mutation.ModifiedAt(); ok {
		_spec.SetField(redirecturi.FieldModifiedAt, field.TypeTime, value)
		_node.ModifiedAt = value
	}
	if nodes := ruc.mutation.RelyingPartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   redirecturi.RelyingPartyTable,
			Columns: []string{redirecturi.RelyingPartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(relyingparty.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RelyingPartyID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RedirectURICreateBulk is the builder for creating many RedirectURI entities in bulk.
type RedirectURICreateBulk struct {
	config
	err      error
	builders []*RedirectURICreate
}

// Save creates the RedirectURI entities in the database.
func (rucb *RedirectURICreateBulk) Save(ctx context.Context) ([]*RedirectURI, error) {
	if rucb.err != nil {
		return nil, rucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rucb.builders))
	nodes := make([]*RedirectURI, len(rucb.builders))
	mutators := make([]Mutator, len(rucb.builders))
	for i := range rucb.builders {
		func(i int, root context.Context) {
			builder := rucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RedirectURIMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = typedef.RedirectURIID(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rucb *RedirectURICreateBulk) SaveX(ctx context.Context) []*RedirectURI {
	v, err := rucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rucb *RedirectURICreateBulk) Exec(ctx context.Context) error {
	_, err := rucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rucb *RedirectURICreateBulk) ExecX(ctx context.Context) {
	if err := rucb.Exec(ctx); err != nil {
		panic(err)
	}
}
