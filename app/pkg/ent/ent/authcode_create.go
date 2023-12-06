// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/42milez/go-oidc-server/app/pkg/ent/ent/authcode"
	"github.com/42milez/go-oidc-server/app/pkg/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-server/app/pkg/typedef"
)

// AuthCodeCreate is the builder for creating a AuthCode entity.
type AuthCodeCreate struct {
	config
	mutation *AuthCodeMutation
	hooks    []Hook
}

// SetCode sets the "code" field.
func (acc *AuthCodeCreate) SetCode(s string) *AuthCodeCreate {
	acc.mutation.SetCode(s)
	return acc
}

// SetUserID sets the "user_id" field.
func (acc *AuthCodeCreate) SetUserID(ti typedef.UserID) *AuthCodeCreate {
	acc.mutation.SetUserID(ti)
	return acc
}

// SetExpireAt sets the "expire_at" field.
func (acc *AuthCodeCreate) SetExpireAt(t time.Time) *AuthCodeCreate {
	acc.mutation.SetExpireAt(t)
	return acc
}

// SetNillableExpireAt sets the "expire_at" field if the given value is not nil.
func (acc *AuthCodeCreate) SetNillableExpireAt(t *time.Time) *AuthCodeCreate {
	if t != nil {
		acc.SetExpireAt(*t)
	}
	return acc
}

// SetUsedAt sets the "used_at" field.
func (acc *AuthCodeCreate) SetUsedAt(t time.Time) *AuthCodeCreate {
	acc.mutation.SetUsedAt(t)
	return acc
}

// SetNillableUsedAt sets the "used_at" field if the given value is not nil.
func (acc *AuthCodeCreate) SetNillableUsedAt(t *time.Time) *AuthCodeCreate {
	if t != nil {
		acc.SetUsedAt(*t)
	}
	return acc
}

// SetCreatedAt sets the "created_at" field.
func (acc *AuthCodeCreate) SetCreatedAt(t time.Time) *AuthCodeCreate {
	acc.mutation.SetCreatedAt(t)
	return acc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acc *AuthCodeCreate) SetNillableCreatedAt(t *time.Time) *AuthCodeCreate {
	if t != nil {
		acc.SetCreatedAt(*t)
	}
	return acc
}

// SetModifiedAt sets the "modified_at" field.
func (acc *AuthCodeCreate) SetModifiedAt(t time.Time) *AuthCodeCreate {
	acc.mutation.SetModifiedAt(t)
	return acc
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (acc *AuthCodeCreate) SetNillableModifiedAt(t *time.Time) *AuthCodeCreate {
	if t != nil {
		acc.SetModifiedAt(*t)
	}
	return acc
}

// SetRelyingPartyID sets the "relying_party_id" field.
func (acc *AuthCodeCreate) SetRelyingPartyID(tpi typedef.RelyingPartyID) *AuthCodeCreate {
	acc.mutation.SetRelyingPartyID(tpi)
	return acc
}

// SetID sets the "id" field.
func (acc *AuthCodeCreate) SetID(tci typedef.AuthCodeID) *AuthCodeCreate {
	acc.mutation.SetID(tci)
	return acc
}

// SetRelyingParty sets the "relying_party" edge to the RelyingParty entity.
func (acc *AuthCodeCreate) SetRelyingParty(r *RelyingParty) *AuthCodeCreate {
	return acc.SetRelyingPartyID(r.ID)
}

// Mutation returns the AuthCodeMutation object of the builder.
func (acc *AuthCodeCreate) Mutation() *AuthCodeMutation {
	return acc.mutation
}

// Save creates the AuthCode in the database.
func (acc *AuthCodeCreate) Save(ctx context.Context) (*AuthCode, error) {
	acc.defaults()
	return withHooks(ctx, acc.sqlSave, acc.mutation, acc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (acc *AuthCodeCreate) SaveX(ctx context.Context) *AuthCode {
	v, err := acc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acc *AuthCodeCreate) Exec(ctx context.Context) error {
	_, err := acc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acc *AuthCodeCreate) ExecX(ctx context.Context) {
	if err := acc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acc *AuthCodeCreate) defaults() {
	if _, ok := acc.mutation.ExpireAt(); !ok {
		v := authcode.DefaultExpireAt()
		acc.mutation.SetExpireAt(v)
	}
	if _, ok := acc.mutation.CreatedAt(); !ok {
		v := authcode.DefaultCreatedAt()
		acc.mutation.SetCreatedAt(v)
	}
	if _, ok := acc.mutation.ModifiedAt(); !ok {
		v := authcode.DefaultModifiedAt()
		acc.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acc *AuthCodeCreate) check() error {
	if _, ok := acc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "AuthCode.code"`)}
	}
	if v, ok := acc.mutation.Code(); ok {
		if err := authcode.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "AuthCode.code": %w`, err)}
		}
	}
	if _, ok := acc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "AuthCode.user_id"`)}
	}
	if _, ok := acc.mutation.ExpireAt(); !ok {
		return &ValidationError{Name: "expire_at", err: errors.New(`ent: missing required field "AuthCode.expire_at"`)}
	}
	if _, ok := acc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AuthCode.created_at"`)}
	}
	if _, ok := acc.mutation.ModifiedAt(); !ok {
		return &ValidationError{Name: "modified_at", err: errors.New(`ent: missing required field "AuthCode.modified_at"`)}
	}
	if _, ok := acc.mutation.RelyingPartyID(); !ok {
		return &ValidationError{Name: "relying_party_id", err: errors.New(`ent: missing required field "AuthCode.relying_party_id"`)}
	}
	if _, ok := acc.mutation.RelyingPartyID(); !ok {
		return &ValidationError{Name: "relying_party", err: errors.New(`ent: missing required edge "AuthCode.relying_party"`)}
	}
	return nil
}

func (acc *AuthCodeCreate) sqlSave(ctx context.Context) (*AuthCode, error) {
	if err := acc.check(); err != nil {
		return nil, err
	}
	_node, _spec := acc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = typedef.AuthCodeID(id)
	}
	acc.mutation.id = &_node.ID
	acc.mutation.done = true
	return _node, nil
}

func (acc *AuthCodeCreate) createSpec() (*AuthCode, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthCode{config: acc.config}
		_spec = sqlgraph.NewCreateSpec(authcode.Table, sqlgraph.NewFieldSpec(authcode.FieldID, field.TypeUint64))
	)
	if id, ok := acc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := acc.mutation.Code(); ok {
		_spec.SetField(authcode.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := acc.mutation.UserID(); ok {
		_spec.SetField(authcode.FieldUserID, field.TypeUint64, value)
		_node.UserID = value
	}
	if value, ok := acc.mutation.ExpireAt(); ok {
		_spec.SetField(authcode.FieldExpireAt, field.TypeTime, value)
		_node.ExpireAt = value
	}
	if value, ok := acc.mutation.UsedAt(); ok {
		_spec.SetField(authcode.FieldUsedAt, field.TypeTime, value)
		_node.UsedAt = &value
	}
	if value, ok := acc.mutation.CreatedAt(); ok {
		_spec.SetField(authcode.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := acc.mutation.ModifiedAt(); ok {
		_spec.SetField(authcode.FieldModifiedAt, field.TypeTime, value)
		_node.ModifiedAt = value
	}
	if nodes := acc.mutation.RelyingPartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   authcode.RelyingPartyTable,
			Columns: []string{authcode.RelyingPartyColumn},
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

// AuthCodeCreateBulk is the builder for creating many AuthCode entities in bulk.
type AuthCodeCreateBulk struct {
	config
	err      error
	builders []*AuthCodeCreate
}

// Save creates the AuthCode entities in the database.
func (accb *AuthCodeCreateBulk) Save(ctx context.Context) ([]*AuthCode, error) {
	if accb.err != nil {
		return nil, accb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(accb.builders))
	nodes := make([]*AuthCode, len(accb.builders))
	mutators := make([]Mutator, len(accb.builders))
	for i := range accb.builders {
		func(i int, root context.Context) {
			builder := accb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthCodeMutation)
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
					_, err = mutators[i+1].Mutate(root, accb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, accb.driver, spec); err != nil {
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
					nodes[i].ID = typedef.AuthCodeID(id)
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
		if _, err := mutators[0].Mutate(ctx, accb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (accb *AuthCodeCreateBulk) SaveX(ctx context.Context) []*AuthCode {
	v, err := accb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (accb *AuthCodeCreateBulk) Exec(ctx context.Context) error {
	_, err := accb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (accb *AuthCodeCreateBulk) ExecX(ctx context.Context) {
	if err := accb.Exec(ctx); err != nil {
		panic(err)
	}
}